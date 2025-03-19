package database

import (
	"dbs2/models"
	"dbs2/utils"
)

// Vytvoření žánru.
//
//	@param genre
//	@return error
func CreateGenre(genre *models.Genre) error {
	err := utils.GetSingleton().PostgresDb.Create(&genre).Error
	if err != nil {
		return err
	}
	return nil
}

// Vrátí všechny žánry.
//
//	@return *[]models.Genre
//	@return error
func GetAllGenres() (*[]models.Genre, error) {
	var genres []models.Genre = []models.Genre{}
	err := utils.GetSingleton().PostgresDb.Model(&models.Genre{}).Find(&genres).Error
	if err != nil {
		return nil, err
	}
	return &genres, nil
}

// Vrtí zda žánr exituje podle ID.
//
//	@param id
//	@return bool
//	@return error
func GenreExistsById(id uint) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.Genre{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}
