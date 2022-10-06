package service

import(
	"twitter/model/repository"
)


func Comment(tweetId int, userId int, message string) error {
	err := repository.InsertCommentInfo(tweetId, userId, message)
	return err
}