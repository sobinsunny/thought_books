package controllers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"thought_books/app/models"
)

var DB *gorm.DB

func InitDB() {
	Gorm, err := gorm.Open("postgres", "host=localhost  dbname=thought_books_dev_db sslmode=disable")
	Gorm.LogMode(true)

	if err != nil {
		panic("failed to connect database")
	}

	Gorm.AutoMigrate(&models.Post{})
}
