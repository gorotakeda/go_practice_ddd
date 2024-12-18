package repository

import (
	"testing"

	"github.com/goro/go_practice_ddd/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestUserRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	user := &domain.User{
		Name:  "テストユーザー",
		Email: "test@example.com",
	}

	err := repo.Create(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
}

func TestUserRepository_GetAll(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	// テストデータの作成
	testUsers := []domain.User{
		{Name: "ユーザー1", Email: "user1@example.com"},
		{Name: "ユーザー2", Email: "user2@example.com"},
	}

	for _, u := range testUsers {
		err := repo.Create(&u)
		assert.NoError(t, err)
	}

	// テスト実行
	users, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, users, len(testUsers))
} 