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

type PostTodoRequestBody struct {
	Title   string
	Content string
}

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
	v1 := server.Group("/dev/api/v1")
	{
		v1.GET("/todos", getTodoLists)
		v1.GET("/todos/:id", getTodoList)
		v1.POST("/todos", postTodo)
		v1.PUT("/todos/:id", putTodo)
		v1.DELETE("/todos/:id", deleteTodo)
	}

	// By default it serves on :8080 unless a PORT environment variable was defined.
	// router.Run(":3000") for a hard coded port
	server.Run()
}

func getTodoLists(c *gin.Context) {
	query := "SELECT title, content FROM todo"
	rows, err := db.QueryContext(c, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	todos := make([]PostTodoRequestBody, 0)

	for rows.Next() {
		var todo PostTodoRequestBody
		if err := rows.Scan(&todo.Title, &todo.Content); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, todos)
}

func getTodoList(c *gin.Context) {
	id := c.Param("id")
	query := "SELECT title, content FROM todo WHERE id = ?"
	row := db.QueryRow(query, id)

	var todo PostTodoRequestBody
	err := row.Scan(&todo.Title, &todo.Content)
	if err != nil {
		fmt.Println("get todo error", err.Error())
	}

	c.JSON(200, todo)
}

func postTodo(c *gin.Context) {
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

	c.JSON(201, nil)
}

func putTodo(c *gin.Context) {
	id := c.Param("id")
	var requestBody PostTodoRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal("error", err)
	}

	// TODO: 修改更update邏輯, `title or content無值時會是空白字串`
	query := "UPDATE todo SET title = IFNULL(?, title), content = IFNULL(?, content) WHERE id = ?"
	_, err := db.Exec(query, requestBody.Title, requestBody.Content, id)
	if err != nil {
		fmt.Println("update todo err:", err.Error())
	}

	c.JSON(204, nil)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	query := "DELETE FROM todo WHERE id = ?"
	_, err := db.ExecContext(c, query, id)
	if err != nil {
		log.Fatal(err)
	}
}
