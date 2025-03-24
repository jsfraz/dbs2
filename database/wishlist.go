package database

import (
	"dbs2/models"
	"dbs2/utils"
)

// Přidání knihy do wishlistu.
//
//	@param bookId
//	@param userId
//	@return error
func AddBookToWishlist(bookId uint, user *models.User) error {
	return utils.GetSingleton().PostgresDb.Exec("INSERT INTO wishlists (user_id, book_id) VALUES (?, ?)", user.ID, bookId).Error
}

// Odstraní knihu z wishlistu.
//
//	@param bookId
//	@param userId
//	@return error
func RemoveBookFromWishlist(bookId uint, userId uint) error {
	return utils.GetSingleton().PostgresDb.Exec("DELETE FROM wishlists WHERE user_id = ? AND book_id = ?", userId, bookId).Error
}

// Kontrola zda už je kniha ve wishlistu.
//
//	@param bookId
//	@param userId
//	@return bool
//	@return error
func IsBookInWishlist(bookId uint, userId uint) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Table("wishlists").
		Where("user_id = ? AND book_id = ?", userId, bookId).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// Vrátí všechny knihy v wishlistu.
//
//	@param userId
//	@return *[]models.Book
//	@return error
func GetAllBooksInWishlist(userId uint) (*[]models.Book, error) {
	var books []models.Book
	err := utils.GetSingleton().PostgresDb.Model(&models.Book{}).Preload("Genres").Preload("Author").Where("id IN (SELECT book_id FROM wishlists WHERE user_id = ?)", userId).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return &books, nil
}
