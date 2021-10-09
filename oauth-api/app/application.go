package app

import (
	"github.com/bookstores-go-microservices/oauth-api/db"
	"github.com/bookstores-go-microservices/oauth-api/https"
	"github.com/bookstores-go-microservices/oauth-api/repository/users"
	"github.com/bookstores-go-microservices/oauth-api/services"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	db.NewRedisClient()

	userService := services.NewUserService(users.NewUserRepository())
	userHandler := https.NewAccessTokenHandler(userService)

	router.POST("/oauth/login", userHandler.Login)
	router.GET("/oauth/verify-token", userHandler.VerifyToken)
	router.POST("/oauth/refresh-token", userHandler.RefeshToken)
	router.POST("/oauth/logout", userHandler.LogOut)

	err := router.Run(":8081")
	if err != nil {
		return
	}
}
