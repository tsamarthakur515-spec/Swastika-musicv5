package plugins

import (
	"fmt"
	"time"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/emojis"
)

var startedAt = time.Now()

func registerStart(b *bot.Bot) {
	b.On("message:/start", func(m *telegram.NewMessage) error {
		text := fmt.Sprintf(`%s <b>Swastika Music V5</b> %s

%s Go + gogram music bot

%s <b>Play</b> — <code>/play song</code>
%s <b>Controls</b> — pause · resume · skip · stop
%s <b>Queue</b> — <code>/queue</code>

%s Premium captions · moderation · player buttons
%s VC stream via gortc when <code>STRING_SESSION</code> set`,
			emojis.FireE(), emojis.WolfE(),
			emojis.Spark(),
			emojis.Play(),
			emojis.Pause(),
			emojis.Queue(),
			emojis.Title(),
			emojis.Stream(),
		)
		_, err := m.Reply(text, &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})

	b.On("message:/help", func(m *telegram.NewMessage) error {
		text := emojis.GearE() + ` <b>Commands</b>

<code>/play</code> <code>/vplay</code> <code>/pause</code> <code>/resume</code>
<code>/skip</code> <code>/stop</code> <code>/queue</code>
<code>/ban</code> <code>/mute</code> <code>/kick</code> <code>/promote</code>
<code>/ping</code> <code>/stats</code>`
		_, err := m.Reply(text, &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})

	b.On("message:/ping", func(m *telegram.NewMessage) error {
		start := time.Now()
		msg, err := m.Reply(emojis.Spark() + " Pong...")
		if err != nil {
			return err
		}
		ms := time.Since(start).Milliseconds()
		up := time.Since(startedAt).Truncate(time.Second)
		text := fmt.Sprintf("%s <b>Pong</b>\n%s %d ms\n%s uptime %s",
			emojis.TG(emojis.CEPingTitle, "⚡"),
			emojis.TG(emojis.CEPingMS, "✨"), ms,
			emojis.TG(emojis.CEPingUptime, "⏱"), up,
		)
		_, err = msg.Edit(text, &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
