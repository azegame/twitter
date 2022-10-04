package service

import(
	//"fmt"

	//"github.com/gin-gonic/gin"
	//"golang.org/x/crypto/bcrypt"

	"twitter/model/repository"
)


func Comment(tweetId int, userId int, message string) error {
	err := repository.InsertCommentInfo(tweetId, userId, message)
	return err
}