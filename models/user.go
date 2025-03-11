package models

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint   `json:"id" validate:"required"`
	FirstName    string `json:"firstName" validate:"required"`
	LastName     string `json:"lastName" validate:"required"`
	Mail         string `json:"mail" validate:"required"`
	Role         Role   `json:"role" validate:"required"`
	Points       int    `json:"points" validate:"required"`
	PasswordHash string `json:"-"`

	// Košík
	Cart []Book `json:"-" gorm:"many2many:carts;"`
	// Seznam přání
	Wishlist []Book `json:"-" gorm:"many2many:wishlists;"`
	// Adresy uživatele
	Addresses []Address `json:"-" gorm:"many2many:user_addresses;"`
}

// Vrátí nového uživatele.
//
//	@param firstName
//	@param lastName
//	@param mail
//	@param role
//	@param password
//	@return *User
//	@return error
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

// Aktualizace atributů uživatele.
//
//	@param uu
//	@return error
func (u *User) Update(uu *UpdateUser) error {
	if uu.FirstName != nil {
		u.FirstName = *uu.FirstName
	}
	if uu.LastName != nil {
		u.LastName = *uu.LastName
	}
	if uu.Mail != nil {
		u.Mail = *uu.Mail
	}
	if uu.Password != nil {
		bytes, err := bcrypt.GenerateFromPassword([]byte(*uu.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.PasswordHash = base64.StdEncoding.EncodeToString(bytes)
	}
	return nil
}
