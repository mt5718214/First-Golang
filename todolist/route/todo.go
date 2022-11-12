package route

import (
	"go-demo/todolist/api"

	"github.com/gin-gonic/gin"
)

func TodoRouter(rg *gin.RouterGroup) {
	todos := rg.Group("/todos")

	todos.GET("", api.GetTodoLists)
	todos.GET("/:id", api.GetTodoList)
	todos.POST("", api.PostTodo)
	todos.PUT("/:id", api.PutTodo)
	todos.DELETE("/:id", api.DeleteTodo)
}
