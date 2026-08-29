package platforms

// YouTube search & stream URL resolution.
// Port from Python PANDAMUSIC/platforms/Youtube.py
// Planned: yt-dlp exec or go library for search + best audio URL.

type Result struct {
	Title    string
	URL      string
	Duration string
	Thumb    string
	ID       string
}

func Search(query string, limit int) ([]Result, error) {
	// TODO: implement yt-dlp / API search
	return []Result{{
		Title: query,
		URL:   "ytsearch:" + query,
	}}, nil
}

func Resolve(urlOrQuery string) (*Result, error) {
	// TODO: resolve direct watch URL or first search hit
	return &Result{
		Title: urlOrQuery,
		URL:   urlOrQuery,
	}, nil
}
