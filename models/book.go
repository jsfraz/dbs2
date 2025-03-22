package models

import "time"

type Book struct {
	ID        uint      `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	AuthorID  uint      `json:"-"`
	Summary   string    `json:"summary" validate:"required"`
	Isbn      string    `json:"isbn" validate:"required"`
	Price     uint64    `json:"price" validate:"required"`
	Published time.Time `json:"published" validate:"required"`
	HasImage  bool      `json:"hasImage" validate:"required"`

	// Žánry
	Genres []Genre `gorm:"many2many:bookGenres;" json:"genres" validate:"required"`
	// Autor
	Author Author `json:"author" validate:"required"`
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
