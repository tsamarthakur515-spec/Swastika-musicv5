package stream

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/amarnathcjd/gogram/telegram"
	"github.com/amarnathcjd/gortc"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/queue"
)

// Manager handles per-chat VC sessions via gortc.
// Real VC join needs a user client (STRING_SESSION). Bot-only falls back to file send.
type Manager struct {
	mu       sync.Mutex
	client   *telegram.Client // preferably userbot
	calls    map[int64]*gortc.Call
	cancels  map[int64]context.CancelFunc
	enabled  bool
}

func New(client *telegram.Client, enabled bool) *Manager {
	return &Manager{
		client:  client,
		calls:   make(map[int64]*gortc.Call),
		cancels: make(map[int64]context.CancelFunc),
		enabled: enabled && client != nil,
	}
}

func (m *Manager) Enabled() bool {
	return m != nil && m.enabled
}

// PlayFile joins VC (if possible) and streams local file.
func (m *Manager) PlayFile(chatID int64, path string, track queue.Track) error {
	if !m.Enabled() {
		return fmt.Errorf("VC stream disabled — set STRING_SESSION for userbot assistant")
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	if cancel, ok := m.cancels[chatID]; ok {
		cancel()
		delete(m.cancels, chatID)
	}
	if old, ok := m.calls[chatID]; ok {
		_ = old.Leave()
		delete(m.calls, chatID)
	}

	call := gortc.NewCall(m.client)
	ctx, cancel := context.WithCancel(context.Background())
	m.cancels[chatID] = cancel
	m.calls[chatID] = call

	call.OnConnected(func() {
		go func() {
			err := call.Stream(ctx, gortc.FromFile(path))
			if err != nil && ctx.Err() == nil {
				log.Printf("[stream] chat %d stream end: %v", chatID, err)
			}
		}()
	})

	if err := call.Join(chatID); err != nil {
		cancel()
		delete(m.cancels, chatID)
		delete(m.calls, chatID)
		return fmt.Errorf("join VC: %w", err)
	}
	log.Printf("[stream] joined VC chat=%d title=%s", chatID, track.Title)
	return nil
}

func (m *Manager) Pause(chatID int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	call, ok := m.calls[chatID]
	if !ok {
		return fmt.Errorf("not streaming")
	}
	return call.SetMuted(true)
}

func (m *Manager) Resume(chatID int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	call, ok := m.calls[chatID]
	if !ok {
		return fmt.Errorf("not streaming")
	}
	return call.SetMuted(false)
}

func (m *Manager) Stop(chatID int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if cancel, ok := m.cancels[chatID]; ok {
		cancel()
		delete(m.cancels, chatID)
	}
	if call, ok := m.calls[chatID]; ok {
		err := call.Leave()
		delete(m.calls, chatID)
		return err
	}
	return nil
}
