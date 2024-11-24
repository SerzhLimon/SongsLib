package repository

const (
	querySetSongInfo = `
		INSERT INTO songs_info (name, group_name, link, release_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	queryGetSong = `
		WITH ids AS (
			SELECT id
			FROM songs_info
			WHERE name = $1
		)
		SELECT couplet_number, couplet_text
		FROM songs_text
		WHERE track_id IN (SELECT id FROM ids)
		ORDER BY couplet_number
		LIMIT 1 OFFSET ($2) - 1
	`

)