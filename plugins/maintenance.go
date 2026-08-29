package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerMaintenance(b *bot.Bot) {
	b.On("message:/maintenance", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🛠 <b>Maintenance mode</b>\n\n<code>/maintenance on|off</code> — disable public commands temporarily.", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
