package service

import(
	//"fmt"

    //"github.com/gin-gonic/gin"
    //"golang.org/x/crypto/bcrypt"

    "twitter/model/repository"
)


func Tweet(userId int, message string) error {
	err := repository.InsertTweetInfo(userId, message)
	return err
}