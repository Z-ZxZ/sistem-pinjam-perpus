package domain

import (
	"time"
)

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleMember UserRole = "member"
)

type User struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Role        UserRole  `json:"role"`
	Status      string    `json:"status"` // active, inactive
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id int64) error
	List() ([]*User, error)
}
