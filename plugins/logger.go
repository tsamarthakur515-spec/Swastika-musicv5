package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerLogger(b *bot.Bot) {
	b.On("message:/logger", func(m *telegram.NewMessage) error {
		_, err := m.Reply("📝 <b>Logger</b>\n\nForwards important events to LOG_GROUP_ID.\nSet LOG_GROUP_ID in env.", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
