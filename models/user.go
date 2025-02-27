package models

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Mail      string
	Admin     bool

	// Košík
	Cart     []Book `gorm:"many2many:carts;"`     // TODO skrýt ve JSONu
	Wishlist []Book `gorm:"many2many:wishlists;"` // TODO skrýt ve JSONu
}
