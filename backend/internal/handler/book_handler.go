package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/domain"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/service"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/pkg/response"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) Handle(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	
	switch r.Method {
	case http.MethodGet:
		if len(pathParts) > 1 {
			h.GetByID(w, r, pathParts[1])
		} else {
			h.List(w, r)
		}
	case http.MethodPost:
		h.Create(w, r)
	case http.MethodPut:
		if len(pathParts) > 1 {
			h.Update(w, r, pathParts[1])
		} else {
			response.Error(w, http.StatusBadRequest, "book ID required")
		}
	case http.MethodDelete:
		if len(pathParts) > 1 {
			h.Delete(w, r, pathParts[1])
		} else {
			response.Error(w, http.StatusBadRequest, "book ID required")
		}
	default:
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if err := h.service.AddBook(&book); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(w, http.StatusCreated, "book added", book)
}

func (h *BookHandler) GetByID(w http.ResponseWriter, r *http.Request, idStr string) {
	id, _ := strconv.ParseInt(idStr, 10, 64)
	book, err := h.service.GetBook(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "book not found")
		return
	}
	response.Success(w, http.StatusOK, "book fetched", book)
}

func (h *BookHandler) List(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	search := query.Get("search")
	category := query.Get("category")
	page, _ := strconv.Atoi(query.Get("page"))
	if page <= 0 { page = 1 }
	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit <= 0 { limit = 10 }

	books, total, err := h.service.ListBooks(search, category, page, limit)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "books listed", map[string]interface{}{
		"books": books,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request, idStr string) {
	id, _ := strconv.ParseInt(idStr, 10, 64)
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}
	book.ID = id
	if err := h.service.UpdateBook(&book); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(w, http.StatusOK, "book updated", book)
}

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request, idStr string) {
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if err := h.service.DeleteBook(id); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(w, http.StatusOK, "book deleted", nil)
}
