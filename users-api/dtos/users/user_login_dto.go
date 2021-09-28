package dtos

type LoginDto struct {
	UserInfo UserDto `json:"user_info"`
	Email string `json:"email"`
	IsLogin bool `json:"is_login"`
}
