package models

type CustomerActivity struct {
	UserID             uint    `json:"userId" db:"user_id"`
	FullName           string  `json:"fullName" db:"full_name"`
	Email              string  `json:"email" db:"email"`
	TotalOrders        int     `json:"totalOrders" db:"total_orders"`
	TotalCartBooks     int     `json:"totalCartBooks" db:"total_cart_books"`
	TotalWishlistBooks int     `json:"totalWishlistBooks" db:"total_wishlist_books"`
	TotalReviews       int     `json:"totalReviews" db:"total_reviews"`
	TotalSpent         float64 `json:"totalSpent" db:"total_spent"`
	Points             int     `json:"points" db:"points"`
}
