package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bayskie/test-book-api/internal/model"
	"github.com/bayskie/test-book-api/internal/usecase"
	"github.com/bayskie/test-book-api/pkg/formatter"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookHandler struct {
	Usecase  usecase.BookUsecase
	validate *validator.Validate
}

func NewBookHandler(u usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		Usecase:  u,
		validate: validator.New(),
	}
}

func (h *BookHandler) RegisterRoutes(r *gin.Engine) {
	books := r.Group("/books")
	{
		books.GET("", h.GetAll)
		books.GET("/:id", h.GetByID)
		books.GET("/search", h.Search)
		books.POST("", h.Create)
		books.PUT("/:id", h.Update)
		books.DELETE("/:id", h.Delete)
	}
}

func (h *BookHandler) GetAll(c *gin.Context) {
	books, err := h.Usecase.GetAll()
	c.JSON(http.StatusOK, formatter.NewResponse(books, err))
}

func (h *BookHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, formatter.NewResponse(nil, fmt.Errorf("invalid book ID")))
		return
	}

	book, err := h.Usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, formatter.NewResponse(nil, err))
		return
	}

	c.JSON(http.StatusOK, formatter.NewResponse(book, nil))
}

func (h *BookHandler) Search(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	yearStr := c.Query("year")

	var year int
	if yearStr != "" {
		y, err := strconv.Atoi(yearStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, formatter.NewResponse(nil, fmt.Errorf("invalid year")))
			return
		}
		year = y
	}

	result, err := h.Usecase.Search(title, author, year)
	c.JSON(http.StatusOK, formatter.NewResponse(result, err))
}

func (h *BookHandler) Create(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, formatter.NewResponse(nil, err))
		return
	}

	if err := h.validate.Struct(book); err != nil {
		errs := formatter.NewValidationErrors(err)
		c.JSON(http.StatusBadRequest, formatter.NewResponse(nil, fmt.Errorf("%s", *errs)))
		return
	}

	newBook, err := h.Usecase.Create(book)
	c.JSON(http.StatusCreated, formatter.NewResponse(newBook, err))
}

func (h *BookHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.Usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, formatter.NewResponse(nil, err))
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, formatter.NewResponse(nil, err))
		return
	}

	if err := h.validate.Struct(book); err != nil {
		errs := formatter.NewValidationErrors(err)
		c.JSON(http.StatusBadRequest, formatter.NewResponse(nil, fmt.Errorf("%s", *errs)))
		return
	}

	updated, err := h.Usecase.Update(book)
	c.JSON(http.StatusOK, formatter.NewResponse(updated, err))
}

func (h *BookHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, formatter.NewResponse(nil, fmt.Errorf("invalid book ID")))
		return
	}

	err = h.Usecase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, formatter.NewResponse(nil, err))
		return
	}

	c.JSON(http.StatusOK, formatter.NewResponse(gin.H{"message": "book deleted successfully"}, nil))
}
