package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goro/go_practice_ddd/domain"
	"github.com/goro/go_practice_ddd/mock"
	"github.com/goro/go_practice_ddd/service"
	"github.com/stretchr/testify/assert"
)

func setupRouter(mockRepo *mock.MockUserRepository) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	userService := service.NewUserService(mockRepo)
	userHandler := NewUserHandler(userService)

	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)

	return r
}

func TestUserHandler_GetUsers(t *testing.T) {
	mockRepo := &mock.MockUserRepository{
		Users: []domain.User{
			{ID: 1, Name: "ユーザー1", Email: "user1@example.com"},
		},
	}

	router := setupRouter(mockRepo)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []domain.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, len(mockRepo.Users))
}

func TestUserHandler_CreateUser(t *testing.T) {
	mockRepo := &mock.MockUserRepository{}
	router := setupRouter(mockRepo)

	user := domain.User{
		Name:  "テストユーザー",
		Email: "test@example.com",
	}

	body, _ := json.Marshal(user)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response domain.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, response.Name)
	assert.Equal(t, user.Email, response.Email)
} 