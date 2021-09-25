package access_token

import (
	"github.com/bookstores/oauth-api/db"
	"github.com/bookstores/oauth-api/domain"
)

func NewAccessTokenRepository() IAccessTokenRepository {
	return &AccessTokenRepository{}
}

type IAccessTokenRepository interface {
	GetTokenById(id string) (*domain.AccessToken, error)
	SetToken(accessToken *domain.AccessToken) error
}

type AccessTokenRepository struct{}

func (r *AccessTokenRepository) GetTokenById(clientId string) (*domain.AccessToken, error) {
	val := &domain.AccessToken{}
	err := db.Get(clientId, val)
	if err != nil {
		return nil, err //errors.NewInternalError(err.Error())
	}

	return val, nil
}

func (r *AccessTokenRepository) SetToken(accessToken *domain.AccessToken) error {
	if err := db.Set(accessToken.ClientId, accessToken, 0); err != nil {
		return err //errors.NewInternalError(err.Error())
	}
	return nil
}
