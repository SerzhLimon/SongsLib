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
	
	queryGetLib = `
		SELECT name, group_name, link, release_date
		FROM songs_info
		WHERE 
		($3::text IS NULL OR name ILIKE '%' || $3 || '%') AND
		($4::text IS NULL OR group_name ILIKE '%' || $4 || '%') AND
		($5::text IS NULL OR link ILIKE '%' || $5 || '%') AND
		($6::text IS NULL OR release_date = $6)
		LIMIT $1 OFFSET ($2) * $1;
	`
)