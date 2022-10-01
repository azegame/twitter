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


func GetTweetInfo(userId int) []entity.TweetInfo {
	var tweets []entity.TweetInfo
	rows, err := db.Query(
		`SELECT
			T.tweet_id,
			T.user_id,
			T.message,
			T.create_at,
			U.user_name
		 FROM
		 	tweets AS T LEFT OUTER JOIN users AS U
		 ON
		 	T.user_id = U.user_id
		 WHERE
		 	T.user_id = ?
		 ORDER BY
		 	T.create_at`,
		userId,
	)
	fmt.Printf("GetTweetInfo()のエラー %v\n ", err)

	for rows.Next() {
		r := entity.TweetInfo{}
		err = rows.Scan(
			&r.TweetId,
			&r.UserId,
			&r.Message,
			&r.CreateAt,
			&r.UserName,
		)
		if err != nil {
			break
		}
		tweets = append(tweets, r)
	}
	return tweets
}

