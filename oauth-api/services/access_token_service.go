package services

import (
	"github.com/bookstores/oauth-api/domain"
	"github.com/bookstores/oauth-api/repository/access_token"
	"github.com/bookstores/oauth-api/utils/errors"
)

type IAccessTokenService interface {
	GetTokenByClientId(id string) (*domain.AccessToken, *errors.RestError)
	CreateNewToken(accessToken *domain.AccessToken) *errors.RestError
}

type AccessTokenService struct {
	repository access_token.IAccessTokenRepository
}

func NewAccessTokenService(repo access_token.IAccessTokenRepository) IAccessTokenService {
	return &AccessTokenService{
		repository: repo,
	}
}

func (s *AccessTokenService) GetTokenByClientId(id string) (*domain.AccessToken, *errors.RestError) {
	if len(id) == 0 {
		return nil, errors.NewBadRequestError("no access token id provided")
	}

	token, err := s.repository.GetTokenById(id)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	if err := token.Validate(); err != nil {
		return nil, err
	}

	// if now < token.Expires && (token.Expires-now) <= 10800 {
	// 	// update token when token is expried soon
	// 	accessToken := domain.AccessToken{}
	// 	accessToken.ClientId = id
	// 	accessToken.Expires = time.Now().Add(time.Hour * 24).UTC().Unix()
	// 	accessToken.UserId = token.UserId

	// 	if err := s.repository.SetToken(&accessToken); err != nil {
	// 		return nil, errors.NewInternalError(err.Error())
	// 	}
	// }

	return token, nil
}

func (s *AccessTokenService) CreateNewToken(accessToken *domain.AccessToken) *errors.RestError {
	if err := accessToken.Validate(); err != nil {
		return err
	}

	if newRrr := s.repository.SetToken(accessToken); newRrr != nil {
		return errors.NewInternalError(newRrr.Error())
	}

	return nil
}
