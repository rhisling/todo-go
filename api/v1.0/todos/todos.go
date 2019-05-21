package todos

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.RouterGroup) {
	todos := r.Group("/todos")
	{
		todos.POST("/", createTodo)
		todos.GET("/", fetchAllTodo)
		todos.GET("/:id", fetchSingleTodo)
		todos.PUT("/:id", updateTodo)
		todos.DELETE("/:id", deleteTodo)
	}
}