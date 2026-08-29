package platforms

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// Result is a resolved YouTube (or yt-dlp supported) track.
type Result struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	URL         string  `json:"webpage_url"`
	Duration    float64 `json:"duration"`
	Thumbnail   string  `json:"thumbnail"`
	Uploader    string  `json:"uploader"`
	StreamURL   string  `json:"url"`
	Extractor   string  `json:"extractor"`
}

func (r *Result) DurationStr() string {
	sec := int(r.Duration)
	if sec <= 0 {
		return "--:--"
	}
	m := sec / 60
	s := sec % 60
	if m >= 60 {
		return fmt.Sprintf("%d:%02d:%02d", m/60, m%60, s)
	}
	return fmt.Sprintf("%d:%02d", m, s)
}

// Resolve accepts a YouTube URL or search query and returns the first match.
func Resolve(query string) (*Result, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, fmt.Errorf("empty query")
	}

	target := query
	if !strings.HasPrefix(query, "http://") && !strings.HasPrefix(query, "https://") {
		target = "ytsearch1:" + query
	}

	ctxTimeout := 45 * time.Second
	cmd := exec.Command("yt-dlp", "-j", "--no-playlist", "--no-warnings", target)
	// hard timeout via process kill is left to host; keep command simple
	_ = ctxTimeout

	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("yt-dlp: %s", strings.TrimSpace(string(ee.Stderr)))
		}
		return nil, fmt.Errorf("yt-dlp failed: %w (is yt-dlp installed?)", err)
	}

	// ytsearch may print one JSON object per line; take first non-empty line
	line := strings.TrimSpace(string(out))
	if idx := strings.IndexByte(line, '\n'); idx >= 0 {
		line = line[:idx]
	}

	var r Result
	if err := json.Unmarshal([]byte(line), &r); err != nil {
		return nil, fmt.Errorf("parse yt-dlp json: %w", err)
	}
	if r.Title == "" {
		r.Title = query
	}
	if r.URL == "" && r.ID != "" {
		r.URL = "https://www.youtube.com/watch?v=" + r.ID
	}
	return &r, nil
}

// DownloadAudio downloads best audio as mp3 into destPath (without extension; yt-dlp adds it).
// Returns the final file path.
func DownloadAudio(pageURL, destPrefix string) (string, error) {
	outTpl := destPrefix + ".%(ext)s"
	cmd := exec.Command(
		"yt-dlp",
		"-x", "--audio-format", "mp3", "--audio-quality", "0",
		"--no-playlist", "--no-warnings",
		"-o", outTpl,
		pageURL,
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("download: %v — %s", err, strings.TrimSpace(string(out)))
	}
	return destPrefix + ".mp3", nil
}
