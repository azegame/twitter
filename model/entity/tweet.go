package entity

type Tweet struct {
	TweetId int     `db:"tweet_id"`
	UserId int      `db:"user_id"`
	Message string  `db:"message"`
	CreateAt string `db:"create_at"`
}