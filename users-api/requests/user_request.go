package requests

import (
	"strings"

	"github.com/bookstores/users-api/untils/errors"
)

type CreateOrUpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUserPasswordRequest struct {
	Password string `json:"password"`
}

func (request *CreateOrUpdateUserRequest) Validate() *errors.RestError {
	request.Email = strings.TrimSpace(strings.ToLower(request.Email))
	if request.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	request.Password = strings.TrimSpace(request.Password)
	if request.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}

func (request *UpdateUserPasswordRequest) Validate() *errors.RestError {
	request.Password = strings.TrimSpace(request.Password)
	if request.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
