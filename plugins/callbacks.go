package plugins

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/emojis"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/formatters"
)

func registerCallbacks(b *bot.Bot) {
	// Menu navigation (exact Python callback names)
	b.Client.On("callback:help_menu", func(c *telegram.CallbackQuery) error {
		_ = c.Answer()
		_, _ = c.Edit(helpListCaption(), &telegram.SendOptions{ParseMode: "HTML", ReplyMarkup: helpMenuMarkup()})
		return nil
	})
	b.Client.On("callback:home_menu", func(c *telegram.CallbackQuery) error {
		_ = c.Answer()
		me, _ := b.Client.GetMe()
		botUser := "bot"
		if me != nil {
			botUser = me.Username
		}
		mention := "User"
		if c.Sender != nil {
			mention = safeMention(c.Sender.FirstName, c.Sender.ID)
		}
		_, _ = c.Edit(startCaption(mention), &telegram.SendOptions{
			ParseMode:   "HTML",
			ReplyMarkup: startMarkup(botUser, b.Cfg.OwnerUsername, b.Cfg.SupportChat, b.Cfg.SupportChannel, b.Cfg.OwnerID),
		})
		return nil
	})
	b.Client.On("callback:about_menu", func(c *telegram.CallbackQuery) error {
		_ = c.Answer()
		kb := telegram.NewKeyboard()
		kb.AddRow(telegram.Button.Data(formatters.Smallcaps("« BACK"), "home_menu"))
		_, _ = c.Edit(aboutCaption(), &telegram.SendOptions{ParseMode: "HTML", ReplyMarkup: kb.Build()})
		return nil
	})
	b.Client.On("callback:music_menu", func(c *telegram.CallbackQuery) error {
		_ = c.Answer()
		_, _ = c.Edit(musicListCaption(), &telegram.SendOptions{
			ParseMode:   "HTML",
			ReplyMarkup: cmdRows([]string{"play", "vplay", "pause", "resume", "skip", "end"}),
		})
		return nil
	})
	b.Client.On("callback:tools_menu", func(c *telegram.CallbackQuery) error {
		_ = c.Answer()
		_, _ = c.Edit(toolsListCaption(), &telegram.SendOptions{
			ParseMode:   "HTML",
			ReplyMarkup: cmdRows([]string{"stats", "active", "welcome", "setwelcome", "resetwelcome", "broadcast"}),
		})
		return nil
	})
	b.Client.On("callback:moderation_menu", func(c *telegram.CallbackQuery) error {
		_ = c.Answer()
		_, _ = c.Edit(moderationListCaption(), &telegram.SendOptions{
			ParseMode:   "HTML",
			ReplyMarkup: cmdRows([]string{"mute", "unmute", "ban", "unban", "kick", "tagall", "noabuse"}),
		})
		return nil
	})
	b.Client.On("callback:chatbot_menu", func(c *telegram.CallbackQuery) error {
		_ = c.Answer()
		_, _ = c.Edit(chatbotListCaption(), &telegram.SendOptions{
			ParseMode:   "HTML",
			ReplyMarkup: cmdRows([]string{"chaton", "chatoff"}),
		})
		return nil
	})
	b.Client.On("callback:locks_menu", func(c *telegram.CallbackQuery) error {
		_ = c.Answer()
		_, _ = c.Edit(locksListCaption(), &telegram.SendOptions{
			ParseMode:   "HTML",
			ReplyMarkup: cmdRows([]string{"lock", "unlock", "locks"}),
		})
		return nil
	})
	b.Client.On("callback:fun_menu", func(c *telegram.CallbackQuery) error {
		_ = c.Answer()
		_, _ = c.Edit(funListCaption(), &telegram.SendOptions{
			ParseMode:   "HTML",
			ReplyMarkup: cmdRows([]string{"couple", "riddle", "dice", "slots", "coinflip"}),
		})
		return nil
	})
	b.Client.On("callback:games_menu", func(c *telegram.CallbackQuery) error {
		_ = c.Answer()
		text := fmt.Sprintf(`<blockquote expandable>%s %s

%s</blockquote>`, emojis.FireE(), formatters.Smallcaps("games"), formatters.Smallcaps("economy rpg social — coming with full port."))
		kb := telegram.NewKeyboard()
		kb.AddRow(telegram.Button.Data(formatters.Smallcaps("« BACK"), "help_menu"))
		_, _ = c.Edit(text, &telegram.SendOptions{ParseMode: "HTML", ReplyMarkup: kb.Build()})
		return nil
	})
	b.Client.On("callback:cmdhelp|", func(c *telegram.CallbackQuery) error {
		data := c.DataString()
		parts := strings.SplitN(data, "|", 2)
		if len(parts) < 2 {
			_ = c.Answer("Invalid", &telegram.CallbackOptions{Alert: true})
			return nil
		}
		key := strings.TrimSpace(strings.ToLower(parts[1]))
		_ = c.Answer()
		kb := telegram.NewKeyboard()
		kb.AddRow(
			telegram.Button.Data(formatters.Smallcaps("COMMANDS"), "help_menu"),
			telegram.Button.Data(formatters.Smallcaps("« BACK"), "home_menu"),
		)
		_, _ = c.Edit(cmdUsageCaption(key), &telegram.SendOptions{ParseMode: "HTML", ReplyMarkup: kb.Build()})
		return nil
	})
	b.Client.On("callback:repo_alert", func(c *telegram.CallbackQuery) error {
		return c.Answer(formatters.Smallcaps("repo private hai")+" 🔒", &telegram.CallbackOptions{Alert: true})
	})
	b.Client.On("callback:support_alert", func(c *telegram.CallbackQuery) error {
		return c.Answer(formatters.Smallcaps("support chat set nahi hai config me"), &telegram.CallbackOptions{Alert: true})
	})
	b.Client.On("callback:update_alert", func(c *telegram.CallbackQuery) error {
		return c.Answer(formatters.Smallcaps("update chat set nahi hai config me"), &telegram.CallbackOptions{Alert: true})
	})

	// Player control callbacks
	b.Client.On("callback:pl:", func(c *telegram.CallbackQuery) error {
		data := c.DataString()
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
			b.Queue.SetPaused(chatID, true)
			if b.Stream != nil {
				_ = b.Stream.Pause(chatID)
			}
			_ = c.Answer(formatters.Smallcaps("paused"))
		case "resume":
			b.Queue.SetPaused(chatID, false)
			if b.Stream != nil {
				_ = b.Stream.Resume(chatID)
			}
			_ = c.Answer(formatters.Smallcaps("resumed"))
		case "skip":
			_ = b.Queue.Skip(chatID)
			_ = c.Answer(formatters.Smallcaps("skipped"))
		case "stop":
			b.Queue.Clear(chatID)
			if b.Stream != nil {
				_ = b.Stream.Stop(chatID)
			}
			_ = c.Answer(formatters.Smallcaps("stopped"))
		case "queue":
			list := b.Queue.List(chatID)
			if len(list) == 0 {
				_ = c.Answer(formatters.Smallcaps("empty queue"), &telegram.CallbackOptions{Alert: true})
				return nil
			}
			_ = c.Answer(formatters.Smallcaps("queue"))
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
		telegram.Button.Data("⏸", "pl:pause:"+id),
		telegram.Button.Data("▶️", "pl:resume:"+id),
		telegram.Button.Data("⏭", "pl:skip:"+id),
		telegram.Button.Data("⏹", "pl:stop:"+id),
	)
	kb.AddRow(telegram.Button.Data("📜 "+formatters.Smallcaps("queue"), "pl:queue:"+id))
	return kb.Build()
}
