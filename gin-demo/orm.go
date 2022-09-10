package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME = "demo"
	PASSWORD = "demo123"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "demo"
)

// gorm.Model definition
// to avoid same name in main.go, here use UserORM.
type UserORM struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
	Username string `json:"username"`
	Password string `json:""`
}

var OrmDb *gorm.DB
var err error

func init() {
	fmt.Println("orm init")

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	OrmDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + err.Error())
	}
}

func CreateUser(db *gorm.DB, user *UserORM) error {
	/*
		Temporarily specify table name with Table method,
		because gorm will use the struct name as table name to insert,
		so use UserORM will throw error.
	*/
	return db.Table("users").Create(user).Error
}

func FindUser(db *gorm.DB, id int64) (*UserORM, error) {
	user := new(UserORM)
	user.ID = id
	error := db.Table("users").First(&user).Error
	return user, error
}

func FindUsers(db *gorm.DB, username interface{}) (*[]UserORM, error) {
	var error error
	var users = new([]UserORM)

	if username == nil {
		error = db.Table("users").Find(&users).Error
	} else {
		error = db.Table("users").Where("username = ?", username).Find(&users).Error
	}
	return users, error
}
