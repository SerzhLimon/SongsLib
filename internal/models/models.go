package models

type SetSongRequest struct {
	SongName string `json:"song"`
	Group    string `json:"group"`
}

type InfoSong struct {
	SongName    string
	Group       string
	ReleaseDate string
	Text        string
	Link        string
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

type GetSongInfo struct {
	ID          int
	SongName    string
	Group       string
	ReleaseDate string
	Link        string
}

type GetLibRequest struct {
	SongName    string `json:"songname"`
	Group       string `json:"group"`
	ReleaseDate string `json:"releasedate"`
	Link        string `json:"link"`
	Offset      int    `json:"offset"`
}

type GetLibResponse struct {
	Songs []GetSongInfo
}

type DeleteSongRequest struct {
	TrackID int `json:"id"`
}

type UpdateSongInfoRequest struct {
	TrackID        int     `json:"id"`
	NewSongName    *string `json:"songname,omitempty"`
	NewGroup       *string `json:"group,omitempty"`
	NewReleaseDate *string `json:"releasedate,omitempty"`
	NewLink        *string `json:"link,omitempty"`
}

type UpdateSongTextRequest struct {
	TrackID    int     `json:"id"`
	CoupletNum int     `json:"coupletnum"`
	NewText    *string `json:"text,omitempty"`
}
