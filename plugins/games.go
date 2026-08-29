package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerGames(b *bot.Bot) {
	b.On("message:/games", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🎮 <b>Games</b>\n\n/dice · /dart · /basket · /football · /bowling · /slot\nPlus custom games from Python port (truth/dare, etc.).", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
	b.On("message:/dice", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🎲 Rolling... (send Telegram dice via API next)")
		return err
	})
	b.On("message:/truth", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🗣 Truth mode — question bank porting...")
		return err
	})
	b.On("message:/dare", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🔥 Dare mode — challenge bank porting...")
		return err
	})
}
