package domain

import (
	"time"
)

type Book struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	Year        int       `json:"year"`
	Category    string    `json:"category"`
	ISBN        string    `json:"isbn"`
	Stock       int       `json:"stock"`
	CoverURL    string    `json:"cover_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookRepository interface {
	Create(book *Book) error
	GetByID(id int64) (*Book, error)
	Update(book *Book) error
	Delete(id int64) error
	List(search string, category string, limit int, offset int) ([]*Book, int64, error)
}
