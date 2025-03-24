package handlers

import (
	"dbs2/database"
	"dbs2/models"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Přidá knihu do wishlistu.
//
//	@param c
//	@param request
//	@return error
func AddBookToWishlist(c *gin.Context, request *models.Id) error {
	// Načtení uživatele
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return errors.New("uživatel není v kontextu")
	}
	// Kontrola zda už je knniha v wishlistu
	bookInWishlist, err := database.IsBookInWishlist(request.Id, u.(*models.User).ID)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	if bookInWishlist {
		c.AbortWithStatus(409)
		return fmt.Errorf("kniha s ID %d již je v wishlistu", request.Id)
	}
	// Přidání do wishlistu
	err = database.AddBookToWishlist(request.Id, u.(*models.User))
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Odstraní knihu z wishlistu.
//
//	@param c
//	@param request
//	@return error
func RemoveBookFromWishlist(c *gin.Context, request *models.Id) error {
	// Načtení uživatele
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return errors.New("uživatel není v kontextu")
	}
	// Kontrola zda už je knniha v wishlistu
	bookInWishlist, err := database.IsBookInWishlist(request.Id, u.(*models.User).ID)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	if !bookInWishlist {
		c.AbortWithStatus(404)
		return fmt.Errorf("kniha s ID %d není v wishlistu", request.Id)
	}
	// Odstranění knihy z wishlistu
	err = database.RemoveBookFromWishlist(request.Id, u.(*models.User).ID)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Vrátí všechny knihy v wishlistu.
//
//	@param c
//	@param request
//	@return *[]models.Book
//	@return error
func GetAllBooksInWishlist(c *gin.Context) (*[]models.Book, error) {
	// Načtení uživatele
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return nil, errors.New("uživatel není v kontextu")
	}
	// Vrácení všech knih v wishlistu
	books, err := database.GetAllBooksInWishlist(u.(*models.User).ID)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return books, nil
}
