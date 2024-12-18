package repository

import (
	"github.com/goro/go_practice_ddd/domain"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetAll() ([]domain.User, error)
	Create(user *domain.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}
