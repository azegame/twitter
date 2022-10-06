package service

import(
	"twitter/model/repository"
)


func Tweet(userId int, message string) error {
	err := repository.InsertTweetInfo(userId, message)
	return err
}