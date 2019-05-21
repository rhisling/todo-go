package apiv1

import (
	"github.com/aravindh/todoApp/api/v1.0/todos"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		todos.ApplyRoutes(v1)
	}
}