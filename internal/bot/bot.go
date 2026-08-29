package bot

import (
	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/config"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/queue"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/stream"
)

type Bot struct {
	Client *telegram.Client
	Cfg    *config.Config
	Queue  *queue.Manager
	Stream *stream.Manager
}

func New(client *telegram.Client, cfg *config.Config, streamMgr *stream.Manager) *Bot {
	return &Bot{
		Client: client,
		Cfg:    cfg,
		Queue:  queue.NewManager(),
		Stream: streamMgr,
	}
}

func (b *Bot) On(pattern string, handler func(m *telegram.NewMessage) error) {
	b.Client.On(pattern, handler)
}
