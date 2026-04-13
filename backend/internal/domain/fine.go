package domain

import (
	"time"
)

type FineStatus string

const (
	FineUnpaid FineStatus = "unpaid"
	FinePaid   FineStatus = "paid"
)

type Fine struct {
	ID        int64      `json:"id"`
	BorrowID  int64      `json:"borrow_id"`
	UserID    int64      `json:"user_id"`
	Amount    float64    `json:"amount"`
	Status    FineStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	// Data hasil join bro
	Borrow *Borrow `json:"borrow,omitempty"`
}

type FineRepository interface {
	Create(fine *Fine) error
	GetByID(id int64) (*Fine, error)
	GetByBorrowID(borrowID int64) (*Fine, error)
	Update(fine *Fine) error
	ListByUser(userID int64) ([]*Fine, error)
	GetTotalUnpaid(userID int64) (float64, error)
}
