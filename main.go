package main

import (
	"log"
	"github.com/goro/go_practice_ddd/handler"
	"github.com/goro/go_practice_ddd/infrastructure"
	"github.com/goro/go_practice_ddd/repository"
	"github.com/goro/go_practice_ddd/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 設定とDB接続
	db := infrastructure.InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	// DI (依存関係の注入)
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Ginルーターの設定
	router := gin.Default()

	// エンドポイント定義
	router.GET("/users", userHandler.GetUsers)
	router.POST("/users", userHandler.CreateUser)

	// サーバー起動
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
