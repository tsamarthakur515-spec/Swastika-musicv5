package plugins

import (
	"fmt"
	"html"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/emojis"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/formatters"
)

// Extra start-menu emoji IDs from Python start.py
const (
	eTitle       = "6111778259374971023"
	eStart       = "6147614817952735246"
	eHelpInside  = "6154314112236001069"
	eCmdBtn      = "5823571441118876120"
	eStartSticker = "6154300385520522693"
)

func sc(s string) string { return formatters.Smallcaps(s) }

func startCaption(mention string) string {
	e := emojis.TG(eTitle, "✨")
	return fmt.Sprintf(
		`%s <b>𝗦𝘄𝗮𝘀𝘁𝗶𝗸𝗮 𝗠𝘂𝘀𝗶𝗰 𝘃𝟱</b>

%s %s %s

%s %s
%s %s

%s %s`,
		e,
		e, sc("hey"), mention,
		e, sc("i am a high quality fast music bot."),
		e, sc("add me to your group and enjoy audio / video streaming."),
		e, sc("use the buttons below for help, owner and support."),
	)
}

func helpListCaption() string {
	body := fmt.Sprintf(
		`%s

%s

• music — play vplay pause resume skip end
• tools — stats active welcome broadcast
• moderation — mute ban kick tagall noabuse
• chatbot — chaton chatoff
• locks — lock unlock content types
• games — economy rpg social
• fun — couple riddle dice slots coinflip`,
		sc("help menu"),
		sc("tap a category below to see commands."),
	)
	return fmt.Sprintf(`<blockquote expandable>%s %s</blockquote>`, emojis.GearE(), body)
}

func musicListCaption() string {
	return fmt.Sprintf(`<blockquote expandable>%s %s

%s
%s</blockquote>`,
		emojis.TG(emojis.Lightning, "⚡"), sc("music commands"),
		sc("audio and video streaming controls."),
		sc("tap a button to see usage."),
	)
}

func toolsListCaption() string {
	return fmt.Sprintf(`<blockquote expandable>%s %s

%s
%s</blockquote>`,
		emojis.GearE(), sc("tools"),
		sc("stats, active chats, welcome and broadcast."),
		sc("tap a button to see usage."),
	)
}

func moderationListCaption() string {
	return fmt.Sprintf(`<blockquote expandable>%s %s

%s
%s
%s</blockquote>`,
		emojis.FireE(), sc("moderation"),
		sc("mute ban kick tagall and abuse filter."),
		sc("admin only for most commands."),
		sc("tap a button to see usage."),
	)
}

func chatbotListCaption() string {
	return fmt.Sprintf(`<blockquote expandable>%s %s

%s
%s</blockquote>`,
		emojis.TG(emojis.Lightning, "⚡"), sc("chatbot commands"),
		sc("enable or disable ai chat in this chat."),
		sc("tap a button to see usage."),
	)
}

func locksListCaption() string {
	body := fmt.Sprintf(
		`%s

%s
%s: /lock url
%s

%s: all url photo video document sticker gif voice videonote audio contact location poll game forward bot command text invitelink phone email emoji media

%s`,
		sc("locks"),
		sc("lock content types in your group."),
		sc("example"),
		sc("when locked, non-admin messages are auto-deleted."),
		sc("types"),
		sc("admin owner sudo only. bot needs delete messages right."),
	)
	return fmt.Sprintf(`<blockquote expandable>%s %s</blockquote>`, emojis.GearE(), body)
}

func funListCaption() string {
	return fmt.Sprintf(`<blockquote expandable>%s %s

%s
%s</blockquote>`,
		emojis.Request(), sc("fun"),
		sc("couple of the day, riddles, dice, slots and coinflip."),
		sc("tap a button to see usage."),
	)
}

func aboutCaption() string {
	body := fmt.Sprintf(
		`%s

%s
%s
%s

%s`,
		sc("about"),
		sc("high quality telegram music bot."),
		sc("supports audio and video streaming."),
		sc("powered by gogram + gortc."),
		sc("add me in your group and start playing."),
	)
	return fmt.Sprintf(`<blockquote expandable>%s %s</blockquote>`, emojis.CheckMark(), body)
}

