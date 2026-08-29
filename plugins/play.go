package plugins

import (
	"fmt"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/queue"
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

	chatID := m.ChatID()
	title := args
	url := args

	// Simple detection: if it looks like a URL keep it, else treat as search query
	if !strings.HasPrefix(args, "http://") && !strings.HasPrefix(args, "https://") {
		title = args
		url = "ytsearch:" + args // placeholder for yt-dlp later
	}

	track := queue.Track{
		Title:       title,
		URL:         url,
		RequestedBy: m.Sender.FirstName,
		IsVideo:     video,
	}

	pos := b.Queue.Add(chatID, track)

	if pos == 1 {
		kind := "🎵"
		if video {
			kind = "🎬"
		}
		text := fmt.Sprintf("%s <b>Now playing</b>\n\n<b>%s</b>\n👤 %s\n\n<code>VC stream engine coming soon — track queued &amp; ready.</code>",
			kind, escapeHTML(track.Title), escapeHTML(track.RequestedBy))
		_, err := m.Reply(text, &telegram.SendOptions{ParseMode: "HTML"})
		return err
	}

	text := fmt.Sprintf("➕ <b>Added to queue</b> (#%d)\n\n<b>%s</b>\n👤 %s",
		pos, escapeHTML(track.Title), escapeHTML(track.RequestedBy))
	_, err := m.Reply(text, &telegram.SendOptions{ParseMode: "HTML"})
	return err
}

func escapeHTML(s string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
	)
	return replacer.Replace(s)
}
