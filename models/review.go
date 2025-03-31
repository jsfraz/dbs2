package models

type Review struct {
	ID       uint   `json:"id" validate:"required"`
	BookID   uint   `json:"-"`
	UserID   uint   `json:"-"`
	Stars    int    `json:"stars" validate:"required"`
	Text     string `json:"text" validate:"required"`
	Approved bool   `json:"approved" validate:"required"`

	// Kniha
	Book Book `json:"book" validate:"required"`
	// Uživatel
	User User `json:"user" validate:"required"`
}

// Vrátí novou recenzi.
//
//	@param bookId
//	@param userId
//	@param stars
//	@param text
//	@return *Review
func NewReview(bookId uint, userId uint, stars int, text string) *Review {
	return &Review{
		BookID: bookId,
		UserID: userId,
		Stars:  stars,
		Text:   text,
	}
}
