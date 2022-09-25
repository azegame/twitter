package repository

import (
	"fmt"

	//sqlite3、entityのimportはrepositoryだけ
	 _ "github.com/mattn/go-sqlite3"

	 "twitter/model/entity"
)


func InsertTweetInfo(userId int, message string) error {
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


func GetTweet(userId int) []entity.Tweet {
	var tweets []entity.Tweet
	rows, err := db.Query(
		`SELECT
			tweet_id,
			user_id,
			message,
			create_at
		 FROM
		 	tweets
		 WHERE
		 	user_id = ?
		 ORDER BY
		 	create_at`,
		userId,
	)
	fmt.Printf("GetTweet())のエラー %v\n ", err)

	for rows.Next() {
		r := entity.Tweet{}
		err = rows.Scan(
			&r.TweetId,
			&r.UserId,
			&r.Message,
			&r.CreateAt,
		)
		if err != nil {
			break
		}
		tweets = append(tweets, r)
	}
	return tweets
}


