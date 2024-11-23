package repository

const (
	querySetUser = `
		insert into users (first_name, last_name, middle_name, gender, age, nationality)
		values ($1, $2, $3, $4, $5, $6)
	`
	queryGetUsers = `
		SELECT first_name, last_name, middle_name, gender, age, country 
		FROM users
		ORDER BY id
		LIMIT $1 OFFSET ($2) * $1
	`

	queryMakeFriendV1 = `
		UPDATE users
		SET friends = array_append(friends, $2)
		WHERE id = $1;
	`

	queryMakeFriendV2 = `
		UPDATE users
		SET friends = array_append(friends, $1)
		WHERE id = $2;
	`

	queryGetFriendsByUserID = `
		SELECT first_name, middle_name, last_name
		FROM users
		WHERE friends @> ARRAY[$1]
		ORDER BY id
		LIMIT $2 OFFSET ($3) * $2
	`

)