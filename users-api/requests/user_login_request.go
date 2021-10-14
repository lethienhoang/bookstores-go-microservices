package requests

import (
	"strings"

	"github.com/lethienhoang/bookstores-utils-go/errors"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginRequest) Validate() *errors.RestError {
	l.Password = strings.TrimSpace(l.Password)
	if l.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	l.Email = strings.TrimSpace(strings.ToLower(l.Email))
	if l.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}
