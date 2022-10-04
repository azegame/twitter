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
			T.create_at Desc`,
		userId,
	)
	fmt.Printf("GetTweetInfo()のエラー %v\n", err)

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


func GetTimeLine(userId int) []entity.TweetInfo {
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
			T.user_id 
		 IN (
			SELECT
				followee_id
			FROM
				follows
			WHERE
				follow_id = ?
			)
		 ORDER BY
			 T.create_at Desc`,
		userId,
	)
	fmt.Printf("GetTimeLine()のエラー %v\n", err)

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


func GetTweetByTweetId(tweetId int) (entity.TweetInfo, error) {
	var tweet entity.TweetInfo

	row := db.QueryRow(
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
			T.tweet_id = ?`,
		tweetId,
	)
	err := row.Scan(
		&tweet.TweetId,
		&tweet.UserId,
		&tweet.Message,
		&tweet.CreateAt,
		&tweet.UserName,
	)
	return tweet, err
}
















