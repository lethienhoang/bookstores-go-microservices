package app

import (
	"github.com/bookstores/oauth-api/db"
	"github.com/bookstores/oauth-api/https"
	"github.com/bookstores/oauth-api/repository/users"
	"github.com/bookstores/oauth-api/services"
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
	router.Run(":8081")
}
