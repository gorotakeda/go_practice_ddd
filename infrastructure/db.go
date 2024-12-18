package infrastructure

import (
	"github.com/goro/go_practice_ddd/config"
	"github.com/goro/go_practice_ddd/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitDB() *gorm.DB {
	var db *gorm.DB
	var err error
	maxRetries := 30
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(config.GetDBURL()), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(time.Second)
	}
	if err != nil {
		log.Fatal("Failed to connect to database after retries:", err)
	}

	// Auto Migrate
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}
