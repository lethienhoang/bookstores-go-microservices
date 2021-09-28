package https

import (
	"github.com/bookstores/oauth-api/requests"
	"github.com/bookstores/oauth-api/services"
	"github.com/bookstores/oauth-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUserHandler interface {
	Login(c *gin.Context)
}

type UserHandler struct {
	service services.IUserService
}

func NewAccessTokenHandler(service services.IUserService) IUserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u UserHandler) Login(c *gin.Context) {
	var req requests.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	login, err := u.service.LoginUser(&req)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, login)
}
