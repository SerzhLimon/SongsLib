package usecase

import (
	"github.com/SerzhLimon/SongsLib/internal/repository"
)

type Usecase struct {
	pgPepo repository.Repository
}

type UseCase interface {
	GetSongs() error
}

func NewUsecase(pgPepo repository.Repository) UseCase {
	return &Usecase{pgPepo: pgPepo}
}

func (u *Usecase) GetSongs() error {
	return nil
}