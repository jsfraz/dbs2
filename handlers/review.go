package handlers

import (
	"dbs2/database"
	"dbs2/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Vytvoří novou recenzi.
//
//	@param c
//	@param request
//	@return error
func CreateReview(c *gin.Context, request *models.CreateReview) error {
	u, _ := c.Get("user")
	// Kontrola zda kniha existuje
	bookExists, err := database.BookExistsById(request.BookId)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	if !bookExists {
		c.AbortWithStatus(404)
		return fmt.Errorf("kniha s ID %d neexistuje", request.BookId)
	}
	// Kontrola zda recenze už existuje
	reviewExists, err := database.ReviewExistsByBookIdAndUserId(request.BookId, u.(*models.User).ID)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	if reviewExists {
		c.AbortWithStatus(409)
		return fmt.Errorf("recenze už existuje")
	}
	// Uložení recenze
	review := models.NewReview(request.BookId, u.(*models.User).ID, request.Stars, request.Text)
	err = database.CreateReview(review)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Vrátí všechny schválené recenze podle knihy.
//
//	@param c
//	@param request
//	@return error
func GetApprovedReviewsByBookId(c *gin.Context, request *models.Id) ([]*models.Review, error) {
	// Kontrola zda kniha existuje
	bookExists, err := database.BookExistsById(request.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	if !bookExists {
		c.AbortWithStatus(404)
		return nil, fmt.Errorf("kniha s ID %d neexistuje", request.Id)
	}
	// Vrácení schválených recenzí
	reviews, err := database.GetApprovedReviewsByBookId(request.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return reviews, nil
}

// Schválí/smaže recenzi.
//
//	@param c
//	@param request
//	@return error
func ApproveReview(c *gin.Context, request *models.ApproveReview) error {
	// Kontrola zda recenze existuje
	reviewExists, err := database.ReviewExistsById(request.ReviewId)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	if !reviewExists {
		c.AbortWithStatus(404)
		return fmt.Errorf("recenze s ID %d neexistuje", request.ReviewId)
	}
	// TODO Kontrola zda je recenze schválená
	// Schválení/smazání recenze
	err = database.ApproveReview(request.ReviewId, request.Approved)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Vrátí recenze ke schválení.
//
//	@param c
//	@param request
//	@return error
func GetReviewsToApprove(c *gin.Context) ([]*models.Review, error) {
	reviews, err := database.GetReviewsToApprove()
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return reviews, nil
}
