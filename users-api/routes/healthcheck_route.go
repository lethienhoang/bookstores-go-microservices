package routes

import (
	"github.com/bookstores-go-microservices/users-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupHealthCheckRouter(router *gin.RouterGroup) {
	healthcheck := router.Group("healthcheck")
	{
		healthcheck.GET("service", controllers.PingSerivce)
		healthcheck.GET("database", controllers.PingDatabases)
	}
}
