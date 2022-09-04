package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Title   string
	Content string
}

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

func main() {
	server := gin.Default()

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
