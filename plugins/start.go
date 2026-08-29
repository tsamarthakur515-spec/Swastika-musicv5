package plugins

import (
	"fmt"
	"time"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerStart(b *bot.Bot) {
	b.On("message:/start", func(m *telegram.NewMessage) error {
		text := `🎵 <b>Swastika Music V5</b>
<code>Go + gogram</code>

Fast Telegram music bot.

<b>Commands</b>
/play &lt;song or url&gt; — Play / queue
/vplay &lt;query&gt; — Video play
/pause — Pause
/resume — Resume
/skip — Skip
/stop — Stop &amp; clear
/queue — Show queue
/ping — Latency
/help — This help

Voice Chat streaming layer coming next.`
		_, err := m.Reply(text, &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})

	b.On("message:/help", func(m *telegram.NewMessage) error {
		_, err := m.Reply(`🎵 <b>Swastika Music V5 — Help</b>

<b>Playback</b>
/play [query|url] — Search &amp; play
/vplay [query|url] — Play as video
/pause / /resume
/skip / /stop
/queue

<b>Other</b>
/ping / /start / /help`, &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})

	b.On("message:/ping", func(m *telegram.NewMessage) error {
		start := time.Now()
		msg, err := m.Reply("🏓 Pinging...")
		if err != nil {
			return err
		}
		latency := time.Since(start).Milliseconds()
		_, err = msg.Edit(fmt.Sprintf("🏓 <b>Pong!</b>\n⏱ %d ms\n🤖 Go + gogram", latency), &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
