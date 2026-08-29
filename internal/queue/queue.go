package queue

import (
	"sync"
)

type Track struct {
	Title       string
	URL         string
	Duration    string
	Thumbnail   string
	RequestedBy string
	IsVideo     bool
	ID          string
}

type ChatQueue struct {
	mu     sync.Mutex
	Tracks []Track
	Paused bool
}

type Manager struct {
	mu     sync.RWMutex
	queues map[int64]*ChatQueue
}

func NewManager() *Manager {
	return &Manager{
		queues: make(map[int64]*ChatQueue),
	}
}

func (m *Manager) get(chatID int64) *ChatQueue {
	m.mu.Lock()
	defer m.mu.Unlock()
	q, ok := m.queues[chatID]
	if !ok {
		q = &ChatQueue{}
		m.queues[chatID] = q
	}
	return q
}

func (m *Manager) Add(chatID int64, t Track) int {
	q := m.get(chatID)
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Tracks = append(q.Tracks, t)
	return len(q.Tracks)
}

func (m *Manager) Current(chatID int64) *Track {
	q := m.get(chatID)
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.Tracks) == 0 {
		return nil
	}
	return &q.Tracks[0]
}

func (m *Manager) List(chatID int64) []Track {
	q := m.get(chatID)
	q.mu.Lock()
	defer q.mu.Unlock()
	out := make([]Track, len(q.Tracks))
	copy(out, q.Tracks)
	return out
}

func (m *Manager) Skip(chatID int64) *Track {
	q := m.get(chatID)
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.Tracks) == 0 {
		return nil
	}
	q.Tracks = q.Tracks[1:]
	q.Paused = false
	if len(q.Tracks) == 0 {
		return nil
	}
	return &q.Tracks[0]
}

func (m *Manager) Clear(chatID int64) {
	q := m.get(chatID)
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Tracks = nil
	q.Paused = false
}

func (m *Manager) SetPaused(chatID int64, paused bool) bool {
	q := m.get(chatID)
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.Tracks) == 0 {
		return false
	}
	q.Paused = paused
	return true
}

func (m *Manager) IsPaused(chatID int64) bool {
	q := m.get(chatID)
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.Paused
}

func (m *Manager) Len(chatID int64) int {
	q := m.get(chatID)
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.Tracks)
}
