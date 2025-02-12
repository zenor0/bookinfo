package service

import (
	"errors"

	"github.com/zenor0/bookinfo/services/details/internal/model"
	"github.com/zenor0/bookinfo/services/details/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(book *model.Book) error {
	if book.ISBN == "" {
		return errors.New("isbn is required")
	}
	if book.Title == "" {
		return errors.New("title is required")
	}
	if book.Author == "" {
		return errors.New("author is required")
	}
	return s.repo.Create(book)
}

func (s *BookService) GetBook(id uint) (*model.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) GetBookByISBN(isbn string) (*model.Book, error) {
	if isbn == "" {
		return nil, errors.New("isbn is required")
	}
	return s.repo.GetByISBN(isbn)
}

func (s *BookService) ListBooks(page, pageSize int) ([]model.Book, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	return s.repo.List(page, pageSize)
}

func (s *BookService) UpdateBook(book *model.Book) error {
	if book.ID == 0 {
		return errors.New("book id is required")
	}
	return s.repo.Update(book)
}

func (s *BookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}

func (s *BookService) SearchBooks(query string) ([]model.Book, error) {
	if query == "" {
		return nil, errors.New("search query is required")
	}
	return s.repo.Search(query)
}
