package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerSudo(b *bot.Bot) {
	b.On("message:/sudo", func(m *telegram.NewMessage) error {
		_, err := m.Reply("👑 <b>Sudo</b>\n\nOwner/sudo user management.\n<code>/addsudo</code> · <code>/rmsudo</code> · <code>/sudolist</code>", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
	b.On("message:/addsudo", func(m *telegram.NewMessage) error {
		_, err := m.Reply("➕ Add sudo — reply to user (owner only, WIP).")
		return err
	})
	b.On("message:/rmsudo", func(m *telegram.NewMessage) error {
		_, err := m.Reply("➖ Remove sudo (WIP).")
		return err
	})
	b.On("message:/sudolist", func(m *telegram.NewMessage) error {
		_, err := m.Reply("📋 Sudo list — DB pending.")
		return err
	})
}
