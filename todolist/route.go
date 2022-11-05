package main

import (
	api "go-demo/todolist/api"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	server := gin.Default()

	// router
	v1 := server.Group("/dev/api/v1")
	{
		v1.POST("/login", api.Auth)

		v1.GET("/todos", api.GetTodoLists)
		v1.GET("/todos/:id", api.GetTodoList)
		v1.POST("/todos", api.PostTodo)
		v1.PUT("/todos/:id", api.PutTodo)
		v1.DELETE("/todos/:id", api.DeleteTodo)
	}

	return server
}
