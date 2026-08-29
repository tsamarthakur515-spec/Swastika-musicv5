package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerNoAbuse(b *bot.Bot) {
	b.On("message:/noabuse", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🛡 <b>No-Abuse</b> filter\n\nToggles anti-abuse word filter for this chat.\n<code>Full word-list + auto-action logic coming from Python port.</code>", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
