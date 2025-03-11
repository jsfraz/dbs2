package models

type Ids struct {
	Ids []uint `query:"ids" validate:"required"`
}
