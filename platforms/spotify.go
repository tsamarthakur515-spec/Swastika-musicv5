package platforms

// Spotify URL metadata stub — full port later with API keys.

type SpotifyTrack struct {
	Title  string
	Artist string
	URL    string
}

func ParseSpotify(url string) (*SpotifyTrack, error) {
	return &SpotifyTrack{
		Title:  "Spotify track",
		Artist: "Unknown",
		URL:    url,
	}, nil
}
