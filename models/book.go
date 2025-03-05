package models

import "gorm.io/datatypes"

type Book struct {
	ID       uint
	Name     string
	AuthorID uint
	Summary  string
	Isbn     string
	Price    uint64
	datatypes.Date

	// Žánry
	Genres []Genre `gorm:"many2many:bookGenres;"`
	// Autor
	Author Author
}