var cmdUsage = map[string]string{
	"play":    sc("command") + ": /play\n\n" + sc("use") + ":\n• /play " + sc("song name") + "\n• /play " + sc("youtube link") + "\n• " + sc("reply to audio with") + " /play\n\n" + sc("plays audio in voice chat."),
	"vplay":   sc("command") + ": /vplay\n\n" + sc("use") + ":\n• /vplay " + sc("song name") + "\n• /vplay " + sc("youtube link") + "\n• " + sc("reply to video with") + " /vplay\n\n" + sc("plays video in voice chat."),
	"pause":   sc("command") + ": /pause\n\n" + sc("use") + ": /pause\n\n" + sc("pauses the current stream."),
	"resume":  sc("command") + ": /resume\n\n" + sc("use") + ": /resume\n\n" + sc("resumes the paused stream."),
	"skip":    sc("command") + ": /skip\n\n" + sc("use") + ": /skip\n\n" + sc("skips to the next track in queue."),
	"end":     sc("command") + ": /end\n\n" + sc("use") + ": /end\n\n" + sc("stops streaming and clears the queue."),
	"stats":   sc("command") + ": /stats\n\n" + sc("use") + ": /stats\n\n" + sc("shows system and bot statistics."),
	"active":  sc("command") + ": /active\n\n" + sc("use") + ": /active\n\n" + sc("shows active voice chats. (owner only)"),
	"broadcast": sc("command") + ": /broadcast\n\n" + sc("use") + ": /broadcast " + sc("message") + "\n\n" + sc("broadcasts message to served users/chats. (owner only)"),
	"couple":  sc("command") + ": /couple\n\n" + sc("use") + ": /couple\n\n" + sc("picks 2 random members and makes couple of the day image."),
	"mute":    sc("command") + ": /mute\n\n" + sc("use") + ":\n• " + sc("reply to user") + ": /mute\n\n" + sc("mutes a user in the group. (admin only)"),
	"unmute":  sc("command") + ": /unmute\n\n" + sc("use") + ":\n• " + sc("reply to user") + ": /unmute\n\n" + sc("unmutes a user in the group. (admin only)"),
	"ban":     sc("command") + ": /ban\n\n" + sc("use") + ":\n• " + sc("reply to user") + ": /ban\n\n" + sc("bans a user from the group. (admin only)"),
	"unban":   sc("command") + ": /unban\n\n" + sc("use") + ":\n• " + sc("reply to user") + ": /unban\n\n" + sc("unbans a user in the group. (admin only)"),
	"kick":    sc("command") + ": /kick\n\n" + sc("use") + ":\n• " + sc("reply to user") + ": /kick\n\n" + sc("kicks a user from the group. (admin only)"),
	"tagall":  sc("command") + ": /tagall\n\n" + sc("use") + ": /tagall " + sc("your message") + "\n\n" + sc("tags all group members."),
	"chaton":  sc("command") + ": /chaton\n\n" + sc("use") + ": /chaton\n\n" + sc("enables chatbot in this chat."),
	"chatoff": sc("command") + ": /chatoff\n\n" + sc("use") + ": /chatoff\n\n" + sc("disables chatbot in this chat."),
	"noabuse": sc("command") + ": /noabuse\n\n" + sc("use") + ":\n• /noabuse on\n• /noabuse off\n\n" + sc("auto deletes abusive messages in group."),
	"welcome": sc("command") + ": /welcome\n\n" + sc("use") + ":\n• /welcome on\n• /welcome off\n\n" + sc("enable or disable welcome messages."),
	"setwelcome": sc("command") + ": /setwelcome\n\n" + sc("use") + ": /setwelcome " + sc("text") + "\n\n" + sc("sets custom welcome message."),
	"resetwelcome": sc("command") + ": /resetwelcome\n\n" + sc("use") + ": /resetwelcome\n\n" + sc("resets welcome message to default."),
	"lock":    sc("command") + ": /lock\n\n" + sc("use") + ": /lock url\n\n" + sc("locks a content type in the group."),
	"unlock":  sc("command") + ": /unlock\n\n" + sc("use") + ": /unlock url\n\n" + sc("unlocks a content type."),
	"locks":   sc("command") + ": /locks\n\n" + sc("use") + ": /locks\n\n" + sc("shows all active locks in this group."),
	"riddle":  sc("command") + ": /riddle\n\n" + sc("use") + ": /riddle\n\n" + sc("sends a random riddle quiz."),
	"dice":    sc("command") + ": /dice\n\n" + sc("use") + ": /dice\n\n" + sc("rolls a dice."),
	"slots":   sc("command") + ": /slots\n\n" + sc("use") + ": /slots\n\n" + sc("virtual slot machine."),
	"coinflip": sc("command") + ": /coinflip\n\n" + sc("use") + ": /coinflip\n\n" + sc("heads or tails."),
}

