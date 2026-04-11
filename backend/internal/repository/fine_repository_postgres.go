package repository

import (
	"database/sql"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/domain"
)

type fineRepositoryPostgres struct {
	db *sql.DB
}

func NewFineRepositoryPostgres(db *sql.DB) domain.FineRepository {
	return &fineRepositoryPostgres{db: db}
}

func (r *fineRepositoryPostgres) Create(fine *domain.Fine) error {
	query := `INSERT INTO fines (borrow_id, user_id, amount, status, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(query, fine.BorrowID, fine.UserID, fine.Amount, fine.Status).
		Scan(&fine.ID, &fine.CreatedAt, &fine.UpdatedAt)
}

func (r *fineRepositoryPostgres) GetByID(id int64) (*domain.Fine, error) {
	fine := &domain.Fine{}
	query := `SELECT id, borrow_id, user_id, amount, status, created_at, updated_at FROM fines WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&fine.ID, &fine.BorrowID, &fine.UserID, &fine.Amount, &fine.Status, &fine.CreatedAt, &fine.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return fine, nil
}

func (r *fineRepositoryPostgres) GetByBorrowID(borrowID int64) (*domain.Fine, error) {
	fine := &domain.Fine{}
	query := `SELECT id, borrow_id, user_id, amount, status, created_at, updated_at FROM fines WHERE borrow_id = $1`
	err := r.db.QueryRow(query, borrowID).Scan(&fine.ID, &fine.BorrowID, &fine.UserID, &fine.Amount, &fine.Status, &fine.CreatedAt, &fine.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return fine, nil
}

func (r *fineRepositoryPostgres) Update(fine *domain.Fine) error {
	query := `UPDATE fines SET amount = $1, status = $2, updated_at = NOW() WHERE id = $3`
	_, err := r.db.Exec(query, fine.Amount, fine.Status, fine.ID)
	return err
}

func (r *fineRepositoryPostgres) ListByUser(userID int64) ([]*domain.Fine, error) {
	query := `SELECT id, borrow_id, user_id, amount, status, created_at, updated_at FROM fines WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fines []*domain.Fine
	for rows.Next() {
		f := &domain.Fine{}
		err := rows.Scan(&f.ID, &f.BorrowID, &f.UserID, &f.Amount, &f.Status, &f.CreatedAt, &f.UpdatedAt)
		if err != nil {
			return nil, err
		}
		fines = append(fines, f)
	}
	return fines, nil
}

func (r *fineRepositoryPostgres) GetTotalUnpaid(userID int64) (float64, error) {
	query := `SELECT COALESCE(SUM(amount), 0) FROM fines WHERE user_id = $1 AND status = 'unpaid'`
	var total float64
	err := r.db.QueryRow(query, userID).Scan(&total)
	return total, err
}
