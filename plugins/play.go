package plugins

import (
	"fmt"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/emojis"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/formatters"
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
		cmd := "play"
		if video {
			cmd = "vplay"
		}
		_, err := m.Reply(fmt.Sprintf("Usage:\n• <code>/%s song name</code>\n• Reply to audio/video with <code>/%s</code>", cmd, cmd), &telegram.SendOptions{ParseMode: "HTML"})
		return err
	}

	status, err := m.Reply(emojis.Progress()+" "+formatters.Smallcaps("getting your query baby..."), &telegram.SendOptions{ParseMode: "HTML"})
	if err != nil {
		return err
	}

	res, err := platforms.Resolve(args)
	if err != nil {
		_, _ = status.Edit("❌ Search failed: "+escapeHTML(err.Error()), &telegram.SendOptions{ParseMode: "HTML"})
		return nil
	}

	reqName := "User"
	if m.Sender != nil {
		reqName = m.Sender.FirstName
	}
	track := queue.Track{
		Title:       res.Title,
		URL:         res.URL,
		Duration:    res.DurationStr(),
		Thumbnail:   res.Thumbnail,
		RequestedBy: reqName,
		IsVideo:     video,
		ID:          res.ID,
	}

	chatID := m.ChatID()
	pos := b.Queue.Add(chatID, track)

	if pos > 1 {
		cap := formatters.QueueCaption(pos, track.Title, track.Duration, track.RequestedBy)
		_, _ = status.Edit(cap, &telegram.SendOptions{ParseMode: "HTML", ReplyMarkup: playerKeyboard(chatID)})
		return nil
	}

	_, _ = status.Edit(emojis.Progress()+" "+formatters.Smallcaps(fmt.Sprintf("downloading %s baby...", map[bool]string{true: "video", false: "audio"}[video])), &telegram.SendOptions{ParseMode: "HTML"})

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
		_, _ = status.Edit("❌ Download failed — "+escapeHTML(msg), &telegram.SendOptions{ParseMode: "HTML"})
		return nil
	}

	_, _ = status.Edit(emojis.Stream()+" "+formatters.Smallcaps("starting stream baby..."), &telegram.SendOptions{ParseMode: "HTML"})

	header := "sᴛʀᴇᴀᴍɪɴɢ ɪɴ ᴠᴄ"
	if video {
		header += " (ᴠɪᴅᴇᴏ)"
	}
	streamed := false
	if b.Stream != nil && b.Stream.Enabled() {
		if err := b.Stream.PlayFile(chatID, file, track); err == nil {
			streamed = true
		}
	}

	cap := formatters.PanelCaption(track.Title, track.Duration, track.RequestedBy, header)
	if !streamed {
		_, _ = b.Client.SendMedia(chatID, file, &telegram.MediaOptions{
			Caption:   cap,
			ParseMode: "HTML",
		})
	}
	_, _ = status.Edit(cap, &telegram.SendOptions{ParseMode: "HTML", ReplyMarkup: playerKeyboard(chatID)})
	return nil
}

func escapeHTML(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;")
	return r.Replace(s)
}
