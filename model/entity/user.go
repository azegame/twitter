package entity

type User struct {
	UserId int      `db:"user_id"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
	CreateAt string `db:"create_at"`
}



