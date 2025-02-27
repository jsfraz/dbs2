package models

type Author struct {
	ID       uint
	Fullname string

	// Knížky autora
	Books []Book // TODO skrýt ve JSONu
}
