package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func main() {
	server := gin.Default()

	// db connection
	var err error
	db, err = sql.Open("mysql", "demo:demo123@tcp(localhost:3307)/demo")
	if err != nil {
		fmt.Println("DB連線資訊有誤請再次確認")
	}
	if err := db.Ping(); err != nil {
		fmt.Println("開啟 MySQL 連線發生錯誤，原因為：", err.Error())
	}

	// router
	server.GET("/todos", getTodoLists)
	server.GET("/todos/:id", getTodoList)
	server.POST("/todos", postTodo)
	server.PUT("/todos", putTodo)
	server.DELETE("/todos/:id", deleteTodo)

	// By default it serves on :8080 unless a PORT environment variable was defined.
	// router.Run(":3000") for a hard coded port
	server.Run()
}

func getTodoLists(c *gin.Context) {
	query := "SELECT title FROM todo"
	rows, err := db.QueryContext(c, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println(rows.Columns())
	}
	fmt.Println(rows)
	fmt.Println("list all todos")
}

func getTodoList(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("list one todo, ID is: %v \n", id)
}

type EmailRequestBody struct {
	Gin string
}

func postTodo(c *gin.Context) {
	// var map = map[string]string
	// err := c.Bind(&map)
	var requestBody EmailRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal("error", err)
		return
	}
	// body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("body:", requestBody)
	// query := "INSERT INTO todo (title, content, is_complete) VALUES (?, ?, ?)"
	// db.Exec(query)
	fmt.Println("post a todo")
}

func putTodo(c *gin.Context) {
	fmt.Println("put a todo")
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	query := "DELETE FROM todo WHERE id = ?"
	_, err := db.ExecContext(c, query, id)
	if err != nil {
		log.Fatal(err)
	}
}
