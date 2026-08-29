package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerChatbot(b *bot.Bot) {
	b.On("message:/chatbot", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🤖 <b>Chatbot</b>\n\n<code>/chatbot on|off</code> — toggle AI replies in this chat.\nPorting conversation logic from Python plugin.", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
