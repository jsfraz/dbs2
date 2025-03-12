package models

type CreateGenre struct {
	Name string `json:"name" validate:"required"`
}
