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
	HasImage  bool

	// Žánry
	Genres []Genre `gorm:"many2many:bookGenres;"`
	// Autor
	Author Author
}

// Vrátí novou knihu.
//
//	@param name
//	@param authorId
//	@param summary
//	@param isbn
//	@param price
//	@param published
//	@param hasImage
//	@return *Book
func NewBook(name string, authorId uint, summary string, isbn string, price uint64, published time.Time, hasImage bool) *Book {
	return &Book{
		Name:      name,
		AuthorID:  authorId,
		Summary:   summary,
		Isbn:      isbn,
		Price:     price,
		Published: published,
		HasImage:  hasImage,
	}
}
