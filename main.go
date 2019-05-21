package main

import (
	"github.com/aravindh/todoApp/api"
	"github.com/aravindh/todoApp/database"

	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	app := gin.Default()
	db, _ := database.Initialize()
	app.Use(database.Inject(db))
	api.ApplyRoutes(app)
	_ = app.Run(":8080")
}
