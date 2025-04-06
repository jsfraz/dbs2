package handlers

import (
	"dbs2/database"
	"dbs2/models"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Přidá knihu do košíku.
//
//	@param c
//	@param request
//	@return error
func AddBookToCart(c *gin.Context, request *models.Id) error {
	// Načtení uživatele
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return errors.New("uživatel není v kontextu")
	}
	// Kontrola zda už je knniha v košíku
	bookInCart, err := database.IsBookInCart(request.Id, u.(*models.User).ID)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	if bookInCart {
		c.AbortWithStatus(409)
		return fmt.Errorf("kniha s ID %d již je v košíku", request.Id)
	}
	// Přidání do košíku
	err = database.AddBookToCart(request.Id, u.(*models.User))
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Odstraní knihu z košíku.
//
//	@param c
//	@param request
//	@return error
func RemoveBookFromCart(c *gin.Context, request *models.Id) error {
	// Načtení uživatele
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return errors.New("uživatel není v kontextu")
	}
	// Kontrola zda už je knniha v košíku
	bookInCart, err := database.IsBookInCart(request.Id, u.(*models.User).ID)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	if !bookInCart {
		c.AbortWithStatus(404)
		return fmt.Errorf("kniha s ID %d není v košíku", request.Id)
	}
	// Odstranění knihy z košíku
	err = database.RemoveBookFromCart(request.Id, u.(*models.User).ID)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Vrátí všechny knihy v košíku.
//
//	@param c
//	@param request
//	@return *[]models.Book
//	@return error
func GetAllBooksInCart(c *gin.Context) (*[]models.Book, error) {
	// Načtení uživatele
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return nil, errors.New("uživatel není v kontextu")
	}
	// Vrácení všech knih v košíku
	books, err := database.GetAllBooksInCart(u.(*models.User).ID)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return books, nil
}

// GVrátí počet knih v košíku.
//
//	@param c
//	@return *models.Count
//	@return error
func GetCartCount(c *gin.Context) (*models.Count, error) {
	// Načtení uživatele
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return nil, errors.New("uživatel není v kontextu")
	}
	// Vrácení počtu knih v košíku
	count, err := database.GetCartCount(u.(*models.User).ID)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return &models.Count{Count: count}, nil
}
