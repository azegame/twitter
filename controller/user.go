package controller

import (
    //"fmt"
    //"strconv"

    "github.com/gin-gonic/gin"
    //"github.com/golang-jwt/jwt/v4"
    //_ "github.com/mattn/go-sqlite3"

    "twitter/service"
    //"twitter/model/repository"
    //"twitter/model/entity" 
)


func signupPage(c *gin.Context) {
    c.HTML(200, "signup.html", gin.H{})
}


func signup(c *gin.Context) {
	userName := c.PostForm("username")
    password := c.PostForm("password")

    if service.Signup(userName, password) {
    	c.Redirect(302, "/login")
    	c.Abort()
        return
    } else {
    	c.HTML(409, "signup.html", gin.H{
    		"error": "既に使われているユーザー名です",
    	})
	}
}


func loginPage(c *gin.Context) {
    c.HTML(200, "login.html", gin.H{})
}


func login(c *gin.Context) {
    userName := c.PostForm("username")
    password := c.PostForm("password")

    userId := service.GetUserId(userName, password)

    if userId == -1 {
        c.HTML(302, "signup.html", gin.H{
        	"error":"ユーザーが存在しません",
        })
        c.Abort()
        return
    }

    if userId == -2 {
        c.HTML(302, "signup.html", gin.H{
        	"error":"パスワードが間違っています",
        })
        c.Abort()
        return
    }
    
    //ほぼ出ないエラーはエラーハンドリングせず握りつぶす。
    jwtStr, _ := service.CreateJWT(userId, userName)

    // localhostは定数にする、変更に対応しやすくするため。
    c.SetCookie("userToken", jwtStr, 86460, "/", "localhost", false, true)
    c.Redirect(303, "/")
}


func logout(c *gin.Context) {
    // localhostは定数にする、変更に対応しやすくするため。
    c.SetCookie("userToken", "", 0, "/", "localhost", false, true)
    c.Redirect(303, "/login")
}

