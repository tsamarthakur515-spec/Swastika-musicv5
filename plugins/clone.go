package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerClone(b *bot.Bot) {
	b.On("message:/clone", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🧬 <b>Clone Bot</b>\n\nSpawn clone instances from bot tokens.\nPython clones.py logic will be ported with multi-client support.", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
	b.On("message:/clones", func(m *telegram.NewMessage) error {
		_, err := m.Reply("📋 Active clones: 0 (manager pending).")
		return err
	})
}
