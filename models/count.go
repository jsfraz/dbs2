package models

type Count struct {
	Count int64 `json:"count" validate:"required"`
}
