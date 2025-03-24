package models

type UpdateBook struct {
	Id        uint   `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	AuthorId  uint   `json:"authorId" validate:"required"`
	Summary   string `json:"summary" validate:"required"`
	Isbn      string `json:"isbn" validate:"required"`
	Price     uint64 `json:"price" validate:"required"`
	Published string `json:"published" validate:"required"`
	GenreIds  []uint `json:"genreIds" validate:"required"`
}
