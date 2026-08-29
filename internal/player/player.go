package player

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/amarnathcjd/gogram/telegram"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/queue"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/platforms"
)

// Player handles download + delivery.
// Full Telegram Voice Chat streaming (gortc/ntgcalls) will plug in here later.
type Player struct {
	mu sync.Mutex
}

func New() *Player {
	return &Player{}
}

// PlayTrack downloads audio and sends it to the chat as media.
func (p *Player) PlayTrack(client *telegram.Client, chatID int64, track queue.Track) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if track.URL == "" {
		return fmt.Errorf("empty track url")
	}

	dir := os.TempDir()
	prefix := filepath.Join(dir, fmt.Sprintf("swastika_%d", chatID))
	file, err := platforms.DownloadAudio(track.URL, prefix)
	if err != nil {
		return err
	}
	defer os.Remove(file)

	caption := fmt.Sprintf("🎵 <b>%s</b>\n👤 %s\n⏱ %s", track.Title, track.RequestedBy, track.Duration)
	_, err = client.SendMedia(chatID, file, &telegram.MediaOptions{
		Caption:   caption,
		ParseMode: "HTML",
	})
	return err
}
