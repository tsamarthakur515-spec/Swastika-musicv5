package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerKill(b *bot.Bot) {
	b.On("message:/kill", func(m *telegram.NewMessage) error {
		_, err := m.Reply("💀 <b>Kill</b>\n\nOwner/sudo process control stub.\nOn Railway use dashboard restart instead.", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
