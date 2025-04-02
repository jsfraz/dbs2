package models

type BookPopularityStat struct {
	BookId        string  `json:"book_id"`
	BookName      string  `json:"book_name"`
	AuthorName    string  `json:"author_name"`
	TotalReviews  int     `json:"total_reviews"`
	AverageRating float64 `json:"average_rating"`
	GenreName     string  `json:"genre_name"`
}
