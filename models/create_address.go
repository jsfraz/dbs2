package models

type CreateAddress struct {
	City     string `json:"city" validate:"required"`
	Street   string `json:"street" validate:"required"`
	PostCode int    `json:"postCode" validate:"required"`
}
