package service

import (
	"github.com/goro/go_practice_ddd/domain"
	"github.com/goro/go_practice_ddd/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers() ([]domain.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.repo.Create(user)
}
