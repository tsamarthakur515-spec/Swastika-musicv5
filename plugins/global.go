package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerGlobal(b *bot.Bot) {
	b.On("message:/gban", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🌍 Global ban — sudo only (WIP).")
		return err
	})
	b.On("message:/ungban", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🌍 Global unban (WIP).")
		return err
	})
}
