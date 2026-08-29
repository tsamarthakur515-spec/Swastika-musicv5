package plugins

import (
	"fmt"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/player"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/queue"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/platforms"
)

var mediaPlayer = player.New()

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

	if pos == 1 {
		text := fmt.Sprintf("%s <b>Now playing</b>\n\n<b>%s</b>\n⏱ %s\n👤 %s\n\n⬇️ Downloading audio...", kind, escapeHTML(track.Title), track.Duration, escapeHTML(track.RequestedBy))
		_, _ = status.Edit(text, &telegram.SendOptions{ParseMode: "HTML"})

		// Deliver audio file to chat (VC WebRTC engine next)
		if err := mediaPlayer.PlayTrack(b.Client, chatID, track); err != nil {
			_, _ = status.Edit(fmt.Sprintf("%s <b>Queued</b> (download failed)\n\n<b>%s</b>\n⏱ %s\n\n<code>%s</code>", kind, escapeHTML(track.Title), track.Duration, escapeHTML(err.Error())), &telegram.SendOptions{ParseMode: "HTML"})
			return nil
		}
		_, _ = status.Edit(fmt.Sprintf("%s <b>Playing</b>\n\n<b>%s</b>\n⏱ %s\n👤 %s\n\n📎 Audio sent to chat.\n<code>Voice Chat stream engine coming next.</code>", kind, escapeHTML(track.Title), track.Duration, escapeHTML(track.RequestedBy)), &telegram.SendOptions{ParseMode: "HTML"})
		return nil
	}

	text := fmt.Sprintf("➕ <b>Added to queue</b> (#%d)\n\n<b>%s</b>\n⏱ %s\n👤 %s", pos, escapeHTML(track.Title), track.Duration, escapeHTML(track.RequestedBy))
	_, _ = status.Edit(text, &telegram.SendOptions{ParseMode: "HTML"})
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
