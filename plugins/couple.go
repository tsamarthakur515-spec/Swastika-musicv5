package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerCouple(b *bot.Bot) {
	b.On("message:/couple", func(m *telegram.NewMessage) error {
		_, err := m.Reply("💕 <b>Couple of the day</b>\n\nPicks two random members.\nNeeds chat member list + daily storage.", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
