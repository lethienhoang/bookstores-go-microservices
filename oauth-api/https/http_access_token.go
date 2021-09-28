package https

import (
	"net/http"
	"strings"

	"github.com/bookstores/oauth-api/domain"
	"github.com/bookstores/oauth-api/services"
	"github.com/bookstores/oauth-api/utils/errors"
	"github.com/gin-gonic/gin"
)

type IAccessTokenHandler interface {
	GetTokenById(c *gin.Context)
	CreateAccessToken(c *gin.Context)
}

type AccessTokenHandler struct {
	service services.IAccessTokenService
}

func NewAccessToken(service services.IAccessTokenService) IAccessTokenHandler {
	return &AccessTokenHandler{
		service: service,
	}
}

func (h *AccessTokenHandler) GetTokenById(c *gin.Context) {
	tokenId := strings.TrimSpace(c.Param("token_id"))
	accessToken, err := h.service.GetTokenByClientId(tokenId)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (h *AccessTokenHandler) CreateAccessToken(c *gin.Context) {
	var at domain.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		parsErr := errors.NewBadRequestError("invalid json body")
		c.JSON(parsErr.Code, parsErr)
		return
	}

	if err := h.service.CreateNewToken(&at); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, at)
}
