package plugins

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/emojis"
)

func registerCallbacks(b *bot.Bot) {
	b.Client.On("callback:pl:", func(c *telegram.CallbackQuery) error {
		data := c.DataString()
		// pl:pause:<chatID> | pl:resume: | pl:skip: | pl:stop: | pl:queue:
		parts := strings.Split(data, ":")
		if len(parts) < 3 {
			_ = c.Answer("Invalid")
			return nil
		}
		action := parts[1]
		chatID, _ := strconv.ParseInt(parts[2], 10, 64)
		if chatID == 0 {
			chatID = c.ChatID
		}

		switch action {
		case "pause":
			if !b.Queue.SetPaused(chatID, true) {
				_ = c.Answer("Nothing playing", &telegram.CallbackOptions{Alert: true})
				return nil
			}
			if b.Stream != nil {
				_ = b.Stream.Pause(chatID)
			}
			_ = c.Answer("Paused ⏸")
			_, _ = c.Edit(emojis.Pause()+" <b>Paused</b>", &telegram.SendOptions{ParseMode: "HTML"})
		case "resume":
			if !b.Queue.SetPaused(chatID, false) {
				_ = c.Answer("Nothing to resume", &telegram.CallbackOptions{Alert: true})
				return nil
			}
			if b.Stream != nil {
				_ = b.Stream.Resume(chatID)
			}
			_ = c.Answer("Resumed ▶️")
			_, _ = c.Edit(emojis.Play()+" <b>Resumed</b>", &telegram.SendOptions{ParseMode: "HTML"})
		case "skip":
			next := b.Queue.Skip(chatID)
			if next == nil {
				if b.Stream != nil {
					_ = b.Stream.Stop(chatID)
				}
				_ = c.Answer("Queue empty")
				_, _ = c.Edit(emojis.Stop()+" Queue finished", &telegram.SendOptions{ParseMode: "HTML"})
				return nil
			}
			_ = c.Answer("Skipped ⏭")
			_, _ = c.Edit(fmt.Sprintf("%s <b>Skipped</b> → <b>%s</b>", emojis.Skip(), escapeHTML(next.Title)), &telegram.SendOptions{ParseMode: "HTML", ReplyMarkup: playerKeyboard(chatID)})
		case "stop":
			b.Queue.Clear(chatID)
			if b.Stream != nil {
				_ = b.Stream.Stop(chatID)
			}
			_ = c.Answer("Stopped")
			_, _ = c.Edit(emojis.Stop()+" <b>Stopped</b>", &telegram.SendOptions{ParseMode: "HTML"})
		case "queue":
			list := b.Queue.List(chatID)
			if len(list) == 0 {
				_ = c.Answer("Empty queue", &telegram.CallbackOptions{Alert: true})
				return nil
			}
			text := emojis.Queue() + " <b>Queue</b>\n\n"
			for i, t := range list {
				if i >= 10 {
					text += fmt.Sprintf("\n… +%d more", len(list)-10)
					break
				}
				text += fmt.Sprintf("%d. <b>%s</b>\n", i+1, escapeHTML(t.Title))
			}
			_ = c.Answer("Queue")
			_, _ = c.Edit(text, &telegram.SendOptions{ParseMode: "HTML", ReplyMarkup: playerKeyboard(chatID)})
		default:
			_ = c.Answer("?")
		}
		return nil
	})
}

func playerKeyboard(chatID int64) *telegram.Keyboard {
	id := fmt.Sprintf("%d", chatID)
	kb := telegram.NewKeyboard()
	kb.AddRow(
		telegram.Button.Data("⏸ Pause", "pl:pause:"+id),
		telegram.Button.Data("▶️ Resume", "pl:resume:"+id),
	)
	kb.AddRow(
		telegram.Button.Data("⏭ Skip", "pl:skip:"+id),
		telegram.Button.Data("⏹ Stop", "pl:stop:"+id),
	)
	kb.AddRow(
		telegram.Button.Data("📜 Queue", "pl:queue:"+id),
	)
	return kb.Build()
}
