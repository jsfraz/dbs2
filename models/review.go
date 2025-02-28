package models

type Review struct {
	ID    uint
	BookId uint
	UserId uint
	Stars int
	Text  string

	// Kniha
	Book     Book
	// Uživatel
	User     User
}
