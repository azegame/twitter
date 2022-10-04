package repository

import (
	"fmt"
	//sqlite3、entityのimportはrepositoryだけ
	 _ "github.com/mattn/go-sqlite3"

	 "twitter/model/entity"
)


func FindUserByUserName(userName string) (entity.User, error) {
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


func FindUserByUserId(userId int) (entity.User, error) {
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
			 user_id = ?`,
		userId,
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


func GetOtherUser(userId int) []entity.User {
	var users []entity.User

	rows, err := db.Query(
		`SELECT
			user_id,
			user_name
		 FROM
			 users
		 WHERE
			 user_id != ?
		 ORDER BY
			 user_id`,
		userId,
	)
	fmt.Printf("GetOtherUser())のエラー %v\n", err)

	for rows.Next() {
		r := entity.User{}
		err = rows.Scan(
			&r.UserId,
			&r.UserName,
		)
		if err != nil {
			break
		}
		users = append(users, r)
	}
	return users
}

