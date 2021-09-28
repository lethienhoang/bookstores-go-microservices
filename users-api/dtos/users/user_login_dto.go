package dtos

type LoginDto struct {
	Email string `json:"email"`
	AccessToken string `json:"access_token"`
	IsLogin bool `json:"is_login"`
}
