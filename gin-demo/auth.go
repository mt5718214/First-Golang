package main

import (
	"errors"
)

var UserData map[string]string

func init() {
	UserData = map[string]string{
		"test": "test",
	}
}

// 註: 沒串接DB 先單純用變數做判斷
func CheckUserIsExist(username string) bool {
	_, isExist := UserData[username]
	return isExist
}
func CheckPassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	}

	return errors.New("password is not correct")
}

func Auth(username string, password string) error {
	if isExist := CheckUserIsExist(username); isExist {
		return CheckPassword(password, UserData[username])
	}

	return errors.New("user is not exist")
}
