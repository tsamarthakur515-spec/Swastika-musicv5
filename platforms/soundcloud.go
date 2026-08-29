package platforms

// SoundCloud stub

type SCTrack struct {
	Title string
	URL   string
}

func ParseSoundCloud(url string) (*SCTrack, error) {
	return &SCTrack{Title: "SoundCloud track", URL: url}, nil
}
