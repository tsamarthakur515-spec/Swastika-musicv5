package plugins

import (
	"fmt"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/admin"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/emojis"
)

func registerModeration(b *bot.Bot) {
	b.On("message:/ban", func(m *telegram.NewMessage) error {
		return modAction(b, m, "ban")
	})
	b.On("message:/unban", func(m *telegram.NewMessage) error {
		return modAction(b, m, "unban")
	})
	b.On("message:/mute", func(m *telegram.NewMessage) error {
		return modAction(b, m, "mute")
	})
	b.On("message:/unmute", func(m *telegram.NewMessage) error {
		return modAction(b, m, "unmute")
	})
	b.On("message:/kick", func(m *telegram.NewMessage) error {
		return modAction(b, m, "kick")
	})
	b.On("message:/promote", func(m *telegram.NewMessage) error {
		return modAction(b, m, "promote")
	})
	b.On("message:/demote", func(m *telegram.NewMessage) error {
		return modAction(b, m, "demote")
	})
}

func modAction(b *bot.Bot, m *telegram.NewMessage, action string) error {
	uid, err := admin.TargetFromReply(m)
	if err != nil {
		_, e := m.Reply(emojis.Warn()+" Reply to a user's message.", &telegram.SendOptions{ParseMode: "HTML"})
		return e
	}
	chatID := m.ChatID()
	var actErr error
	switch action {
	case "ban":
		actErr = admin.Ban(b.Client, chatID, uid)
	case "unban":
		actErr = admin.Unban(b.Client, chatID, uid)
	case "mute":
		actErr = admin.Mute(b.Client, chatID, uid)
	case "unmute":
		actErr = admin.Unmute(b.Client, chatID, uid)
	case "kick":
		actErr = admin.Kick(b.Client, chatID, uid)
	case "promote":
		actErr = admin.Promote(b.Client, chatID, uid)
	case "demote":
		actErr = admin.Demote(b.Client, chatID, uid)
	}
	if actErr != nil {
		_, e := m.Reply(fmt.Sprintf("%s <b>%s failed</b>\n<code>%s</code>\nBot needs admin rights.", emojis.Warn(), action, escapeHTML(actErr.Error())), &telegram.SendOptions{ParseMode: "HTML"})
		return e
	}
	_, e := m.Reply(fmt.Sprintf("%s <b>%s</b> done for <code>%d</code>", emojis.CheckMark(), action, uid), &telegram.SendOptions{ParseMode: "HTML"})
	return e
}
