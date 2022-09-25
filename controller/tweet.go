package controller

import (
    "fmt"
    //"strconv"
    
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"

    "twitter/service"
)


func indexPage(c *gin.Context) {
	claims := c.Keys["claims"]
    //claimsはjwt.MapClaims型と型アサーション、取り出したuserNameはstringなのでstring型と認識
    userName := claims.(jwt.MapClaims)["user_name"]//.(string)
    

    c.HTML(200, "index.html", gin.H{
    	"name": userName,
    })
}


func postTweetAnd(c *gin.Context) {
	claims := c.Keys["claims"]
    userId := int(claims.(jwt.MapClaims)["user_id"].(float64))
	userName := claims.(jwt.MapClaims)["user_name"]

	message := c.PostForm("message")
	fmt.Println(message)

	err := service.Tweet(userId, message)

	c.HTML(200, "index.html", gin.H{
		"name": userName,
		"error": err,
	})
}

