package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/SerzhLimon/SongsLib/internal/models"
)

type Repository interface {
	SetSong(data models.SetSongInPostgres) error
	GetSong(data models.GetSongRequest) (models.GetSongResponse, error)
}

type pgRepo struct {
	db *sql.DB
}

func NewPGRepository(db *sql.DB) Repository {
	return &pgRepo{db: db}
}

func (r *pgRepo) SetSong(data models.SetSongInPostgres) error {
	tx, err := r.db.Begin()
	if err != nil {
		err := errors.Errorf("pgRepo.SetSong %v", err)
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var songID int
	if err = tx.QueryRow(querySetSongInfo,
		data.InfoSong.SongName,
		data.InfoSong.Group,
		data.InfoSong.Link,
		data.InfoSong.ReleaseDate,
	).Scan(&songID); err != nil {
		err := errors.Errorf("pgRepo.SetSong %v", err)
		return err
	}

	valueStrings := make([]string, 0, len(data.SongPagination.CoupletNumber))
	valueArgs := make([]interface{}, 0)

	for i, number := range data.SongPagination.CoupletNumber {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
		valueArgs = append(valueArgs, songID, number, data.SongPagination.Text[i])
	}

	querySetSongText := `
		INSERT INTO songs_text (track_id, couplet_number, couplet_text)
		VALUES %s
	`
	querySetSongText = fmt.Sprintf(querySetSongText, strings.Join(valueStrings, ","))

	if _, err = tx.Exec(querySetSongText, valueArgs...); err != nil {
		err = errors.Errorf("pgRepo.SetSong %v", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		err = errors.Errorf("pgRepo.SetSong %v", err)
		return err
	}

	return nil
}

func (r *pgRepo) GetSong(data models.GetSongRequest) (models.GetSongResponse, error) {
	var res models.GetSongResponse 

	err := r.db.QueryRow(queryGetSong, data.SongName, data.Offset).Scan(
		&res.CoupletNumber,
		&res.Couplet,
    )
    if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		err = errors.Errorf("pgRepo.GetSong %v", err)
		return res, err
    }

	return res, nil
}