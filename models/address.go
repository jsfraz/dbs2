package models

type Address struct {
	ID       uint
	UserID   uint
	City     string
	Street   string
	PostCode int

	// Uživatel
	User User // TODO skrýt ve JSONu
}
