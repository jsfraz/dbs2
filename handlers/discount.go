package handlers

import (
	"dbs2/database"
	"dbs2/models"
	"dbs2/utils"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

// Vytvoří slevu.
//
//	@param c
//	@param discount
//	@return error
func CreateDiscount(c *gin.Context, required *models.CreateDiscount) error {
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return errors.New("uživatel není v kontextu")
	}
	// TODO zajistit aby se nevygeneroval duplicitní kód
	// Vytvoření slevového kódu platného měsíc, kód stojí 100 věrnostních bodů a má hodnotu 100 Kč
	err := database.CreateDiscount(u.(*models.User).ID, utils.GenerateRandomString(10), required.PointPrice, time.Now().Add(time.Hour*24*30))
	if err != nil {
		return err
	}
	return nil
}

// Vrátí všechny slevy zákazníka.
//
//	@param c
//	@return error
func GetAllCustomerDiscounts(c *gin.Context) ([]models.Discount, error) {
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return nil, errors.New("uživatel není v kontextu")
	}
	return database.GetAllCustomerDiscounts(u.(*models.User).ID)
}
