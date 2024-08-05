package service

import (
	"github.com/usmonzodasomon/shortener/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryI interface {
	CreateUser(user models.User) (int, error)
}

type UserService struct {
	repo UserRepositoryI
}

func NewUserService(repo UserRepositoryI) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(user models.User) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	id, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
