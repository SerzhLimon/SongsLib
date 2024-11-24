package models

type InfoSong struct {
	SongName    string
	Group       string
	ReleaseDate string
	Text        string
	Link        string
}

type SetSongRequest struct {
	SongName string `json:"song"`
	Group    string `json:"group"`
}

type SongPagination struct {
	CoupletNumber []int
	Text          []string
}

type SetSongInPostgres struct {
	InfoSong       InfoSong
	SongPagination SongPagination
}

type GetSongRequest struct {
	SongName string `json:"songname"`
	Offset   int    `json:"offset"`
}

type GetSongResponse struct {
	CoupletNumber int
	Couplet       string
}
