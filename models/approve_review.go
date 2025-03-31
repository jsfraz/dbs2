package models

type ApproveReview struct {
	ReviewId uint `json:"reviewId"`
	Approved bool `json:"approved"`
}
