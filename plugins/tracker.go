package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerTracker(b *bot.Bot) {
	b.On("message:/track", func(m *telegram.NewMessage) error {
		_, err := m.Reply("📍 Tracker / analytics stub.")
		return err
	})
}
