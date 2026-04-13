package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/domain"
)

type bookRepositoryPostgres struct {
	db *sql.DB
}

func NewBookRepositoryPostgres(db *sql.DB) domain.BookRepository {
	return &bookRepositoryPostgres{db: db}
}

func (r *bookRepositoryPostgres) Create(book *domain.Book) error {
	query := `INSERT INTO books (title, author, publisher, year, category, isbn, stock, cover_url, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW()) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(query, book.Title, book.Author, book.Publisher, book.Year, book.Category, book.ISBN, book.Stock, book.CoverURL).
		Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
}

func (r *bookRepositoryPostgres) GetByID(id int64) (*domain.Book, error) {
	book := &domain.Book{}
	query := `SELECT id, title, author, publisher, year, category, isbn, stock, cover_url, created_at, updated_at FROM books WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Publisher, &book.Year, &book.Category, &book.ISBN, &book.Stock, &book.CoverURL, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *bookRepositoryPostgres) Update(book *domain.Book) error {
	query := `UPDATE books SET title = $1, author = $2, publisher = $3, year = $4, category = $5, isbn = $6, stock = $7, cover_url = $8, updated_at = NOW() WHERE id = $9`
	_, err := r.db.Exec(query, book.Title, book.Author, book.Publisher, book.Year, book.Category, book.ISBN, book.Stock, book.CoverURL, book.ID)
	return err
}

func (r *bookRepositoryPostgres) Delete(id int64) error {
	query := `DELETE FROM books WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *bookRepositoryPostgres) List(search string, category string, limit int, offset int) ([]*domain.Book, int64, error) {
	var whereClauses []string
	var args []interface{}
	argCount := 1

	if search != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("(LOWER(title) LIKE LOWER($%d) OR LOWER(author) LIKE LOWER($%d) OR isbn LIKE $%d)", argCount, argCount, argCount))
		args = append(args, "%"+search+"%")
		argCount++
	}

	if category != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("category = $%d", argCount))
		args = append(args, category)
		argCount++
	}

	whereSQL := ""
	if len(whereClauses) > 0 {
		whereSQL = " WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Ngitung total data cuy
	countQuery := "SELECT COUNT(*) FROM books" + whereSQL
	var total int64
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// List datanya nih
	listQuery := fmt.Sprintf("SELECT id, title, author, publisher, year, category, isbn, stock, cover_url, created_at, updated_at FROM books%s LIMIT $%d OFFSET $%d", whereSQL, argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var books []*domain.Book
	for rows.Next() {
		book := &domain.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Publisher, &book.Year, &book.Category, &book.ISBN, &book.Stock, &book.CoverURL, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}
		books = append(books, book)
	}

	return books, total, nil
}
