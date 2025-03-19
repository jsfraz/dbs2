package models

import "time"

type Discount struct {
	ID         uint
	UserId     uint
	Code       string
	ValidUntil time.Time
	Used       bool
	Price      uint64

	// Uživatel
	User User
}
