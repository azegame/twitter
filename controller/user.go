package controller

import (
    //"fmt"

    "github.com/gin-gonic/gin"
    //_ "github.com/mattn/go-sqlite3"

    "twitter/service" 
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

    jwtStr, _ := service.CreateJWT(userId, userName)

    /* ほぼ出ないエラーはエラーハンドリングせず握りつぶす。
    if err != nil {
        c.HTML(401, "index.html", gin.H{"error": "tokenの作成に失敗しました。"})
        c.Abort()
        return
    }
    */

    // localhostは定数にする、変更に対応しやすくするため。
    c.SetCookie("userToken", jwtStr, 86460, "/", "localhost", false, true)
    c.Redirect(303, "/")
}


func logout(c *gin.Context) {
    // localhostは定数にする、変更に対応しやすくするため。
    c.SetCookie("userToken", "", 0, "/", "localhost", false, true)
    c.Redirect(303, "/login")
}


