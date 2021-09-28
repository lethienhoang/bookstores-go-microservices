package users

import (
	"encoding/json"
	"github.com/bookstores/oauth-api/dtos"
	"github.com/bookstores/oauth-api/requests"
	"github.com/bookstores/oauth-api/utils/crypto"
	"github.com/bookstores/oauth-api/utils/errors"
	"github.com/go-resty/resty/v2"
	_ "github.com/go-resty/resty/v2"
)

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

type IUserRepository interface {
	LoginUser(request *requests.LoginRequest) (*dtos.UserDto, *errors.RestError)
}

type UserRepository struct{}

type HttpError struct {
	ID, Message string
}

type HttpSuccess struct {
	ID, Message string
}

func (s *UserRepository) LoginUser(request *requests.LoginRequest) (*dtos.UserDto, *errors.RestError) {
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
		userDto := dtos.UserDto{}
		if err := json.Unmarshal(resp.Body(), &userDto); err != nil {
			return nil, errors.NewBadRequestError("error in convert byte to object")
		}

		return &userDto, nil
	} else {
		httError := resp.Error().(*HttpError)
		return nil, errors.NewInternalError(httError.Message)
	}

	return nil, nil
}
