package app

import (
	"github.com/bookstores/oauth-api/db"
	"github.com/bookstores/oauth-api/https"
	"github.com/bookstores/oauth-api/repository/access_token"
	"github.com/bookstores/oauth-api/services"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	db.NewRedisClient()

	atService := services.NewAccessTokenService(access_token.NewAccessTokenRepository())
	atHandler := https.NewHandler(atService)

	router.GET("/oauth/access_token/:token_id", atHandler.GetTokenById)
	router.POST("/oauth/access_token", atHandler.CreateAccessToken)
	router.Run(":8081")
}
