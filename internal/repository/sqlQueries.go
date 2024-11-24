package repository

const (
	querySetSongInfo = `
		INSERT INTO songs_info (name, group_name, link, release_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	// querySetSongText = `
	// 	INSERT INTO songs_text (track_id, couplet_number, couplet_text)
	// 	VALUES %s
	// `

)