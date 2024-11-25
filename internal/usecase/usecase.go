package usecase

import (
	"strings"

	"github.com/SerzhLimon/SongsLib/internal/models"
	"github.com/SerzhLimon/SongsLib/internal/repository"
)

type Usecase struct {
	pgPepo repository.Repository
}

type UseCase interface {
	SetSong(data models.InfoSong) error
	GetSong(data models.GetSongRequest) (models.GetSongResponse, error)
	GetLib(data models.GetLibRequest) (models.GetLibResponse, error)
	DeleteSong(data models.DeleteSongRequest) error
	UpdateSongInfo(data models.UpdateSongInfoRequest) error
	UpdateSongText(data models.UpdateSongTextRequest) error
}

func NewUsecase(pgPepo repository.Repository) UseCase {
	return &Usecase{pgPepo: pgPepo}
}

func (u *Usecase) SetSong(data models.InfoSong) error {

	text := u.parseText(data.Text)
	song := models.SetSongInPostgres{
		InfoSong:       data,
		SongPagination: text,
	}

	return u.pgPepo.SetSong(song)
}

func (u *Usecase) parseText(text string) models.SongPagination {

	couplets := strings.Split(text, "\n\n")
	result := models.SongPagination{
		CoupletNumber: make([]int, 0, len(couplets)),
		Text:           make([]string, 0, len(couplets)),
	}

	for i, couplet := range couplets {
		result.CoupletNumber = append(result.CoupletNumber, i+1)
		result.Text = append(result.Text, couplet)
	}

	return result
}

func (u *Usecase) GetSong(data models.GetSongRequest) (models.GetSongResponse, error) {
	return u.pgPepo.GetSong(data)
}

func (u *Usecase) GetLib(data models.GetLibRequest) (models.GetLibResponse, error) {
	return u.pgPepo.GetLib(data)
}

func (u *Usecase) DeleteSong(data models.DeleteSongRequest) error {
	return u.pgPepo.DeleteSong(data)
}

func (u *Usecase) UpdateSongInfo(data models.UpdateSongInfoRequest) error {
	return u.pgPepo.UpdateSongInfo(data)
}

func (u *Usecase) UpdateSongText(data models.UpdateSongTextRequest) error {
	return u.pgPepo.UpdateSongText(data)
}