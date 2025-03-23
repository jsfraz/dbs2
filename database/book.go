package database

import (
	"dbs2/models"
	"dbs2/utils"
)

// Vrátí zda kniha podle ISBN existuje.
//
//	@param isbn
//	@return bool
//	@return error
func BookExistsByIsbn(isbn string) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.Book{}).Where("isbn = ?", isbn).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// Vloží knihu do databáze.
//
//	@param book
//	@return error
func CreateBook(book *models.Book) error {
	err := utils.GetSingleton().PostgresDb.Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}

// Vrátí zda kniha podle ID existuje.
//
//	@param id
//	@return bool
//	@return error
func BookExistsById(id uint) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.Book{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// Vrátí knihu podle ID.
//
//	@param id
//	@return *models.Book
//	@return error
func GetBookById(id uint) (*models.Book, error) {
	var book models.Book
	err := utils.GetSingleton().PostgresDb.Model(&models.Book{}).Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// Aktualizace knihy.
//
//	@param book
//	@return error
func UpdateBook(book *models.Book) error {
	return utils.GetSingleton().PostgresDb.Save(book).Error
}

// Vrátí všechny knihy i s autory a žánry.
//
//	@return *[]models.Book
//	@return error
func GetAllBooks() (*[]models.Book, error) {
	var books []models.Book = []models.Book{}
	err := utils.GetSingleton().PostgresDb.Model(&models.Book{}).Preload("Genres").Preload("Author").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return &books, nil
}

// Vyhledávání knih podle různých kritérií.
//
//	@param searchBooks
//	@return *[]models.Book
//	@return error
func SearchBooks(searchBooks *models.SearchBooks) (*[]models.Book, error) {
	var books []models.Book = []models.Book{}
	tx := utils.GetSingleton().PostgresDb.Model(&models.Book{}).Preload("Genres").Preload("Author")
	if searchBooks.Name != nil {
		tx = tx.Where("name LIKE ?", "%"+*searchBooks.Name+"%")
	}
	if searchBooks.AuthorIds != nil {
		tx = tx.Where("author_id IN ?", *searchBooks.AuthorIds)
	}
	if searchBooks.GenreIds != nil {
		tx = tx.Where("genre_id IN ?", *searchBooks.GenreIds)
	}
	err := tx.Where("price >= ? AND price <= ?", searchBooks.MinPrice, searchBooks.MaxPrice).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return &books, nil
}
