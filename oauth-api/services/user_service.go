package services

import (
	"github.com/bookstores-go-microservices/oauth-api/dtos"
	"github.com/bookstores-go-microservices/oauth-api/repository/access_token"
	"github.com/bookstores-go-microservices/oauth-api/repository/users"
	"github.com/bookstores-go-microservices/oauth-api/requests"
	"github.com/bookstores-go-microservices/oauth-api/utils/errors"
)

type IUserService interface {
	LoginUser(request *requests.LoginRequest) (*dtos.TokenDto, *errors.RestError)
	RefeshToken(userId int64) (*dtos.TokenDto, *errors.RestError)
	DelToken(tokenId string) *errors.RestError
}

type UserService struct {
	repository users.IUserRepository
}

func NewUserService(repo users.IUserRepository) IUserService {
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) LoginUser(request *requests.LoginRequest) (*dtos.TokenDto, *errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := s.repository.LoginUser(request)
	if err != nil {
		return nil, err
	}

	atService := NewAccessTokenService(access_token.NewAccessTokenRepository())
	token, err := atService.CreateNewToken(user.UserInfo.Id)
	if err != nil {
		return nil, err
	}

	return &dtos.TokenDto{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (s *UserService) RefeshToken(userId int64) (*dtos.TokenDto, *errors.RestError) {
	atService := NewAccessTokenService(access_token.NewAccessTokenRepository())
	token, err := atService.CreateNewToken(userId)
	if err != nil {
		return nil, err
	}

	return &dtos.TokenDto{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (s *UserService) DelToken(tokenId string) *errors.RestError {
	atService := NewAccessTokenService(access_token.NewAccessTokenRepository())
	return atService.DeleteToken(tokenId)
}
