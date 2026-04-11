package repository

import (
	"database/sql"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/domain"
)

type borrowRepositoryPostgres struct {
	db *sql.DB
}

func NewBorrowRepositoryPostgres(db *sql.DB) domain.BorrowRepository {
	return &borrowRepositoryPostgres{db: db}
}

func (r *borrowRepositoryPostgres) Create(borrow *domain.Borrow) error {
	query := `INSERT INTO borrows (user_id, book_id, borrow_date, due_date, status, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(query, borrow.UserID, borrow.BookID, borrow.BorrowDate, borrow.DueDate, borrow.Status).
		Scan(&borrow.ID, &borrow.CreatedAt, &borrow.UpdatedAt)
}

func (r *borrowRepositoryPostgres) GetByID(id int64) (*domain.Borrow, error) {
	borrow := &domain.Borrow{}
	query := `SELECT id, user_id, book_id, borrow_date, due_date, return_date, status, created_at, updated_at FROM borrows WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&borrow.ID, &borrow.UserID, &borrow.BookID, &borrow.BorrowDate, &borrow.DueDate, &borrow.ReturnDate, &borrow.Status, &borrow.CreatedAt, &borrow.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return borrow, nil
}

func (r *borrowRepositoryPostgres) Update(borrow *domain.Borrow) error {
	query := `UPDATE borrows SET user_id = $1, book_id = $2, borrow_date = $3, due_date = $4, return_date = $5, status = $6, updated_at = NOW() WHERE id = $7`
	_, err := r.db.Exec(query, borrow.UserID, borrow.BookID, borrow.BorrowDate, borrow.DueDate, borrow.ReturnDate, borrow.Status, borrow.ID)
	return err
}

func (r *borrowRepositoryPostgres) ListByUser(userID int64) ([]*domain.Borrow, error) {
	query := `SELECT b.id, b.user_id, b.book_id, b.borrow_date, b.due_date, b.return_date, b.status, b.created_at, b.updated_at,
			  bk.title, bk.author
			  FROM borrows b
			  JOIN books bk ON b.book_id = bk.id
			  WHERE b.user_id = $1
			  ORDER BY b.borrow_date DESC`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var borrows []*domain.Borrow
	for rows.Next() {
		b := &domain.Borrow{Book: &domain.Book{}}
		err := rows.Scan(&b.ID, &b.UserID, &b.BookID, &b.BorrowDate, &b.DueDate, &b.ReturnDate, &b.Status, &b.CreatedAt, &b.UpdatedAt, &b.Book.Title, &b.Book.Author)
		if err != nil {
			return nil, err
		}
		borrows = append(borrows, b)
	}
	return borrows, nil
}

func (r *borrowRepositoryPostgres) ListAll() ([]*domain.Borrow, error) {
	query := `SELECT b.id, b.user_id, b.book_id, b.borrow_date, b.due_date, b.return_date, b.status, b.created_at, b.updated_at,
			  u.name, bk.title
			  FROM borrows b
			  JOIN users u ON b.user_id = u.id
			  JOIN books bk ON b.book_id = bk.id
			  ORDER BY b.borrow_date DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var borrows []*domain.Borrow
	for rows.Next() {
		b := &domain.Borrow{User: &domain.User{}, Book: &domain.Book{}}
		err := rows.Scan(&b.ID, &b.UserID, &b.BookID, &b.BorrowDate, &b.DueDate, &b.ReturnDate, &b.Status, &b.CreatedAt, &b.UpdatedAt, &b.User.Name, &b.Book.Title)
		if err != nil {
			return nil, err
		}
		borrows = append(borrows, b)
	}
	return borrows, nil
}

func (r *borrowRepositoryPostgres) GetActiveBorrowCount(userID int64) (int, error) {
	query := `SELECT COUNT(*) FROM borrows WHERE user_id = $1 AND status != 'returned'`
	var count int
	err := r.db.QueryRow(query, userID).Scan(&count)
	return count, err
}
