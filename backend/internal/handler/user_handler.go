package handler

import (
	"encoding/json"
	"net/http"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/middleware"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/service"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/pkg/response"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.Register(req.Name, req.Email, req.Password); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, "user registered successfully", nil)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	token, user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "login successful", map[string]interface{}{
		"token": token,
		"user":  user,
	})
}

func (h *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserContextKey).(int64)
	user, err := h.service.GetProfile(userID)
	if err != nil {
		response.Error(w, http.StatusNotFound, "user not found")
		return
	}
	response.Success(w, http.StatusOK, "profile fetched", user)
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(w, http.StatusOK, "users listed", users)
}
