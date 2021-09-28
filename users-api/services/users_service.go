package services

import (
	"github.com/bookstores/users-api/domains/users"
	dtos "github.com/bookstores/users-api/dtos/users"
	"github.com/bookstores/users-api/requests"
	"github.com/bookstores/users-api/untils/crypto"
	"github.com/bookstores/users-api/untils/errors"
	"github.com/jinzhu/copier"
)

type userService struct {
}

type userServiceInterface interface {
	Create(request *requests.CreateOrUpdateUserRequest) (*dtos.UserDto, *errors.RestError)
	Get(userId int64) (*dtos.UserDto, *errors.RestError)
	Update(request *requests.CreateOrUpdateUserRequest, userId int64) (*dtos.UserDto, *errors.RestError)
	Delete(userId int64) *errors.RestError
	FindByStatus(status string) ([]dtos.UserDto, *errors.RestError)
	LoginUser(request *requests.LoginRequest) (*dtos.LoginDto, *errors.RestError)
}

var (
	UserService userServiceInterface = &userService{}
)

func (s *userService) Create(request *requests.CreateOrUpdateUserRequest) (*dtos.UserDto, *errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user := new(users.User)
	request.Password= crypto.GetMd5Hash(request.Password)

	if err := copier.Copy(&user, &request); err != nil {
		return nil, errors.NewBadRequestError("can't copy user: " + err.Error())
	}

	if err := user.Create(); err != nil {
		return nil, err
	}

	res := dtos.UserDto{}
	if err := copier.Copy(&res, &user); err != nil {
		return nil, errors.NewBadRequestError("can't copy user: " + err.Error())
	}

	return &res, nil
}

func (s *userService) Get(userId int64) (*dtos.UserDto, *errors.RestError) {
	user := users.User{
		Id: userId,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	res := dtos.UserDto{}
	if err := copier.Copy(&res, &user); err != nil {
		return nil, errors.NewBadRequestError("can't copy user: " + err.Error())
	}

	return &res, nil
}

func (s *userService) Update(request *requests.CreateOrUpdateUserRequest, userId int64) (*dtos.UserDto, *errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if userId == 0 {
		return nil, errors.NewBadRequestError("invalid userId")
	}

	user := users.User{
		Id: userId,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	if err := copier.Copy(&user, &request); err != nil {
		return nil, errors.NewBadRequestError("can't copy user: " + err.Error())
	}

	if err := user.Update(); err != nil {
		return nil, err
	}

	res := dtos.UserDto{}
	if err := copier.Copy(&res, &user); err != nil {
		return nil, errors.NewBadRequestError("can't copy user: " + err.Error())
	}

	return &res, nil
}

func (s *userService) Delete(userId int64) *errors.RestError {
	user := users.User{
		Id: userId,
	}

	return user.Delete()
}

func (s *userService) FindByStatus(status string) ([]dtos.UserDto, *errors.RestError) {
	user := &users.User{}
	users, err := user.FindByStatus(status)
	if err != nil {
		return nil, err
	}

	res := make([]dtos.UserDto, 0)
	if err := copier.Copy(&res, &users); err != nil {
		return nil, errors.NewBadRequestError("can't copy user: " + err.Error())
	}

	return res, nil
}

func (s *userService) LoginUser(request *requests.LoginRequest) (*dtos.LoginDto, *errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user := users.User{
		Email: request.Email,
		//Password: crypto.GetMd5Hash(request.Password),
		Password: request.Password,
	}

	if err := user.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	userDtos := dtos.UserDto{}
	if err := copier.Copy(&userDtos, &user); err != nil {
		return nil, errors.NewBadRequestError("can't copy user: " + err.Error())
	}

	res := dtos.LoginDto{
		UserInfo: userDtos,
		Email: user.Email,
		IsLogin: true,
	}

	return &res, nil
}
