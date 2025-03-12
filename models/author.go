package models

import (
	"time"
)

type Author struct {
	ID        uint      `json:"id" validate:"required"`
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Birth     time.Time `json:"birth" validate:"required"`
}

// Vrátí nového autora.
//
//	@param firstName
//	@param lastName
//	@param birth
//	@return *Author
func NewAuthor(firstName string, lastName string, birth time.Time) *Author {
	return &Author{
		FirstName: firstName,
		LastName:  lastName,
		Birth:     birth,
	}
}
