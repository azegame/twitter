package service

import (
	"fmt"

	//"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"twitter/model/repository"
)


// 数字で返す意味、2パターンならboolの方がいい
func Signup(userName string, password string) bool {
	if userExists(userName) {
		return false
	}
	ep, _ := encryptPassword(password)

	//空文字処理

	err := repository.CreateUser(userName, ep)
	if err != nil {
		fmt.Println(err)
	}

	return true
}


func userExists(userName string) bool {
	// ユーザーが見つからなければエラー(重複していない時)
	_, err := repository.FindUserByUserName(userName)
	if err != nil {
		return false
	}
	return true
}


func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func GetUserId(userName string, password string) int {
	user, err := repository.FindUserByUserName(userName)

	if err != nil {
		fmt.Println(err)
		return -1
	}
	if CompareHashAndPassword(user.Password, password) != nil {
		return -2
	}

	fmt.Printf("GetUserId()のuser.UIdは、%d\n", user.UserId)
	return user.UserId
}


// hashと非hashパスワード比較(ログイン用)
func CompareHashAndPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	//fmt.Println(err)    //hashedPassword is not the hash of the given password
	return err
}










