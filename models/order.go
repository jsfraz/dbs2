package models

type Order struct {
	ID         uint
	UserID     uint
	TotalPrice uint64
	Status     string
	AddressID  uint
	DiscountID *uint

	// Uživatel
	User User // TODO skrýt ve JSONu
	// Knížky v objednávce
	OrderedBooks Book `gorm:"many2many:user_order_books;"`
	// (Ne)Použitá sleva
	Discount *Discount
	// Adresa
	Address Address
}
