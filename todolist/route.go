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
		v1.POST("/login", api.AuthHandler)

		v1.GET("/todos", api.JWTAuthMiddleware(), api.GetTodoLists)
		v1.GET("/todos/:id", api.JWTAuthMiddleware(), api.GetTodoList)
		v1.POST("/todos", api.JWTAuthMiddleware(), api.PostTodo)
		v1.PUT("/todos/:id", api.JWTAuthMiddleware(), api.PutTodo)
		v1.DELETE("/todos/:id", api.JWTAuthMiddleware(), api.DeleteTodo)
	}

	return server
}
