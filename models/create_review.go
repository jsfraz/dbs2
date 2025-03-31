package models

type CreateReview struct {
	BookId uint   `json:"bookId" validate:"required"`
	Stars  int    `json:"stars" validate:"required"`
	Text   string `json:"text" validate:"required"`
}
