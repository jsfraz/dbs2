package models

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint
	FirstName    string
	LastName     string
	Mail         string
	Role         Role
	Points       int
	PasswordHash string

	// Košík
	Cart []Book `gorm:"many2many:carts;"` // TODO skrýt ve JSONu
	// Seznam přání
	Wishlist []Book `gorm:"many2many:wishlists;"` // TODO skrýt ve JSONu
	// Adresy uživatele
	Addresses []Address `gorm:"many2many:user_addresses;"` // TODO skrýt ve JSONu
}

func NewUser(firstName, lastName, mail string, role Role, password string) (*User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:    firstName,
		LastName:     lastName,
		Mail:         mail,
		Role:         role,
		Points:       0,
		PasswordHash: base64.StdEncoding.EncodeToString(bytes),
		Cart:         []Book{},
		Wishlist:     []Book{},
		Addresses:    []Address{},
	}, nil
}
