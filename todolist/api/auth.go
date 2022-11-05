package api

import (
	"fmt"
	db "go-demo/todolist/database"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type userInfoReqBody struct {
	Username, Password, CheckPassword string
}

type userInfo struct {
	Id                 int
	Username, Password string
}

func Auth(c *gin.Context) {
	var userInfoReqBody userInfoReqBody
	err := c.BindJSON(&userInfoReqBody)
	if err != nil {
		fmt.Println("BindJSON error: ", err.Error())
	}

	username := strings.Trim(userInfoReqBody.Username, " ")
	password := strings.Trim(userInfoReqBody.Password, " ")
	checkPassword := strings.Trim(userInfoReqBody.CheckPassword, " ")

	if username == "" || password == "" || checkPassword == "" {
		c.JSON(400, gin.H{
			"message": "field can't be empty",
		})
	}

	if password != checkPassword {
		c.JSON(400, gin.H{
			"message": "password and checkPassword is not equal",
		})
	}

	row := db.SqlDB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	var userInfo userInfo
	err = row.Scan(&userInfo.Id, &userInfo.Username, &userInfo.Password)
	if err != nil {
		fmt.Println("QueryRow error: ", err.Error())
		c.JSON(400, gin.H{
			"message": "user is not exist",
		})
	}

	// sign JWT token and return to client
	token, err := createJWT("token", userInfo.Id, userInfo.Username)
	if err != nil {
		fmt.Println("createJWT error: ", err.Error())
		c.JSON(400, gin.H{
			"token": "",
		})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

// type JWTClaim struct {
// 	*jwt.RegisteredClaims
// 	UserInfo interface{}
// }

func createJWT(sub string, userId int, username string) (string, error) {
	// Secret key
	mySigningKey := []byte("mySigningKey")
	// Create the Claims
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	claim := jwt.RegisteredClaims{
		Issuer:    "kemp",
		Subject:   sub,
		Audience:  []string{username},
		ExpiresAt: jwt.NewNumericDate(expireTime),
		IssuedAt:  jwt.NewNumericDate(nowTime),
		ID:        strconv.Itoa(userId),
	}

	// token instance
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("generate token fail: ", err.Error())
		return "", err
	}

	return ss, err
}

func ParseToken(c *gin.Context, token string) {
}