func cmdUsageCaption(key string) string {
	u, ok := cmdUsage[key]
	if !ok {
		u = sc("unknown command")
	}
	return fmt.Sprintf(`<blockquote expandable>%s %s</blockquote>`, emojis.Title(), u)
}

func addGroupURL(botUsername string) string {
	admin := "delete_messages+manage_video_chats+pin_messages+invite_users+ban_users+change_info+promote_members"
	return fmt.Sprintf("https://t.me/%s?startgroup=s&admin=%s", botUsername, admin)
}

func startMarkup(botUsername, owner, support, channel string, ownerID int64) *telegram.Keyboard {
	kb := telegram.NewKeyboard()
	kb.AddRow(telegram.Button.URL(sc("➕ add me in your group ➕"), addGroupURL(botUsername)))

	var ownerBtn telegram.Button
	if owner != "" {
		ownerBtn = telegram.Button.URL(sc("owner"), "https://t.me/"+owner)
	} else if ownerID != 0 {
		ownerBtn = telegram.Button.URL(sc("owner"), fmt.Sprintf("tg://user?id=%d", ownerID))
	} else {
		ownerBtn = telegram.Button.Data(sc("owner"), "about_menu")
	}
	kb.AddRow(ownerBtn, telegram.Button.Data(sc("about"), "about_menu"))

	var supportBtn, updateBtn telegram.Button
	if support != "" {
		supportBtn = telegram.Button.URL(sc("support"), "https://t.me/"+support)
	} else {
		supportBtn = telegram.Button.Data(sc("support"), "support_alert")
	}
	if channel != "" {
		updateBtn = telegram.Button.URL(sc("update"), "https://t.me/"+channel)
	} else {
		updateBtn = telegram.Button.Data(sc("update"), "update_alert")
	}
	kb.AddRow(supportBtn, updateBtn)
	kb.AddRow(telegram.Button.Data(sc("help and commands"), "help_menu"))
	kb.AddRow(telegram.Button.Data(sc("source"), "repo_alert"))
	return kb.Build()
}

func helpMenuMarkup() *telegram.Keyboard {
	kb := telegram.NewKeyboard()
	kb.AddRow(
		telegram.Button.Data(sc("MUSIC"), "music_menu"),
		telegram.Button.Data(sc("TOOLS"), "tools_menu"),
	)
	kb.AddRow(
		telegram.Button.Data(sc("MODERATION"), "moderation_menu"),
		telegram.Button.Data(sc("CHATBOT"), "chatbot_menu"),
	)
	kb.AddRow(
		telegram.Button.Data(sc("LOCKS"), "locks_menu"),
		telegram.Button.Data(sc("GAMES"), "games_menu"),
	)
	kb.AddRow(telegram.Button.Data(sc("FUN"), "fun_menu"))
	kb.AddRow(telegram.Button.Data(sc("« BACK"), "home_menu"))
	return kb.Build()
}

func cmdRows(keys []string) *telegram.Keyboard {
	kb := telegram.NewKeyboard()
	for i := 0; i < len(keys); i += 2 {
		if i+1 < len(keys) {
			kb.AddRow(
				telegram.Button.Data(sc(keys[i]), "cmdhelp|"+keys[i]),
				telegram.Button.Data(sc(keys[i+1]), "cmdhelp|"+keys[i+1]),
			)
		} else {
			kb.AddRow(telegram.Button.Data(sc(keys[i]), "cmdhelp|"+keys[i]))
		}
	}
	kb.AddRow(
		telegram.Button.Data(sc("COMMANDS"), "help_menu"),
		telegram.Button.Data(sc("« BACK"), "home_menu"),
	)
	return kb.Build()
}

func safeMention(first string, id int64) string {
	name := html.EscapeString(first)
	if name == "" {
		name = "User"
	}
	return fmt.Sprintf(`<a href="tg://user?id=%d">%s</a>`, id, name)
}
