package dtos

type UserDto struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type UserResponseDto struct {
	UserInfo UserDto `json:"user_info"`
	Email    string  `json:"email"`
	IsLogin  bool    `json:"is_login"`
}
