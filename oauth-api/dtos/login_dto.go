package dtos

type LoginDto struct {
	Email       string `json:"email"`
	Password    bool   `json:"is_login"`
	AccessToken string `json:"access_token"`
}
