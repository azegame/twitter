package controller

import (
	"fmt"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"twitter/model/repository"
)


func follow(c *gin.Context) {
	claims := c.Keys["claims"]
	userIdByJWT := int(claims.(jwt.MapClaims)["user_id"].(float64))
	otherUserId, _ := strconv.Atoi(c.PostForm("followeeid"))

	err := repository.InsertFollowInfo(userIdByJWT, otherUserId)
	if err != nil {
		fmt.Println(err)
	}

	c.Redirect(303, "/")
}


func unFollow(c *gin.Context) {
	claims := c.Keys["claims"]
	userIdByJWT := int(claims.(jwt.MapClaims)["user_id"].(float64))
	otherUserId, _ := strconv.Atoi(c.PostForm("followeeid"))

	err := repository.DeleteFollowInfo(userIdByJWT, otherUserId)
	if err != nil {
		fmt.Println(err)
	}

	c.Redirect(303, "/")
}