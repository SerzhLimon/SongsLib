package repository

import (
	"database/sql"
)

const limit = 3

type Repository interface {
	GetSongs() error
}

type pgRepo struct {
	db *sql.DB
}

func NewPGRepository(db *sql.DB) Repository {
	return &pgRepo{db: db}
}

func (r *pgRepo) GetSongs() error {
	return nil
}
