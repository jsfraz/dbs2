package models

type Review struct {
	ID     uint
	UserID uint
	BookID uint
	Stars  int
	Text   string
}
