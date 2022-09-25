package repository

import (
	"fmt"
	
	//sqlite3、entityのimportはrepositoryだけ
	 _ "github.com/mattn/go-sqlite3"

	 //"twitter/model/entity"
)


func InsertTweetInfo(userId int, message string) error {
	fmt.Println(message)
	_, err := db.Exec(
		`INSERT INTO tweets (
			user_id,
			message
		 ) VALUES (?,?)`,
		userId,
		message,
	)
	return err
}