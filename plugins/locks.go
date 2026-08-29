package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerLocks(b *bot.Bot) {
	b.On("message:/locks", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🔒 <b>Locks</b>\n\nAvailable: media, urls, stickers, gifs, games, inline, all\n<code>/lock [type]</code> · <code>/unlock [type]</code>\n\nFull lock engine porting from Python...", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
	b.On("message:/lock", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🔒 Lock applied (stub). Configure types via /locks.")
		return err
	})
	b.On("message:/unlock", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🔓 Unlocked (stub).")
		return err
	})
}
