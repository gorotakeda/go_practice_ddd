package mock

import (
	"github.com/goro/go_practice_ddd/domain"
	"github.com/goro/go_practice_ddd/repository"
)

var _ repository.UserRepositoryInterface = (*MockUserRepository)(nil)

type MockUserRepository struct {
	Users []domain.User
}

func (m *MockUserRepository) GetAll() ([]domain.User, error) {
	return m.Users, nil
}

func (m *MockUserRepository) Create(user *domain.User) error {
	user.ID = uint(len(m.Users) + 1)
	m.Users = append(m.Users, *user)
	return nil
} 