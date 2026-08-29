package plugins

import (
	"fmt"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerControls(b *bot.Bot) {
	b.On("message:/pause", func(m *telegram.NewMessage) error {
		ok := b.Queue.SetPaused(m.ChatID(), true)
		if !ok {
			_, err := m.Reply("❌ Queue is empty.")
			return err
		}
		_, err := m.Reply("⏸ Paused.")
		return err
	})

	b.On("message:/resume", func(m *telegram.NewMessage) error {
		ok := b.Queue.SetPaused(m.ChatID(), false)
		if !ok {
			_, err := m.Reply("❌ Queue is empty.")
			return err
		}
		_, err := m.Reply("▶️ Resumed.")
		return err
	})

	b.On("message:/skip", func(m *telegram.NewMessage) error {
		next := b.Queue.Skip(m.ChatID())
		if next == nil {
			_, err := m.Reply("⏭ Skipped. Queue is now empty.")
			return err
		}
		_, err := m.Reply(fmt.Sprintf("⏭ Skipped.\n🎵 Now: <b>%s</b>", escapeHTML(next.Title)), &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})

	b.On("message:/stop", func(m *telegram.NewMessage) error {
		b.Queue.Clear(m.ChatID())
		_, err := m.Reply("⏹ Stopped. Queue cleared.")
		return err
	})

	b.On("message:/end", func(m *telegram.NewMessage) error {
		b.Queue.Clear(m.ChatID())
		_, err := m.Reply("⏹ Stopped. Queue cleared.")
		return err
	})

	b.On("message:/queue", func(m *telegram.NewMessage) error {
		list := b.Queue.List(m.ChatID())
		if len(list) == 0 {
			_, err := m.Reply("📭 Queue is empty.")
			return err
		}

		var sb strings.Builder
		sb.WriteString("📜 <b>Current Queue</b>\n\n")
		for i, t := range list {
			prefix := ""
			if i == 0 {
				prefix = "▶️ "
			} else {
				prefix = fmt.Sprintf("%d. ", i+1)
			}
			kind := "🎵"
			if t.IsVideo {
				kind = "🎬"
			}
			sb.WriteString(fmt.Sprintf("%s%s <b>%s</b>\n   👤 %s\n", prefix, kind, escapeHTML(t.Title), escapeHTML(t.RequestedBy)))
		}
		if b.Queue.IsPaused(m.ChatID()) {
			sb.WriteString("\n⏸ <i>Paused</i>")
		}
		_, err := m.Reply(sb.String(), &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
