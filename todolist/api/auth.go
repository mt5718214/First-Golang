package api

import (
	"errors"
	"fmt"
	db "go-demo/todolist/database"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

var (
	// Secret key
	mySigningKey = []byte("mySigningKey")
)

type userInfoReqBody struct {
	Username, Password, CheckPassword string
}

type userInfo struct {
	Id                 int
	Username, Password string
}

// AuthHandler @Summary
// @version 1.0
// @produce application/json
// @param register body Login true "login"
// @Success 200 string successful return token
// @Router /dev/api/v1/login [post]
func AuthHandler(c *gin.Context) {
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

// 如果需要包含客製化使用者資訊的Claim可使用以下的struct
// type JWTClaim struct {
// 	*jwt.RegisteredClaims
// 	UserInfo interface{}
// }

func createJWT(sub string, userId int, username string) (string, error) {
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

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	if strings.Trim(tokenString, " ") == "" {
		return nil, errors.New("invalid token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		/**
		* Comma-ok 斷言
		* 可以直接判斷是否是該型別的變數： value, ok = element.(T)
		* value 就是變數的值，ok 是一個 bool 型別，element 是 interface 變數，T 是斷言的型別。
		* 如果 element 裡面確實儲存了 T 型別的數值，那麼 ok 回傳 true，否則回傳 false。
		 */
		// 驗證 alg 是否為預期的HMAC演算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println("parse token error: ", err.Error())
		return nil, err
	}

	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println("parsed token: ", token.Header, token.Signature, err)
		fmt.Println("claim", claim)
		return claim, nil
	}

	return nil, errors.New("invalid token")
}
