package usecase

import (
	"github.com/bayskie/test-book-api/internal/model"
	"github.com/bayskie/test-book-api/internal/repository"
)

type BookUsecase interface {
	GetAll() ([]model.Book, error)
	GetByID(id uint) (model.Book, error)
	Search(title, author string, year int) ([]model.Book, error)
	Create(book model.Book) (model.Book, error)
	Update(book model.Book) (model.Book, error)
	Delete(id uint) error
}

type bookUsecase struct {
	repo repository.BookRepository
}

func NewBookUsecase(r repository.BookRepository) BookUsecase {
	return &bookUsecase{r}
}

func (u *bookUsecase) GetAll() ([]model.Book, error) {
	return u.repo.FindAll()
}

func (u *bookUsecase) GetByID(id uint) (model.Book, error) {
	return u.repo.FindByID(id)
}

func (u *bookUsecase) Search(title, author string, year int) ([]model.Book, error) {
	return u.repo.Search(title, author, year)
}

func (u *bookUsecase) Create(book model.Book) (model.Book, error) {
	return u.repo.Create(book)
}

func (u *bookUsecase) Update(book model.Book) (model.Book, error) {
	return u.repo.Update(book)
}

func (u *bookUsecase) Delete(id uint) error {
	book, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	return u.repo.Delete(book)
}
