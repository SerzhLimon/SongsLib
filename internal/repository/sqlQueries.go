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
		SELECT id, name, group_name, link, release_date
		FROM songs_info
		WHERE 
		($3::text IS NULL OR name ILIKE '%' || $3 || '%')
		 AND
		($4::text IS NULL OR group_name ILIKE '%' || $4 || '%') 
		AND
		($5::text IS NULL OR link ILIKE '%' || $5 || '%') 
		AND
		($6::text IS NULL OR release_date ILIKE '%' || $6 || '%')
		ORDER BY id
		LIMIT $1::integer OFFSET ($2::integer) * $1::integer;
	`

	queryDeleteSongText = `
		DELETE FROM songs_text
		WHERE track_id = $1;
	`
	
	queryDeleteSongInfo = `
		DELETE FROM songs_info
		WHERE id = $1;
	`
	
	queryUpdateSongInfo = `
		UPDATE songs_info
		SET 
			name = COALESCE($2, name),
			group_name = COALESCE($3, group_name),
			release_date = COALESCE($4, release_date),
			link = COALESCE($5, link)
		WHERE id = $1;
	`

	queryUpdateSongText = `
		UPDATE songs_text
		SET 
			couplet_text = COALESCE($3, couplet_text)
		WHERE couplet_number = $1 AND track_id = $2;
	`


)