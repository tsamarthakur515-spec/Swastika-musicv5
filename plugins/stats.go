package plugins

import (
	"fmt"
	"runtime"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func registerStats(b *bot.Bot) {
	b.On("message:/stats", func(m *telegram.NewMessage) error {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		text := fmt.Sprintf(`📊 <b>Swastika Music V5 Stats</b>

🐹 Go %s
🧵 Goroutines: %d
💾 Alloc: %.2f MB
📦 Sys: %.2f MB

<code>Full DB stats (users/chats/plays) after database layer.</code>`,
			runtime.Version(),
			runtime.NumGoroutine(),
			float64(mem.Alloc)/1024/1024,
			float64(mem.Sys)/1024/1024,
		)
		_, err := m.Reply(text, &telegram.SendOptions{ParseMode: "HTML"})
		return err
	})
}
