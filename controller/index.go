package controller

import (
	//"fmt"
	//"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"twitter/model/repository"
	"twitter/service"
)


func indexPage(c *gin.Context) {
	claims := c.Keys["claims"]
	userId := int(claims.(jwt.MapClaims)["user_id"].(float64))
	//claimsはjwt.MapClaims型と型アサーション、
	//取り出したuserNameはstringなのでstring型と認識
	userName := claims.(jwt.MapClaims)["user_name"]//.(string)

	users := repository.GetOtherUser(userId)
	timeLine := repository.GetTimeLine(userId)


	c.HTML(200, "index.html", gin.H{
		"myname": userName,
		"users": users,
		"timeline": timeLine,
	})
}


func postTweet(c *gin.Context) {
	claims := c.Keys["claims"]
	userId := int(claims.(jwt.MapClaims)["user_id"].(float64))
	userName := claims.(jwt.MapClaims)["user_name"]

	users := repository.GetOtherUser(userId)
	message := c.PostForm("message")

	err := service.Tweet(userId, message)

	timeLine := repository.GetTimeLine(userId)

	c.HTML(200, "index.html", gin.H{
		"error": err,
		"myname": userName,
		"users": users,
		"timeline": timeLine,
	})
}



