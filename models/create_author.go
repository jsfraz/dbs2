package models

type CreateAuthor struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Birth     string `json:"birth" validate:"required"`
}
