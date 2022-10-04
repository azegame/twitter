package repository

import (
	"fmt"

	//sqlite3、entityのimportはrepositoryだけ
	 _ "github.com/mattn/go-sqlite3"

	"twitter/model/entity"
)


func InsertCommentInfo(tweetId int, userId int, message string) error {
	_, err := db.Exec(
		`INSERT INTO comments (
			tweet_id,
			user_id,
			message
		 ) VALUES (?,?,?)`,
		tweetId,
		userId,
		message,
	)
	return err
}


func GetComments(tweetId int) []entity.CommentInfo {
	var comments []entity.CommentInfo
	rows, err := db.Query(
		`SELECT
			C.comment_id,
			C.tweet_id,
			C.user_id,
			C.message,
			C.create_at,
			U.user_name
		 FROM
			comments AS C LEFT OUTER JOIN users AS U
		 ON
			C.user_id = U.user_id
		 WHERE
			C.tweet_id = ?
		 ORDER BY
			C.create_at`,
		tweetId,
	)
	fmt.Printf("GetComments()のエラー %v\n", err)

	for rows.Next() {
		r := entity.CommentInfo{}
		err = rows.Scan(
			&r.CommentId,
			&r.TweetId,
			&r.UserId,
			&r.Message,
			&r.CreateAt,
			&r.UserName,
		)
		if err != nil {
			break
		}
		comments = append(comments, r)
	}
	return comments
}



