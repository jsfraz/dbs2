package handlers

import (
	"dbs2/database"
	"dbs2/models"
	"dbs2/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Vytoření autora.
//
//	@param c
//	@param request
//	@return error
func CreateAuthor(c *gin.Context, request *models.CreateAuthor) error {
	// Validace narození autora
	date, err := utils.ParseISO8601String(request.Birth)
	if err != nil {
		c.AbortWithStatus(400)
		return fmt.Errorf("chyba parsování data narození: %s", err)
	}
	// Vytvoření autora
	author := models.NewAuthor(request.FirstName, request.LastName, *date)
	err = database.CreateAuthor(author)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Všichni autoři.
//
//	@param c
//	@return *[]models.Author
//	@return error
func GetAllAuthors(c *gin.Context) (*[]models.Author, error) {
	authors, err := database.GetAllAuthors()
	if err != nil {
		return nil, err
	}
	return authors, nil
}
