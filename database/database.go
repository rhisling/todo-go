package database

import (
	"github.com/aravindh/todoApp/database/models"
	"github.com/jinzhu/gorm"
)

func Initialize() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=havok dbname=havok sslmode=disable password=havok")
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.TodoModel{}, &models.TransformedTodo{})
	return db, err
}
