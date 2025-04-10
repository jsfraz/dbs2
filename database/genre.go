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
	err := utils.GetSingleton().PostgresDb.Model(&models.Genre{}).Order("id ASC").Find(&genres).Error
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

// Vrátí žánry podle ID.
//
//	@param ids
//	@return *[]models.Genre
//	@return error
func GetGenresByIds(ids []uint) (*[]models.Genre, error) {
	var genres []models.Genre = []models.Genre{}
	err := utils.GetSingleton().PostgresDb.Model(&models.Genre{}).Where("id IN ?", ids).Order("id ASC").Find(&genres).Error
	if err != nil {
		return nil, err
	}
	return &genres, nil
}
