package plugins

import (
	"fmt"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerActive(b *bot.Bot) {
	b.On("message:/active", func(m *telegram.NewMessage) error {
		n := b.Queue.Len(m.ChatID())
		_, err := m.Reply(fmt.Sprintf("📶 This chat queue size: <b>%d</b>", n), &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
	b.On("message:/activevc", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🎙 Active voice chats: tracked after VC engine is live.")
		return err
	})
}
