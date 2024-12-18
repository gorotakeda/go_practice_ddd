package service

import (
	"testing"

	"github.com/goro/go_practice_ddd/domain"
	"github.com/goro/go_practice_ddd/mock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUsers(t *testing.T) {
	mockRepo := &mock.MockUserRepository{
		Users: []domain.User{
			{ID: 1, Name: "ユーザー1", Email: "user1@example.com"},
			{ID: 2, Name: "ユーザー2", Email: "user2@example.com"},
		},
	}

	service := NewUserService(mockRepo)

	users, err := service.GetUsers()
	assert.NoError(t, err)
	assert.Len(t, users, len(mockRepo.Users))
}

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := &mock.MockUserRepository{}
	service := NewUserService(mockRepo)

	user := &domain.User{
		Name:  "テストユーザー",
		Email: "test@example.com",
	}

	err := service.CreateUser(user)
	assert.NoError(t, err)
	assert.Len(t, mockRepo.Users, 1)
} 