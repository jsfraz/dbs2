package models

type LoginResponse struct {
	AccessToken string `json:"accessToken" validate:"required"`
}

func NewLoginResponse(accessToken string) *LoginResponse {
	l := new(LoginResponse)
	l.AccessToken = accessToken
	return l
}
