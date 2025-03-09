package models

type Login struct {
	Mail     string `json:"mail" validate:"required"`
	Password string `json:"password" validate:"required"`
}
