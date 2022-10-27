package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// router
	server.GET("/todos", getTodoLists)
	server.GET("/todos/:id", getTodoList)
	server.POST("/todos", postTodo)
	server.PUT("/todos", putTodo)
	server.DELETE("/todos", deleteTodo)

	// By default it serves on :8080 unless a PORT environment variable was defined.
	// router.Run(":3000") for a hard coded port
	server.Run()
}

func getTodoLists(c *gin.Context) {
	fmt.Println("list all todos")
}

func getTodoList(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("list one todo, ID is: %v \n", id)
}

func postTodo(c *gin.Context) {
	fmt.Println("post a todo")
}

func putTodo(c *gin.Context) {
	fmt.Println("put a todo")
}

func deleteTodo(c *gin.Context) {
	fmt.Println("delete a todo")
}
