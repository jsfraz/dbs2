package models

import "time"

type Order struct {
	ID         uint      `json:"id" validate:"required"`
	UserID     uint      `json:"userId" validate:"required"`
	TotalPrice uint64    `json:"totalPrice" validate:"required"`
	Status     string    `json:"status" validate:"required"`
	AddressID  uint      `json:"addressId" validate:"required"`
	DiscountID *uint     `json:"-"`
	CreatedAt  time.Time `json:"createdAt" validate:"required"`

	// Uživatel
	User User `json:"-"`
	// Knížky v objednávce
	OrderedBooks []Book `json:"orderedBooks" validate:"required" gorm:"many2many:user_order_books;"`
	// (Ne)Použitá sleva
	Discount *Discount `json:"discount"`
	// Adresa
	Address Address `json:"address" validate:"required"`
}
