package database

import (
	"dbs2/models"
	"dbs2/utils"
)

// Vytvoření autora.
//
//	@param author
//	@return error
func CreateAuthor(author *models.Author) error {
	err := utils.GetSingleton().PostgresDb.Create(&author).Error
	if err != nil {
		return err
	}
	return nil
}

// Vrátí všechny autory.
//
//	@return *[]models.Author
//	@return error
func GetAllAuthors() (*[]models.Author, error) {
	var authors []models.Author = []models.Author{}
	err := utils.GetSingleton().PostgresDb.Model(&models.Author{}).Order("id ASC").Find(&authors).Error
	if err != nil {
		return nil, err
	}
	return &authors, nil
}

// Vrátí zda autor existuje podle ID.
//
//	@param id
//	@return bool
//	@return error
func AuthorExistsById(id uint) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.Author{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}
