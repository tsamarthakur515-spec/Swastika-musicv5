package plugins

import (
	"fmt"
	"os"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
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
		_, err := m.Reply("❗ Usage: <code>/play song name or youtube url</code>", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	}

	status, err := m.Reply("🔎 Searching...")
	if err != nil {
		return err
	}

	res, err := platforms.Resolve(args)
	if err != nil {
		_, _ = status.Edit("❌ Search failed: "+escapeHTML(err.Error()), &telegram.SendOptions{ParseMode: "HTML"})
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

	kind := "🎵"
	if video {
		kind = "🎬"
	}

	if pos > 1 {
		text := fmt.Sprintf("➕ <b>Added to queue</b> (#%d)\n\n<b>%s</b>\n⏱ %s\n👤 %s", pos, escapeHTML(track.Title), track.Duration, escapeHTML(track.RequestedBy))
		_, _ = status.Edit(text, &telegram.SendOptions{ParseMode: "HTML"})
		return nil
	}

	_, _ = status.Edit(fmt.Sprintf("%s <b>Downloading</b>\n\n<b>%s</b>\n⏱ %s\n👤 %s", kind, escapeHTML(track.Title), track.Duration, escapeHTML(track.RequestedBy)), &telegram.SendOptions{ParseMode: "HTML"})

	var file string
	var dlErr error
	if video {
		file, dlErr = platforms.DownloadVideo(track.ID)
	} else {
		file, dlErr = platforms.DownloadSong(track.ID)
	}
	if dlErr != nil || file == "" {
		msg := "unknown error"
		if dlErr != nil {
			msg = dlErr.Error()
		}
		_, _ = status.Edit(fmt.Sprintf("❌ Download failed\n\n<b>%s</b>\n<code>%s</code>", escapeHTML(track.Title), escapeHTML(msg)), &telegram.SendOptions{ParseMode: "HTML"})
		return nil
	}
	defer func() {
		// keep cache like Python downloads/ — do not always delete
		_ = file
	}()

	caption := fmt.Sprintf("%s <b>%s</b>\n⏱ %s\n👤 %s\n\n<code>VC stream (ntgcalls) — next layer</code>", kind, escapeHTML(track.Title), track.Duration, escapeHTML(track.RequestedBy))
	_, sendErr := b.Client.SendMedia(chatID, file, &telegram.MediaOptions{
		Caption:   caption,
		ParseMode: "HTML",
	})
	if sendErr != nil {
		_, _ = status.Edit("❌ Send failed: "+escapeHTML(sendErr.Error()), &telegram.SendOptions{ParseMode: "HTML"})
		return nil
	}
	_, _ = status.Edit(fmt.Sprintf("%s <b>Playing</b>\n\n<b>%s</b>\n⏱ %s\n👤 %s\n📎 File sent (API → yt-dlp same as Python)", kind, escapeHTML(track.Title), track.Duration, escapeHTML(track.RequestedBy)), &telegram.SendOptions{ParseMode: "HTML"})
	_ = os.Getenv // keep imports honest if needed
	return nil
}

func escapeHTML(s string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
	)
	return replacer.Replace(s)
}
