package service

import(
	"fmt"

	"twitter/model/repository"
)


func IsFollowing(userIdByJWT int, otherUserId int) bool {
	_, err := repository.SearchFollowInfo(userIdByJWT, otherUserId)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}