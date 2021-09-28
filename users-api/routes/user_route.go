package routes

import (
	"github.com/bookstores/users-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRouter(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET(":user_id", controllers.GetUser)
		users.POST("", controllers.CreateUser)
		users.PUT(":user_id", controllers.UpdateUser)
		users.GET("status", controllers.FindByStatus)
		users.DELETE(":user_id", controllers.DeleteUser)
		users.POST("login", controllers.Login)
	}

}
