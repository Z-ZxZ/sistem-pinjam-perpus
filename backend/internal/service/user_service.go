package service

import (
	"errors"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/domain"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/pkg/utils"
)

type UserService interface {
	Register(name, email, password string) error
	Login(email, password string) (string, *domain.User, error)
	GetProfile(id int64) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
}

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(name, email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     domain.RoleMember,
		Status:   "active",
	}

	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (string, *domain.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *userService) GetProfile(id int64) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) ListUsers() ([]*domain.User, error) {
	return s.repo.List()
}
