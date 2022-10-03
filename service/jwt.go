package service

import (
	"fmt"
	"time"
	"errors"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/golang-jwt/jwt/v4"

	//"twitter/model/entity"
	//"twitter/model/repository"
)


func CreateJWT(userId int, userName string) (string, error) {
	//Claimsオブジェクトの作成
	claims := jwt.MapClaims{
		"user_id" : userId,
		"user_name" : userName,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	}
	//ヘッダー(暗号化方式を任意で指定、HMAC)とClaimsからtokenの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//トークンに署名を付与
	signedToken, err:= token.SignedString([]byte("SECRET_KEY"))
	return signedToken, err
}


func AuthorizationMiddleware()  gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := Authorization(c)

		if err != nil {
			fmt.Printf("AuthorizationMiddleware()のエラー %v\n", err)
			c.Redirect(303, "/login")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

//Cookie内のJWTからトークンを取得して、そのトークンが正しいか検証、正しかったらClaimsを返す
func Authorization(c *gin.Context) (jwt.MapClaims, error) {
	signedToken, err := c.Cookie("userToken")
	if err != nil {
		return nil, err
	}

	token, err := validateToken(signedToken)
	if err != nil {
		//fmt.Printf("Authorization()のtokenは %v\n", token)
		return nil, err
	}

	if claims, ok := extractClaims(token); ok {
		return claims, nil
	}

	return nil, errors.New("Claimsを取得できませんでした")
}


func validateToken(signedToken string) (*jwt.Token, error){
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("SECRET_KEY"), nil
	})

	return token, err
}


func extractClaims(token *jwt.Token) (jwt.MapClaims, bool) {
	//tokenのClaimsの型はjwt.MapClaimsと型アサーションしている
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		//log.Printf("Invalid JWT Token")
		return nil, false
	}
}

