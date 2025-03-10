package models

type Register struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Mail      string `json:"mail" validate:"email,required"`
	Password  string `json:"password" validate:"required"` // TODO validace délky a tak
}
