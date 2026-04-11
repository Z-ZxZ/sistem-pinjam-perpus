package repository

import (
	"database/sql"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/domain"
)

type userRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) domain.UserRepository {
	return &userRepositoryPostgres{db: db}
}

func (r *userRepositoryPostgres) Create(user *domain.User) error {
	query := `INSERT INTO users (name, email, password, role, status, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(query, user.Name, user.Email, user.Password, user.Role, user.Status).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *userRepositoryPostgres) GetByID(id int64) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, name, email, password, role, status, created_at, updated_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryPostgres) GetByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, name, email, password, role, status, created_at, updated_at FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryPostgres) Update(user *domain.User) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3, role = $4, status = $5, updated_at = NOW() WHERE id = $6`
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.Role, user.Status, user.ID)
	return err
}

func (r *userRepositoryPostgres) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *userRepositoryPostgres) List() ([]*domain.User, error) {
	query := `SELECT id, name, email, role, status, created_at, updated_at FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
