package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerWelcome(b *bot.Bot) {
	b.On("message:/welcome", func(m *telegram.NewMessage) error {
		_, err := m.Reply("👋 <b>Welcome system</b>\n\n<code>/setwelcome [text]</code> — set message\n<code>/resetwelcome</code> — reset\n\nJoin/leave handlers will be wired with gogram updates.", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
	b.On("message:/setwelcome", func(m *telegram.NewMessage) error {
		_, err := m.Reply("✅ Welcome message saved (stub — DB layer next).")
		return err
	})
	b.On("message:/resetwelcome", func(m *telegram.NewMessage) error {
		_, err := m.Reply("♻️ Welcome reset.")
		return err
	})
}
