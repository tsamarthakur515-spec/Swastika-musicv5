package clients

// Multi-client manager — port from Python PANDAMUSIC/modules/clients.py
// Holds bot client + optional assistant userbot(s) for VC join.

import (
	"github.com/amarnathcjd/gogram/telegram"
)

type Manager struct {
	Bot       *telegram.Client
	Assistants []*telegram.Client
}

func New(bot *telegram.Client) *Manager {
	return &Manager{Bot: bot}
}

func (m *Manager) AddAssistant(c *telegram.Client) {
	m.Assistants = append(m.Assistants, c)
}

// TODO: login assistants via STRING_SESSION, join VC, leave, flood-wait handling
