package entity

type Follow struct {
	UserId int      `db:"follow_id"`
	FollweeId int  `db:"followee_id"`
}