package domain

import (
	"time"

	"github.com/bookstores/oauth-api/utils/errors"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    string `json:"client_id"`
	Expires     int64  `json:"expires"`
}

const (
	expiresIn = 12
)

func GetAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expiresIn * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expiresIn := time.Unix(at.Expires, 0)

	return expiresIn.Before(now)
}

func (at AccessToken) IsExpiresSoon() bool {
	now := time.Now().UTC().Unix()
	if now < at.Expires && (at.Expires-now) <= 10800 {
		return true
	}

	return false
}

func (at AccessToken) Validate() *errors.RestError {
	if at.UserId <= 0 {
		return errors.NewBadRequestError("user id must be greater than 0")
	}

	if len(at.AccessToken) == 0 {
		return errors.NewBadRequestError("token must be at least one character")
	}

	if at.ClientId == "" {
		return errors.NewBadRequestError("client id must be set")
	}

	if at.Expires == 0 {
		return errors.NewBadRequestError("invalid expires time")
	}

	return nil
}
