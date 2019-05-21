package todos

import (
	"fmt"
	"github.com/aravindh/todoApp/database/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// type alias for models
type TodoModel = models.TodoModel
type TransformedTodoModel = models.TransformedTodo

func createTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todo TodoModel
	e := c.BindJSON(&todo)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("Todo title" + todo.Title)
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.Title})
}

func fetchAllTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todos []TodoModel
	var _todos []TransformedTodoModel
	db.Find(&todos)
	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	//transforms the todos for building a good response
	for _, item := range todos {

		_todos = append(_todos, TransformedTodoModel{ID: item.ID, Title: item.Title, Completed: item.Completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

func fetchSingleTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todo TodoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	_todo := TransformedTodoModel{ID: todo.ID, Title: todo.Title, Completed: todo.Completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

func updateTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todoNew TodoModel
	var todoOld TodoModel
	e := c.BindJSON(&todoNew)
	if e != nil {
		fmt.Println(e)
		return
	}
	todoID := c.Param("id")
	db.First(&todoOld, todoID)
	if todoOld.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	db.Model(&todoOld).Update(&todoNew)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

func deleteTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todo TodoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
