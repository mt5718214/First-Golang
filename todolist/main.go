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
	db, err = sql.Open("mysql", "demo:demo123@tcp(localhost:3306)/demo")
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

func postTodo(c *gin.Context) {
	type PostTodoRequestBody struct {
		Title   string
		Content string
	}
	var requestBody PostTodoRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal("error", err)
	}

	query := "INSERT INTO todo (title, content, is_complete) VALUES (?, ?, 0)"
	result, err := db.Exec(query, requestBody.Title, requestBody.Content)
	if err != nil {
		log.Fatal("insert todo err", err.Error())
	}
	if row, _ := result.RowsAffected(); row != 1 {
		log.Fatal("insert todo data unmatched.")
	}
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
