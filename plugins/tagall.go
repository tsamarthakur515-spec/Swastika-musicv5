package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerTagAll(b *bot.Bot) {
	b.On("message:/tagall", func(m *telegram.NewMessage) error {
		_, err := m.Reply("📢 <b>Tag All</b>\n\nMentions members in batches.\n<code>Needs admin rights + participant fetch via gogram.</code>", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
	b.On("message:/all", func(m *telegram.NewMessage) error {
		_, err := m.Reply("📢 Tag-all alias — same as /tagall (WIP).")
		return err
	})
}
