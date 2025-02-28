package models

type User struct {
	ID           uint
	FirstName    string
	LastName     string
	Mail         string
	Admin        bool
	Points       int
	PasswordHash string

	// Košík
	Cart []Book `gorm:"many2many:carts;"` // TODO skrýt ve JSONu
	// Seznam přání
	Wishlist []Book `gorm:"many2many:wishlists;"` // TODO skrýt ve JSONu
	// Adresy uživatele
	Addresses []Address `gorm:"many2many:user_addresses;"` // TODO skrýt ve JSONu
}
