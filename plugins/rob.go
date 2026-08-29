package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerRob(b *bot.Bot) {
	b.On("message:/rob", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🥷 <b>Rob</b>\n\nFun economy/rob command from Python port (WIP).", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
