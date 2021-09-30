package services

import (
	"github.com/bookstores/oauth-api/dtos"
	"github.com/bookstores/oauth-api/repository/access_token"
	"github.com/bookstores/oauth-api/utils/errors"
	"github.com/bookstores/oauth-api/utils/jwt_auth"
)

type IAccessTokenService interface {
	GetTokenByClientId(id string) (int64, *errors.RestError)
	CreateNewToken(userId int64) (*dtos.TokenDetailsDto, *errors.RestError)
}

type AccessTokenService struct {
	repository access_token.IAccessTokenRepository
}

func NewAccessTokenService(repo access_token.IAccessTokenRepository) IAccessTokenService {
	return &AccessTokenService{
		repository: repo,
	}
}

func (s *AccessTokenService) GetTokenByClientId(id string) (int64, *errors.RestError) {
	if len(id) == 0 {
		return 0, errors.NewBadRequestError("no access token id provided")
	}

	userId, err := s.repository.GetTokenById(id)
	if err != nil {
		return 0, errors.NewInternalError(err.Error())
	}

	return userId, nil
}

func (s *AccessTokenService) CreateNewToken(userId int64) (*dtos.TokenDetailsDto, *errors.RestError) {
	tokenDetail := new(dtos.TokenDetailsDto)
	if err := jwt_auth.CreateToken(userId, tokenDetail); err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	if err := jwt_auth.CreateRefreshToken(userId, tokenDetail); err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	var err error
	err = s.repository.SetToken(tokenDetail.AccessUuid, userId, tokenDetail.AtExpires)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	err = s.repository.SetToken(tokenDetail.RefreshUuid, userId, tokenDetail.RtExpires)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	return tokenDetail, nil
}
