package plugins

import (
	"fmt"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/emojis"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/queue"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/platforms"
)

func registerPlay(b *bot.Bot) {
	b.On("message:/play", func(m *telegram.NewMessage) error {
		return handlePlay(b, m, false)
	})
	b.On("message:/vplay", func(m *telegram.NewMessage) error {
		return handlePlay(b, m, true)
	})
}

func handlePlay(b *bot.Bot, m *telegram.NewMessage, video bool) error {
	args := strings.TrimSpace(m.Args())
	if args == "" {
		_, err := m.Reply(emojis.Warn()+" Usage: <code>/play song name or youtube url</code>", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	}

	status, err := m.Reply(emojis.Spark()+" Searching...")
	if err != nil {
		return err
	}

	res, err := platforms.Resolve(args)
	if err != nil {
		_, _ = status.Edit(emojis.Warn()+" Search failed: "+escapeHTML(err.Error()), &telegram.SendOptions{ParseMode: "HTML"})
		return nil
	}

	track := queue.Track{
		Title:       res.Title,
		URL:         res.URL,
		Duration:    res.DurationStr(),
		Thumbnail:   res.Thumbnail,
		RequestedBy: m.Sender.FirstName,
		IsVideo:     video,
		ID:          res.ID,
	}

	chatID := m.ChatID()
	pos := b.Queue.Add(chatID, track)

	if pos > 1 {
		text := fmt.Sprintf("%s <b>Added to queue</b> (#%d)\n\n%s <b>%s</b>\n%s %s\n%s %s",
			emojis.Queue(), pos,
			emojis.Title(), escapeHTML(track.Title),
			emojis.Duration(), track.Duration,
			emojis.Request(), escapeHTML(track.RequestedBy),
		)
		_, _ = status.Edit(text, &telegram.SendOptions{ParseMode: "HTML", ReplyMarkup: playerKeyboard(chatID)})
		return nil
	}

	_, _ = status.Edit(fmt.Sprintf("%s <b>Downloading</b>\n\n%s <b>%s</b>", emojis.Progress(), emojis.Title(), escapeHTML(track.Title)), &telegram.SendOptions{ParseMode: "HTML"})

	var file string
	var dlErr error
	if video {
		file, dlErr = platforms.DownloadVideo(track.ID)
	} else {
		file, dlErr = platforms.DownloadSong(track.ID)
	}
	if dlErr != nil || file == "" {
		msg := "unknown"
		if dlErr != nil {
			msg = dlErr.Error()
		}
		_, _ = status.Edit(fmt.Sprintf("%s Download failed\n<code>%s</code>", emojis.Warn(), escapeHTML(msg)), &telegram.SendOptions{ParseMode: "HTML"})
		return nil
	}

	// Try VC stream first
	streamed := false
	if b.Stream != nil && b.Stream.Enabled() {
		if err := b.Stream.PlayFile(chatID, file, track); err == nil {
			streamed = true
		}
	}

	cap := fmt.Sprintf("%s <b>Now playing</b>\n\n%s <b>%s</b>\n%s %s\n%s %s",
		emojis.Play(),
		emojis.Title(), escapeHTML(track.Title),
		emojis.Duration(), track.Duration,
		emojis.Request(), escapeHTML(track.RequestedBy),
	)
	if streamed {
		cap += "\n\n" + emojis.Stream() + " <b>Streaming in Voice Chat</b>"
	} else {
		_, _ = b.Client.SendMedia(chatID, file, &telegram.MediaOptions{
			Caption:   cap + "\n\n<code>VC needs STRING_SESSION — file sent to chat</code>",
			ParseMode: "HTML",
		})
		cap += "\n\n📎 Audio file sent (set STRING_SESSION for VC)"
	}

	_, _ = status.Edit(cap, &telegram.SendOptions{ParseMode: "HTML", ReplyMarkup: playerKeyboard(chatID)})
	return nil
}

func escapeHTML(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;")
	return r.Replace(s)
}
