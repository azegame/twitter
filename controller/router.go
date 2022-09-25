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
        auth.POST("/tweets", postTweetAnd)
        auth.GET("/users/:UserId", userPage)
    }

	return r
}

