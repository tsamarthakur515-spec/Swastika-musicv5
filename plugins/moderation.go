package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerModeration(b *bot.Bot) {
	b.On("message:/ban", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🔨 <b>Ban</b> — reply to a user.\n<code>Full moderation logic porting from Python...</code>", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
	b.On("message:/unban", func(m *telegram.NewMessage) error {
		_, err := m.Reply("✅ Unban — reply to a user. (WIP)")
		return err
	})
	b.On("message:/mute", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🔇 Mute — reply to a user. (WIP)")
		return err
	})
	b.On("message:/unmute", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🔊 Unmute — reply to a user. (WIP)")
		return err
	})
	b.On("message:/kick", func(m *telegram.NewMessage) error {
		_, err := m.Reply("👢 Kick — reply to a user. (WIP)")
		return err
	})
	b.On("message:/promote", func(m *telegram.NewMessage) error {
		_, err := m.Reply("⬆️ Promote — reply to a user. (WIP)")
		return err
	})
	b.On("message:/demote", func(m *telegram.NewMessage) error {
		_, err := m.Reply("⬇️ Demote — reply to a user. (WIP)")
		return err
	})
}
