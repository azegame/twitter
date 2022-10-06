package controller

import (
	"github.com/gin-gonic/gin"

	"twitter/service"
)


func GetRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("view/template/*.html")

	r.GET("/signup", signupPage)
	r.POST("/signup", signup)
	r.GET("/login", loginPage)
	r.POST("/login", login)
	r.GET("/logout", logout)
	
	auth := r.Group("/")
	auth.Use(service.AuthorizationMiddleware()) 
	{
		auth.GET("/", indexPage)
		auth.POST("/tweets", postTweet)
		auth.GET("/users/:UserId", otherUserPage)
		auth.POST("/follow", follow)
		auth.POST("/unfollow", unFollow)
		auth.GET("/tweets/:TweetId/comments", commentPage)
		auth.POST("/tweets/:TweetId/comments", postComment)
	}
	return r
}