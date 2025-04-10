package main

import (
	"github.com/bayskie/test-book-api/internal/config"
	"github.com/bayskie/test-book-api/internal/handler"
	"github.com/bayskie/test-book-api/internal/repository"
	"github.com/bayskie/test-book-api/internal/seed"
	"github.com/bayskie/test-book-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.NewDatabase()

	seed.SeedBooks(db)

	bookRepo := repository.NewBookRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookHandler := handler.NewBookHandler(bookUsecase)

	r := gin.Default()
	bookHandler.RegisterRoutes(r)

	r.Run(":8080")
}
