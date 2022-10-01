package entity

type TweetInfo struct {
	TweetId int     `db:"tweet_id"`
	UserId int      `db:"user_id"`
	Message string  `db:"message"`
	CreateAt string `db:"create_at"`
	UserName string `db:"user_name"`
}