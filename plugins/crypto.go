package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerCrypto(b *bot.Bot) {
	b.On("message:/crypto", func(m *telegram.NewMessage) error {
		_, err := m.Reply("₿ <b>Crypto</b>\n\n<code>/crypto btc</code> — price lookup\nAPI integration from Python crypto plugin.", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
	b.On("message:/price", func(m *telegram.NewMessage) error {
		_, err := m.Reply("💰 Price lookup — specify coin (e.g. /price btc)")
		return err
	})
}
