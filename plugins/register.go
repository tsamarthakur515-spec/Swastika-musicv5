package plugins

import (
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

func Register(b *bot.Bot) {
	registerStart(b)
	registerPlay(b)
	registerControls(b)
}
