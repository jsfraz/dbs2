package models

type CreateDiscount struct {
	PointPrice int `json:"pointPrice" validate:"required,min=100"`
}
