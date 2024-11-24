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
		Couplet_number: make([]int, 0, len(couplets)),
		Text:           make([]string, 0, len(couplets)),
	}

	for i, couplet := range couplets {
		result.Couplet_number = append(result.Couplet_number, i+1)
		result.Text = append(result.Text, couplet)
	}

	return result
}
