package config

import (
	"log"

	"github.com/bayskie/test-book-api/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&model.Book{})
	return db
}
