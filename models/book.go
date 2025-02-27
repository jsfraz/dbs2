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

	// Recenze
	Reviews []Review // TODO skrýt ve JSONu
	// Žánry
	Genres []Genre `gorm:"many2many:bookGenres;"`
}
