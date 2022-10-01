package repository

import (
	//"fmt"

	//sqlite3、entityのimportはrepositoryだけ
	 _ "github.com/mattn/go-sqlite3"

	 "twitter/model/entity"
)


func InsertFollowInfo(userIdByJWT int, otherUserId int) error {
	_, err := db.Exec(
		`INSERT INTO follows (
			follow_id,
			followee_id
		 ) VALUES (?,?)`,
		userIdByJWT,
		otherUserId,
	)
	return err
}


func DeleteFollowInfo(userIdByJWT int, otherUserId int) error {
		_, err := db.Exec(
		`DELETE FROM follows
		 WHERE 
		 	follow_id = ?
		 AND
		 	followee_id = ?`,
		userIdByJWT,
		otherUserId,
	)
	return err
}


func SearchFollowInfo(userIdByJWT int, otherUserId int) (entity.Follow, error) {
	var followInfo entity.Follow

	row := db.QueryRow(
		`SELECT
			follow_id,
			followee_id
		 FROM
		 	follows
		 WHERE 
		 	follow_id = ?
		 AND
		 	followee_id = ?`,
		userIdByJWT,
		otherUserId,
	)

	err := row.Scan(
		&followInfo.UserId,
		&followInfo.FollweeId,
	)
	return followInfo, err
}

