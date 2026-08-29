package platforms

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Defaults match Python SWASTIKA-MUSIC-V5 Youtube.py
var (
	APIURL = "https://aruyt.up.railway.app"
	APIKey = "YUKI-XJd3KfUSWeuOWsiZyuIrlmQf"
	DLDir  = "downloads"
)

// Configure from env/config (call once at startup).
func Configure(apiURL, apiKey string) {
	if apiURL != "" {
		APIURL = strings.TrimRight(apiURL, "/")
	}
	if apiKey != "" {
		APIKey = apiKey
	}
}

type Result struct {
	ID        string
	Title     string
	URL       string
	Duration  float64
	Thumbnail string
	Uploader  string
}

func (r *Result) DurationStr() string {
	sec := int(r.Duration)
	if sec <= 0 {
		return "0:00"
	}
	m := sec / 60
	s := sec % 60
	if m >= 60 {
		return fmt.Sprintf("%d:%02d:%02d", m/60, m%60, s)
	}
	return fmt.Sprintf("%d:%02d", m, s)
}

func ToVidID(value string) string {
	v := strings.TrimSpace(value)
	if strings.Contains(v, "v=") {
		v = strings.Split(strings.Split(v, "v=")[1], "&")[0]
	}
	if strings.Contains(v, "youtu.be/") {
		v = strings.Split(strings.Split(v, "youtu.be/")[1], "?")[0]
	}
	return strings.TrimSpace(v)
}

// Resolve — search like Python (URL/id fast-path, else yt-dlp ytsearch).
func Resolve(query string) (*Result, error) {
	q := strings.TrimSpace(query)
	if q == "" {
		return nil, fmt.Errorf("empty query")
	}

	// Direct YouTube URL / 11-char id
	if strings.Contains(q, "youtube.com") || strings.Contains(q, "youtu.be/") || (len(q) == 11 && !strings.Contains(q, " ")) {
		vid := ToVidID(q)
		if len(vid) >= 10 {
			return &Result{
				ID:        vid,
				Title:     "YouTube Video",
				URL:       "https://www.youtube.com/watch?v=" + vid,
				Thumbnail: "https://i.ytimg.com/vi/" + vid + "/hqdefault.jpg",
				Uploader:  "YouTube Music",
			}, nil
		}
	}

	return searchYTDLP(q)
}

func searchYTDLP(query string) (*Result, error) {
	target := "ytsearch1:" + query
	cmd := exec.Command("yt-dlp", "-j", "--flat-playlist", "--no-playlist", "--no-warnings", target)
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("yt-dlp search: %s", strings.TrimSpace(string(ee.Stderr)))
		}
		return nil, fmt.Errorf("yt-dlp search failed: %w", err)
	}
	line := strings.TrimSpace(string(out))
	if i := strings.IndexByte(line, '\n'); i >= 0 {
		line = line[:i]
	}
	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(line), &raw); err != nil {
		return nil, fmt.Errorf("parse search json: %w", err)
	}
	id, _ := raw["id"].(string)
	title, _ := raw["title"].(string)
	webpage, _ := raw["webpage_url"].(string)
	if webpage == "" {
		webpage, _ = raw["url"].(string)
	}
	if webpage == "" && id != "" {
		webpage = "https://www.youtube.com/watch?v=" + id
	}
	var dur float64
	switch d := raw["duration"].(type) {
	case float64:
		dur = d
	case json.Number:
		dur, _ = d.Float64()
	}
	thumb, _ := raw["thumbnail"].(string)
	if thumb == "" && id != "" {
		thumb = "https://i.ytimg.com/vi/" + id + "/hqdefault.jpg"
	}
	uploader, _ := raw["uploader"].(string)
	if title == "" {
		title = query
	}
	return &Result{
		ID:        id,
		Title:     title,
		URL:       webpage,
		Duration:  dur,
		Thumbnail: thumb,
		Uploader:  uploader,
	}, nil
}

// DownloadSong — Python download_song: API first, then yt-dlp audio.
func DownloadSong(vidid string) (string, error) {
	vidid = ToVidID(vidid)
	if vidid == "" {
		return "", fmt.Errorf("invalid vidid")
	}
	_ = os.MkdirAll(DLDir, 0o755)

	if path, err := downloadAPI(vidid, "audio", "mp3", 80*time.Second); err == nil && path != "" {
		return path, nil
	}
	return downloadYTDLPAudio(vidid)
}

