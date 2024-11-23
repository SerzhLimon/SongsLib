package models

type InfoSongResponse struct {
	ReleaseDate string
	Text        string
	Link        string
}

type SetSongRequest struct {
	SongName string `json:"song"`
	Group    string `json:"group"`
}
