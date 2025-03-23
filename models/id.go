package models

type Id struct {
	Id uint `query:"id" validate:"required"`
}
