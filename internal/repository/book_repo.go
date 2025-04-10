package repository

import (
	"github.com/bayskie/test-book-api/internal/model"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]model.Book, error)
	FindByID(id uint) (model.Book, error)
	Search(title, author string, year int) ([]model.Book, error)
	Create(book model.Book) (model.Book, error)
	Update(book model.Book) (model.Book, error)
	Delete(book model.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) FindByID(id uint) (model.Book, error) {
	var book model.Book
	err := r.db.First(&book, id).Error
	return book, err
}

func (r *bookRepository) Search(title, author string, year int) ([]model.Book, error) {
	var books []model.Book
	query := r.db

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if author != "" {
		query = query.Where("author LIKE ?", "%"+author+"%")
	}
	if year > 0 {
		query = query.Where("year = ?", year)
	}

	err := query.Find(&books).Error
	return books, err
}

func (r *bookRepository) Create(book model.Book) (model.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *bookRepository) Update(book model.Book) (model.Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *bookRepository) Delete(book model.Book) error {
	return r.db.Delete(&book).Error
}
