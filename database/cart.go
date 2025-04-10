package database

import (
	"dbs2/models"
	"dbs2/utils"
)

// Přidání knihy do košíku.
//
//	@param bookId
//	@param userId
//	@return error
func AddBookToCart(bookId uint, user *models.User) error {
	return utils.GetSingleton().PostgresDb.Exec("INSERT INTO carts (user_id, book_id) VALUES (?, ?)", user.ID, bookId).Error
}

// Odstraní knihu z košíku.
//
//	@param bookId
//	@param userId
//	@return error
func RemoveBookFromCart(bookId uint, userId uint) error {
	return utils.GetSingleton().PostgresDb.Exec("DELETE FROM carts WHERE user_id = ? AND book_id = ?", userId, bookId).Error
}

// Kontrola zda už je kniha v košíku.
//
//	@param bookId
//	@param userId
//	@return bool
//	@return error
func IsBookInCart(bookId uint, userId uint) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Table("carts").
		Where("user_id = ? AND book_id = ?", userId, bookId).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// Vrátí všechny knihy v košíku.
//
//	@param userId
//	@return *[]models.Book
//	@return error
func GetAllBooksInCart(userId uint) (*[]models.Book, error) {
	var books []models.Book
	err := utils.GetSingleton().PostgresDb.Model(&models.Book{}).Preload("Genres").Preload("Author").Where("id IN (SELECT book_id FROM carts WHERE user_id = ?)", userId).Order("id ASC").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return &books, nil
}

// Vrátí počet knih v košíku.
//
//	@param userId
//	@return int
//	@return error
func GetCartCount(userId uint) (int64, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Raw("SELECT get_user_cart_count(?)", userId).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
