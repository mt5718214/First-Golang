package main

import (
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

func main() {
	server := gin.Default()

	// set the path to read HTML file
	server.LoadHTMLGlob("template/*")

	// route
	server.GET("/", test)

	// listen and serve on 0.0.0.0:8080
	server.Run(":8080")
}
