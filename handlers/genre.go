package handlers

import (
	"dbs2/database"
	"dbs2/models"

	"github.com/gin-gonic/gin"
)

// Vytvoření žánru.
//
//	@param c
//	@param request
//	@return error
func CreateGenre(c *gin.Context, request *models.CreateGenre) error {
	// TODO kontrola duplicity jména???
	genre := models.NewGenre(request.Name)
	err := database.CreateGenre(genre)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Vrátí všechny žánry
//
//	@param c
//	@return *[]models.Genre
//	@return error
func GetAllGenres(c *gin.Context) (*[]models.Genre, error) {
	authors, err := database.GetAllGenres()
	if err != nil {
		return nil, err
	}
	return authors, nil
}
