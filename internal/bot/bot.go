package bot

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/config"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/queue"
)

type Bot struct {
	Client *telegram.Client
	Cfg    *config.Config
	Queue  *queue.Manager
}

func New(client *telegram.Client, cfg *config.Config) *Bot {
	return &Bot{
		Client: client,
		Cfg:    cfg,
		Queue:  queue.NewManager(),
	}
}

func (b *Bot) On(pattern string, handler func(m *telegram.NewMessage) error) {
	b.Client.On(pattern, handler)
}
