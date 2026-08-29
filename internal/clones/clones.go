package clones

// Clone bot spawner — port from Python PANDAMUSIC/modules/clones.py + plugins/clone.py

import (
	"sync"

	"github.com/amarnathcjd/gogram/telegram"
)

type Instance struct {
	Token  string
	Client *telegram.Client
	Owner  int64
}

type Manager struct {
	mu        sync.Mutex
	instances map[string]*Instance
}

func New() *Manager {
	return &Manager{instances: make(map[string]*Instance)}
}

func (m *Manager) List() []*Instance {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]*Instance, 0, len(m.instances))
	for _, v := range m.instances {
		out = append(out, v)
	}
	return out
}

func (m *Manager) Count() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.instances)
}

// TODO: start/stop clone from bot token, isolate queues per clone
