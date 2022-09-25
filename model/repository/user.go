package repository

import (
	//sqlite3、entityのimportはrepositoryだけ
	 _ "github.com/mattn/go-sqlite3"

	 "twitter/model/entity"
)


func FindUser(userName string) (entity.User, error) {
	var user entity.User

	row := db.QueryRow(
		`SELECT
			user_id,
			user_name,
			password,
			create_at
		 FROM
		 	users
		 WHERE
		 	user_name = ?`,
		userName,
	)

	err := row.Scan(
		&user.UserId,
		&user.UserName,
		&user.Password,
		&user.CreateAt,
	)
	return user, err
}


func CreateUser(userName string, password string) error {
	_, err := db.Exec(
		`INSERT INTO users (
			user_name,
			password
		 ) VALUES (?,?)`, 
		userName,
		password,
	)
	return err
}

