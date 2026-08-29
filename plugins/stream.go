package plugins

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerStream(b *bot.Bot) {
	b.On("message:/stream", func(m *telegram.NewMessage) error {
		_, err := m.Reply("📡 <b>Stream</b>\n\nRTMP / live stream support is planned.\nUse /play for tracks for now.", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
	b.On("message:/streamstatus", func(m *telegram.NewMessage) error {
		_, err := m.Reply("📡 No active RTMP stream.")
		return err
	})
	b.On("message:/replay", func(m *telegram.NewMessage) error {
		cur := b.Queue.Current(m.ChatID())
		if cur == nil {
			_, err := m.Reply("❌ Nothing to replay.")
			return err
		}
		_, err := m.Reply("🔁 Replaying: <b>"+escapeHTML(cur.Title)+"</b>\n<code>(VC engine pending)</code>", &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
