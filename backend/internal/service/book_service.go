package service

import (
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/domain"
)

type BookService interface {
	AddBook(book *domain.Book) error
	GetBook(id int64) (*domain.Book, error)
	UpdateBook(book *domain.Book) error
	DeleteBook(id int64) error
	ListBooks(search string, category string, page int, limit int) ([]*domain.Book, int64, error)
}

type bookService struct {
	repo domain.BookRepository
}

func NewBookService(repo domain.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) AddBook(book *domain.Book) error {
	return s.repo.Create(book)
}

func (s *bookService) GetBook(id int64) (*domain.Book, error) {
	return s.repo.GetByID(id)
}

func (s *bookService) UpdateBook(book *domain.Book) error {
	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id int64) error {
	return s.repo.Delete(id)
}

func (s *bookService) ListBooks(search string, category string, page int, limit int) ([]*domain.Book, int64, error) {
	offset := (page - 1) * limit
	return s.repo.List(search, category, limit, offset)
}
