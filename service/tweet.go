package service

import(
	"twitter/model/repository"
)


func Tweet(userId int, message string) error {
	err := repository.InsertTweet(userId, message)
	return err
}