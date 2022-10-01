package controller

import (
    "fmt"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"

    "twitter/model/repository"
    "twitter/service"
)




func otherUserPage(c *gin.Context) {
    claims := c.Keys["claims"]
    userIdByJWT := int(claims.(jwt.MapClaims)["user_id"].(float64))
    otherUserId, _ := strconv.Atoi(c.Param("UserId"))

    user, _ := repository.FindUserByUserId(otherUserId)
    tweets := repository.GetTweet(otherUserId)

    if service.IsFollowing(userIdByJWT, otherUserId) {
        //followする
           c.HTML(200, "follow.html", gin.H{
            "username": user.UserName,
            "followid": userIdByJWT,
            "followeeid": otherUserId,
            "followstate": "unfollow",
            "tweets": tweets,
        })
    } else {
        //unfollowにする
        c.HTML(200, "follow.html", gin.H{
            "username": user.UserName,
            "followid": userIdByJWT,
            "followeeid": otherUserId,
            "followstate": "follow",
            "tweets": tweets,
        })
    }
    //tweetがまだありません処理
}


func follow(c *gin.Context) {
    claims := c.Keys["claims"]
    userIdByJWT := int(claims.(jwt.MapClaims)["user_id"].(float64))
    otherUserId, _ := strconv.Atoi(c.PostForm("followeeid"))

    //user, _ := repository.FindUserByUserId(otherUserId)
    //tweets := repository.GetTweet(otherUserId)

    err := repository.InsertFollowInfo(userIdByJWT, otherUserId)
    fmt.Println(err)

    c.Redirect(303, "/")
}


func unFollow(c *gin.Context) {
    claims := c.Keys["claims"]
    userIdByJWT := int(claims.(jwt.MapClaims)["user_id"].(float64))
    otherUserId, _ := strconv.Atoi(c.PostForm("followeeid"))

    //user, _ := repository.FindUserByUserId(otherUserId)
    //tweets := repository.GetTweet(otherUserId)

    err := repository.DeleteFollowInfo(userIdByJWT, otherUserId)
    fmt.Println(err)

    c.Redirect(303, "/")
}




