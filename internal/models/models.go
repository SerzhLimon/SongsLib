package models

type InfoSong struct {
	ReleaseDate string
	Text        string
	Link        string
}

type SetSongRequest struct {
	SongName string `json:"song"`
	Group    string `json:"group"`
}

type SongPagination struct {
	Couplet_number []int
	Text           []string
}

