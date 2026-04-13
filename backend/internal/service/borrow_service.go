package service

import (
	"errors"
	"math"
	"time"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/domain"
)

type BorrowService interface {
	BorrowBook(userID int64, bookID int64) error
	ReturnBook(borrowID int64) (float64, error)
	GetUserHistory(userID int64) ([]*domain.Borrow, error)
	ListAllBorrows() ([]*domain.Borrow, error)
	GetUserFines(userID int64) ([]*domain.Fine, float64, error)
}

type borrowService struct {
	borrowRepo domain.BorrowRepository
	bookRepo   domain.BookRepository
	fineRepo   domain.FineRepository
}

func NewBorrowService(borrowRepo domain.BorrowRepository, bookRepo domain.BookRepository, fineRepo domain.FineRepository) BorrowService {
	return &borrowService{
		borrowRepo: borrowRepo,
		bookRepo:   bookRepo,
		fineRepo:   fineRepo,
	}
}

func (s *borrowService) BorrowBook(userID int64, bookID int64) error {
	// Cek dlu dia punya denda yg belom dibayar ga, jgn ngutang terus bos
	totalUnpaid, err := s.fineRepo.GetTotalUnpaid(userID)
	if err != nil {
		return err
	}
	if totalUnpaid > 0 {
		return errors.New("user has unpaid fines")
	}

	// Cek pinjeman aktifnya kebanyakan ga (max 3 yak)
	activeCount, err := s.borrowRepo.GetActiveBorrowCount(userID)
	if err != nil {
		return err
	}
	if activeCount >= 3 {
		return errors.New("user has reached borrowing limit (3 books)")
	}

	// Cek stok buku masih ada ga
	book, err := s.bookRepo.GetByID(bookID)
	if err != nil {
		return err
	}
	if book.Stock <= 0 {
		return errors.New("book is out of stock")
	}

	// Bikin record pinjeman nih
	borrow := &domain.Borrow{
		UserID:     userID,
		BookID:     bookID,
		BorrowDate: time.Now(),
		DueDate:    time.Now().AddDate(0, 0, 7), // Jatah minjem 7 hari doang
		Status:     domain.StatusBorrowed,
	}

	err = s.borrowRepo.Create(borrow)
	if err != nil {
		return err
	}

	// Kurangin stok buku lerrr
	book.Stock--
	return s.bookRepo.Update(book)
}

func (s *borrowService) ReturnBook(borrowID int64) (float64, error) {
	borrow, err := s.borrowRepo.GetByID(borrowID)
	if err != nil {
		return 0, err
	}

	if borrow.Status == domain.StatusReturned {
		return 0, errors.New("book already returned")
	}

	now := time.Now()
	borrow.ReturnDate = &now
	borrow.Status = domain.StatusReturned

	// Itung denda nya nih
	var fineAmount float64 = 0
	if now.After(borrow.DueDate) {
		daysLate := math.Ceil(now.Sub(borrow.DueDate).Hours() / 24)
		fineAmount = daysLate * 2000 // 2 rebu sehari gpp lah ya buat kas
	}

	err = s.borrowRepo.Update(borrow)
	if err != nil {
		return 0, err
	}

	// Kalo ada denda kita masukin database langsung
	if fineAmount > 0 {
		fine := &domain.Fine{
			BorrowID: borrow.ID,
			UserID:   borrow.UserID,
			Amount:   fineAmount,
			Status:   domain.FineUnpaid,
		}
		err = s.fineRepo.Create(fine)
		if err != nil {
			return fineAmount, err
		}
	}

	// Balikin stok buku nya soalnya udah dibalikin
	book, err := s.bookRepo.GetByID(borrow.BookID)
	if err != nil {
		return fineAmount, err
	}
	book.Stock++
	err = s.bookRepo.Update(book)
	if err != nil {
		return fineAmount, err
	}

	return fineAmount, nil
}

func (s *borrowService) GetUserHistory(userID int64) ([]*domain.Borrow, error) {
	return s.borrowRepo.ListByUser(userID)
}

func (s *borrowService) ListAllBorrows() ([]*domain.Borrow, error) {
	return s.borrowRepo.ListAll()
}

func (s *borrowService) GetUserFines(userID int64) ([]*domain.Fine, float64, error) {
	fines, err := s.fineRepo.ListByUser(userID)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.fineRepo.GetTotalUnpaid(userID)
	return fines, total, err
}
