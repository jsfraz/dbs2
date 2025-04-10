package models

import "time"

type Discount struct {
	ID         uint      `json:"id" validate:"required"`
	UserId     uint      `json:"userId " validate:"required"`
	Code       string    `json:"code" validate:"required"`
	ValidUntil time.Time `json:"validUntil" validate:"required"`
	Used       bool      `json:"used" validate:"required"`
	Price      uint64    `json:"price" validate:"required"`

	// Uživatel
	User User `json:"-"`
}
