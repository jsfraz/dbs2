package models

import "time"

type Book struct {
	ID        uint
	Name      string
	AuthorID  uint
	Summary   string
	Isbn      string
	Price     uint64
	Published time.Time

	// Žánry
	Genres []Genre `gorm:"many2many:bookGenres;"`
	// Autor
	Author Author
}
