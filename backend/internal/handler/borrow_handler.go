package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/middleware"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/service"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/pkg/response"
)

type BorrowHandler struct {
	service service.BorrowService
}

func NewBorrowHandler(service service.BorrowService) *BorrowHandler {
	return &BorrowHandler{service: service}
}

func (h *BorrowHandler) Borrow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		BookID int64 `json:"book_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	userID := r.Context().Value(middleware.UserContextKey).(int64)
	if err := h.service.BorrowBook(userID, req.BookID); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, "book borrowed successfully", nil)
}

func (h *BorrowHandler) Return(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		BorrowID int64 `json:"borrow_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	fine, err := h.service.ReturnBook(req.BorrowID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "book returned successfully", map[string]interface{}{
		"fine_amount": fine,
	})
}

func (h *BorrowHandler) History(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserContextKey).(int64)
	history, err := h.service.GetUserHistory(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(w, http.StatusOK, "history fetched", history)
}

func (h *BorrowHandler) AllBorrows(w http.ResponseWriter, r *http.Request) {
	borrows, err := h.service.ListAllBorrows()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(w, http.StatusOK, "all borrows fetched", borrows)
}

func (h *BorrowHandler) Fines(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserContextKey).(int64)
	fines, total, err := h.service.GetUserFines(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(w, http.StatusOK, "fines fetched", map[string]interface{}{
		"fines": fines,
		"total": total,
	})
}
