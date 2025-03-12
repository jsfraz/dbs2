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
