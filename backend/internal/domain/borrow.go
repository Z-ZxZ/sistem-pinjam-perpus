package domain

import (
	"time"
)

type BorrowStatus string

const (
	StatusBorrowed BorrowStatus = "borrowed"
	StatusReturned BorrowStatus = "returned"
	StatusOverdue  BorrowStatus = "overdue"
)

type Borrow struct {
	ID         int64        `json:"id"`
	UserID     int64        `json:"user_id"`
	BookID     int64        `json:"book_id"`
	BorrowDate time.Time    `json:"borrow_date"`
	DueDate    time.Time    `json:"due_date"`
	ReturnDate *time.Time   `json:"return_date,omitempty"`
	Status     BorrowStatus `json:"status"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`

	// Data hasil join nih asik dah
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Book *Book `json:"book,omitempty" gorm:"foreignKey:BookID"`
}

type BorrowRepository interface {
	Create(borrow *Borrow) error
	GetByID(id int64) (*Borrow, error)
	Update(borrow *Borrow) error
	ListByUser(userID int64) ([]*Borrow, error)
	ListAll() ([]*Borrow, error)
	GetActiveBorrowCount(userID int64) (int, error)
}