// DownloadVideo — Python download_video: API first, then yt-dlp video.
func DownloadVideo(vidid string) (string, error) {
	vidid = ToVidID(vidid)
	if vidid == "" {
		return "", fmt.Errorf("invalid vidid")
	}
	_ = os.MkdirAll(DLDir, 0o755)

	if path, err := downloadAPI(vidid, "video", "mp4", 150*time.Second); err == nil && path != "" {
		return path, nil
	}
	return downloadYTDLPVideo(vidid)
}

func downloadAPI(vidid, mediaType, ext string, timeout time.Duration) (string, error) {
	if APIKey == "" {
		return "", fmt.Errorf("API_KEY missing")
	}
	filePath := filepath.Join(DLDir, vidid+"."+ext)
	if st, err := os.Stat(filePath); err == nil && st.Size() > 1024 {
		return filePath, nil
	}

	fullURL := "https://www.youtube.com/watch?v=" + vidid
	client := &http.Client{Timeout: timeout}
	u, _ := url.Parse(strings.TrimRight(APIURL, "/") + "/download")
	q := u.Query()
	q.Set("url", fullURL)
	q.Set("type", mediaType)
	q.Set("api_key", APIKey)
	u.RawQuery = q.Encode()

	resp, err := client.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 400))
		return "", fmt.Errorf("API HTTP %d: %s", resp.StatusCode, string(body))
	}
	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := io.Copy(f, resp.Body); err != nil {
		_ = os.Remove(filePath)
		return "", err
	}
	st, err := os.Stat(filePath)
	if err != nil || st.Size() < 1024 {
		_ = os.Remove(filePath)
		return "", fmt.Errorf("API file too small")
	}
	return filePath, nil
}

func downloadYTDLPAudio(vidid string) (string, error) {
	outTpl := filepath.Join(DLDir, vidid+".%(ext)s")
	final := filepath.Join(DLDir, vidid+".mp3")
	cmd := exec.Command(
		"yt-dlp",
		"-f", "bestaudio[ext=m4a]/bestaudio/best",
		"-x", "--audio-format", "mp3", "--audio-quality", "192",
		"--no-playlist", "--no-warnings", "--retries", "3",
		"-o", outTpl,
		"https://www.youtube.com/watch?v="+vidid,
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("yt-dlp audio: %v — %s", err, strings.TrimSpace(string(out)))
	}
	if st, err := os.Stat(final); err == nil && st.Size() > 1024 {
		return final, nil
	}
	// find any audio produced
	entries, _ := os.ReadDir(DLDir)
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, vidid+".") {
			for _, ext := range []string{".mp3", ".m4a", ".webm", ".opus", ".ogg"} {
				if strings.HasSuffix(name, ext) {
					return filepath.Join(DLDir, name), nil
				}
			}
		}
	}
	return "", fmt.Errorf("yt-dlp audio: no output file")
}

func downloadYTDLPVideo(vidid string) (string, error) {
	outTpl := filepath.Join(DLDir, vidid+".%(ext)s")
	final := filepath.Join(DLDir, vidid+".mp4")
	cmd := exec.Command(
		"yt-dlp",
		"-f", "best[height<=720][ext=mp4]/best[ext=mp4]/best",
		"--merge-output-format", "mp4",
		"--no-playlist", "--no-warnings", "--retries", "3",
		"-o", outTpl,
		"https://www.youtube.com/watch?v="+vidid,
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("yt-dlp video: %v — %s", err, strings.TrimSpace(string(out)))
	}
	if st, err := os.Stat(final); err == nil && st.Size() > 1024 {
		return final, nil
	}
	entries, _ := os.ReadDir(DLDir)
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, vidid+".") {
			for _, ext := range []string{".mp4", ".mkv", ".webm"} {
				if strings.HasSuffix(name, ext) {
					return filepath.Join(DLDir, name), nil
				}
			}
		}
	}
	return "", fmt.Errorf("yt-dlp video: no output file")
}

// DownloadAudio keeps older player API: resolve page URL → song file.
func DownloadAudio(pageURL, _destPrefix string) (string, error) {
	vid := ToVidID(pageURL)
	if vid == "" {
		vid = pageURL
	}
	return DownloadSong(vid)
}
