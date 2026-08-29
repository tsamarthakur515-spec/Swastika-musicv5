package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerGuessGame(b *bot.Bot) {
	b.On("message:/guess", func(m *telegram.NewMessage) error {
		_, err := m.Reply("❓ <b>Guess Game</b>\n\nStart a number/word guess session.\n<code>State machine from Python guessgame.py will be ported.</code>", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
