package models

type UpdateUser struct {
	Id        uint    `json:"id" validate:"required"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Mail      *string `json:"mail" validate:"omitempty,email"`
	Password  *string `json:"password"` // TODO validace délky a tak
}
