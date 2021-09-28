package services

import (
	"github.com/bookstores/oauth-api/dtos"
	"github.com/bookstores/oauth-api/repository/users"
	"github.com/bookstores/oauth-api/requests"
	"github.com/bookstores/oauth-api/utils/errors"
)

type IUserService interface {
	LoginUser(request *requests.LoginRequest) (*dtos.UserDto, *errors.RestError)
}

type UserService struct {
	repository users.IUserRepository
}

func NewUserService(repo users.IUserRepository) IUserService {
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) LoginUser(request *requests.LoginRequest) (*dtos.UserDto, *errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//atService := NewAccessTokenService(access_token.NewAccessTokenRepository())
	//atService.CreateNewToken()

	return s.repository.LoginUser(request)
}
