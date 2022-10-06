package controller

import (
	"fmt"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"twitter/model/repository"
	"twitter/service"
)


func commentPage(c *gin.Context) {
	tweetId, _ := strconv.Atoi(c.Param("TweetId"))

	claims := c.Keys["claims"]
	userName := claims.(jwt.MapClaims)["user_name"]

	tweet, err := repository.GetTweetByTweetId(tweetId)
	if err != nil {
		fmt.Println(err)
	}
	comments := repository.GetComments(tweetId)

	c.HTML(200, "comment.html", gin.H{
		"TweetId":tweetId,
		"tweet": tweet,
		"myname": userName,
		"comments": comments,
	})
}


func postComment(c *gin.Context) {
	tweetId, _ := strconv.Atoi(c.Param("TweetId"))

	claims := c.Keys["claims"]
	userId := int(claims.(jwt.MapClaims)["user_id"].(float64))
	userName := claims.(jwt.MapClaims)["user_name"]

	message := c.PostForm("message")

	err := service.Comment(tweetId, userId, message)
	if err != nil {
		fmt.Println(err)
	}

	tweet, err := repository.GetTweetByTweetId(tweetId)
	if err != nil {
		fmt.Println(err)
	}
	
	comments := repository.GetComments(tweetId)

	c.HTML(200, "comment.html", gin.H{
		"TweetId":tweetId,
		"tweet": tweet,
		"myname": userName,
		"comments": comments,
	})
}