package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerReload(b *bot.Bot) {
	b.On("message:/reload", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🔄 Admin cache / config reload (stub).")
		return err
	})
	b.On("message:/restart", func(m *telegram.NewMessage) error {
		_, err := m.Reply("♻️ Restart requested — process manager / Railway redeploy handles this.")
		return err
	})
}
