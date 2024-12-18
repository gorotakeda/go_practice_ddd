package main

import (
	"github.com/goro/go_practice_ddd/config"
	"github.com/goro/go_practice_ddd/handler"
	"github.com/goro/go_practice_ddd/infrastructure"
	"github.com/goro/go_practice_ddd/repository"
	"github.com/goro/go_practice_ddd/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 設定とDB接続
	db := infrastructure.InitDB()
	defer db.Close()

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
	router.Run(":8080")
}
