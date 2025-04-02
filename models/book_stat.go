package models

type BookStat struct {
	BookID        uint    `json:"bookId" db:"book_id"`
	BookName      string  `json:"bookName" db:"book_name"`
	AuthorName    string  `json:"authorName" db:"author_name"`
	TotalSales    int     `json:"totalSales" db:"total_sales"`
	TotalRevenue  float64 `json:"totalRevenue" db:"total_revenue"`
	AverageRating float32 `json:"averageRating" db:"average_rating"`
	TotalReviews  int     `json:"totalReviews" db:"total_reviews"`
}
