package https

import (
	"github.com/bookstores/oauth-api/repository/access_token"
	"github.com/bookstores/oauth-api/requests"
	"github.com/bookstores/oauth-api/services"
	"github.com/bookstores/oauth-api/utils/errors"
	"github.com/bookstores/oauth-api/utils/jwt_auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUserHandler interface {
	Login(c *gin.Context)
	LogOut(c *gin.Context)
	VerifyToken(c *gin.Context)
	RefeshToken(c *gin.Context)
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

func (u UserHandler) LogOut(c *gin.Context)  {
	bearToken := c.Request.Header.Get("Authorization")
	tokenDecoded, err := jwt_auth.DecodeToken(bearToken, false)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	restErr := u.service.DelToken(tokenDecoded.AccessTokenId)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	restErr = u.service.DelToken(tokenDecoded.RefeshTokenId)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(200, "user has log out")
}

func (u UserHandler) VerifyToken(c *gin.Context) {
	bearToken := c.Request.Header.Get("Authorization")
	tokenDecoded, err := jwt_auth.DecodeToken(bearToken, false)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	atService := services.NewAccessTokenService(access_token.NewAccessTokenRepository())
	userId, errRest := atService.GetTokenByClientId(tokenDecoded.AccessTokenId)
	if errRest != nil {
		c.JSON(http.StatusUnauthorized, errRest.Message)
		return
	}
	if tokenDecoded.UserId != userId {
		c.JSON(http.StatusUnauthorized, "user is not authorized")
		return
	}

	c.JSON(http.StatusOK, "Ok")
}

func (u *UserHandler) RefeshToken(c *gin.Context)  {
	bearToken := map[string]string{}
	if err := c.ShouldBindJSON(&bearToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	tokenDecoded, err := jwt_auth.DecodeToken(bearToken["refresh_token"], false)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	restErr := u.service.DelToken(tokenDecoded.AccessTokenId)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	restErr = u.service.DelToken(tokenDecoded.RefeshTokenId)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	reLogin, restErr := u.service.RefeshToken(tokenDecoded.UserId)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, reLogin)
}
