package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type SearchBooks struct {
	Name      *string `query:"name"`
	AuthorIds *[]uint `query:"authorIds"`
	GenreIds  *[]uint `query:"genreIds"`
	MinPrice  uint64  `query:"minPrice" validate:"required,min=100,max=1000"`
	MaxPrice  uint64  `query:"maxPrice" validate:"required,min=100,max=1000"`
}

// Validate pro SearchBooks
//
//	@return error
func (s *SearchBooks) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return err
	}

	// Kontrola, že MinPrice je menší nebo rovno MaxPrice
	if s.MinPrice > s.MaxPrice {
		return fmt.Errorf("MinPrice musí být menší nebo se rovnat MaxPrice")
	}

	return nil
}
