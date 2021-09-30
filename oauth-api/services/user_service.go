package services

import (
	"github.com/bookstores/oauth-api/dtos"
	"github.com/bookstores/oauth-api/repository/access_token"
	"github.com/bookstores/oauth-api/repository/users"
	"github.com/bookstores/oauth-api/requests"
	"github.com/bookstores/oauth-api/utils/errors"
)

type IUserService interface {
	LoginUser(request *requests.LoginRequest) (*dtos.TokenDto, *errors.RestError)
	//RefeshToken()
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
		return  nil, err
	}

	atService := NewAccessTokenService(access_token.NewAccessTokenRepository())
	token, err := atService.CreateNewToken(user.UserInfo.Id)
	if err != nil {
		return  nil, err
	}

	return &dtos.TokenDto {
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
