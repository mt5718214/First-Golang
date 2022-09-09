package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type IndexData struct {
	Title   string
	Content string
}

type User struct {
	id                 int
	username, password string
}

const (
	USERNAME = "demo"
	PASSWORD = "demo123"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "golangtest"
)

func test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一支 gin 專案"
	c.HTML(http.StatusOK, "index.html", data)
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginAuth(c *gin.Context) {
	var username, password string

	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入使用者名稱"),
		})
		return
	}

	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入密碼"),
		})
		return
	}

	if error := Auth(username, password); error == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "登入成功",
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": error,
		})
		return
	}
}

func selectUser(DB *sql.DB, username string) {
	user := new(User)
	stmtOut, err := DB.Prepare("SELECT * FROM users WHERE username = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	err = stmtOut.QueryRow(username).Scan(&user.id, &user.username, &user.password)
	if err != nil {
		fmt.Println("select user error", err)
		return
	}
	fmt.Println("users", *user)
}

func insertUser(DB *sql.DB, username, password string) error {
	_, err := DB.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, username, password)
	if err != nil {
		fmt.Println("err", err)
		return errors.New("insert users error")
	}
	return nil
}

func main() {
	server := gin.Default()

	// db connection
	db, err := sql.Open("mysql", "root:password@/demo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err != nil {
		fmt.Println("開啟 MySQL 連線發生錯誤，原因為：", err)
		return
	}

	// Open doesn't open a connection. Validate DSN data:
	if err = db.Ping(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	insertUser(db, "user2", "password")
	selectUser(db, "user2")

	// set the path to read HTML file
	server.LoadHTMLGlob("template/html/*")
	// set the path to read Static file
	server.Static("/assets", "./template/assets")

	// route
	server.GET("/", test)
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth)

	// listen and serve on 0.0.0.0:8080
	server.Run(":8080")
}
