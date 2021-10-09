package users

import (
	"encoding/json"

	"github.com/bookstores-go-microservices/oauth-api/dtos"
	"github.com/bookstores-go-microservices/oauth-api/requests"
	"github.com/bookstores-go-microservices/oauth-api/utils/crypto"
	"github.com/bookstores-go-microservices/oauth-api/utils/errors"
	"github.com/go-resty/resty/v2"
	_ "github.com/go-resty/resty/v2"
)

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

type IUserRepository interface {
	LoginUser(request *requests.LoginRequest) (*dtos.UserResponseDto, *errors.RestError)
}

type UserRepository struct{}

type HttpError struct {
	ID, Message string
}

type HttpSuccess struct {
	ID, Message string
}

func (s *UserRepository) LoginUser(request *requests.LoginRequest) (*dtos.UserResponseDto, *errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	request.Password = crypto.GetMd5Hash(request.Password)
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		Post("http://localhost:8080/api/v1/users/login")

	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	if resp.IsSuccess() {
		userDto := dtos.UserResponseDto{}
		if err := json.Unmarshal(resp.Body(), &userDto); err != nil {
			return nil, errors.NewBadRequestError("error in convert byte to object")
		}

		return &userDto, nil
	} else {
		httError := resp.Error().(*HttpError)
		return nil, errors.NewInternalError(httError.Message)
	}
}
