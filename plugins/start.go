package plugins

import (
	"fmt"
	"time"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/emojis"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/formatters"
)

var startedAt = time.Now()

func registerStart(b *bot.Bot) {
	b.On("message:/start", func(m *telegram.NewMessage) error {
		me, _ := b.Client.GetMe()
		botUser := "swastika_music_bot"
		if me != nil && me.Username != "" {
			botUser = me.Username
		}
		mention := "User"
		if m.Sender != nil {
			mention = safeMention(m.Sender.FirstName, m.Sender.ID)
		}
		cap := startCaption(mention)
		mk := startMarkup(botUser, b.Cfg.OwnerUsername, b.Cfg.SupportChat, b.Cfg.SupportChannel, b.Cfg.OwnerID)

		if b.Cfg.StartImageURL != "" {
			_, err := b.Client.SendMedia(m.ChatID(), b.Cfg.StartImageURL, &telegram.MediaOptions{
				Caption:     cap,
				ParseMode:   "HTML",
				ReplyMarkup: mk,
			})
			if err == nil {
				return nil
			}
		}
		_, err := m.Reply(cap, &telegram.SendOptions{ParseMode: "HTML", ReplyMarkup: mk})
		return err
	})

	b.On("message:/help", func(m *telegram.NewMessage) error {
		_, err := m.Reply(helpListCaption(), &telegram.SendOptions{
			ParseMode:   "HTML",
			ReplyMarkup: helpMenuMarkup(),
		})
		return err
	})

	b.On("message:/ping", func(m *telegram.NewMessage) error {
		start := time.Now()
		msg, err := m.Reply(emojis.Spark() + " " + formatters.Smallcaps("pong..."))
		if err != nil {
			return err
		}
		ms := time.Since(start).Milliseconds()
		up := time.Since(startedAt).Truncate(time.Second)
		text := fmt.Sprintf("%s <b>%s</b>\n%s %d ms\n%s %s",
			emojis.TG(emojis.CEPingTitle, "⚡"), formatters.Smallcaps("pong"),
			emojis.TG(emojis.CEPingMS, "✨"), ms,
			emojis.TG(emojis.CEPingUptime, "⏱"), formatters.Smallcaps(up.String()),
		)
		_, err = msg.Edit(text, &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
