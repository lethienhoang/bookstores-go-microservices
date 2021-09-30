package middlewares

import (
	"github.com/bookstores/oauth-api/utils/jwt_auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		bearToken := context.Request.Header.Get("Authorization")
		_, err := jwt_auth.DecodeToken(bearToken, false)
		if err != nil {
			context.JSON(http.StatusUnauthorized, err.Error())
			context.Abort()
			return
		}

		context.Next()
	}
}