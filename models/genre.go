package models

type Genre struct {
	ID   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// Vrátí nový žánr.
//
//	@param name
//	@return *Genre
func NewGenre(name string) *Genre {
	return &Genre{
		Name: name,
	}
}
