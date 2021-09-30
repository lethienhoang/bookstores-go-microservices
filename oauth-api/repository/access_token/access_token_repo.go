package access_token

import (
	"github.com/bookstores/oauth-api/db"
	"time"
)

func NewAccessTokenRepository() IAccessTokenRepository {
	return &AccessTokenRepository{}
}

type IAccessTokenRepository interface {
	GetTokenById(clientId string) (int64, error)
	SetToken(clientId string, userId int64, expToken int64) error
}

type AccessTokenRepository struct{}

func (r *AccessTokenRepository) GetTokenById(clientId string) (int64, error) {
	var val int64
	err := db.Get(clientId, val)
	if err != nil {
		return 0, err //errors.NewInternalError(err.Error())
	}

	return val, nil
}

func (r *AccessTokenRepository) SetToken(clientId string, userId int64, expToken int64) error {
	at := time.Unix(expToken, 0)
	now := time.Now()
	if err := db.Set(clientId, userId, at.Sub(now)); err != nil {
		return err //errors.NewInternalError(err.Error())
	}
	return nil
}
