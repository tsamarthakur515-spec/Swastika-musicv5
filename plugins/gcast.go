package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerGcast(b *bot.Bot) {
	b.On("message:/gcast", func(m *telegram.NewMessage) error {
		_, err := m.Reply("📣 <b>Broadcast</b>\n\nOwner-only global cast to all chats.\nRequires chat list from DB.", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
	b.On("message:/broadcast", func(m *telegram.NewMessage) error {
		_, err := m.Reply("📣 Broadcast alias (sudo only, WIP).")
		return err
	})
}
