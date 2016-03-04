package netflixroulette

type Media struct {
	Unit        int    `json:"unit"`
	ShowId      int    `json:"show_id"`
	ShowTitle   string `json:"show_title"`
	ReleaseYear string `json:"release_year"`
	Rating      string `json:"rating"`
	Category    string `json:"category"`
	ShowCast    string `json:"show_cast"`
	Director    string `json:"director"`
	Summary     string `json:"summary"`
	Poster      string `json:"poster"`
	MediaType   int    `json:"mediatype"`
	Runtime     string `json:"runtime"`
}

func (media Media) IsMovie() bool {
	return media.MediaType == 0
}

func (media Media) IsTvShow() bool {
	return media.MediaType == 1
}

func (media Media) GetReadableMediaType() string {
	if media.IsMovie() {
		return "Movie"
	} else if media.IsTvShow() {
		return "TV Show"
	}

	return ""
}
