package handlers

import (
	"dbs2/database"
	"dbs2/models"
	"errors"

	"github.com/gin-gonic/gin"
)

// Vytvoří objednávku.
//
//	@param c
//	@param order
//	@return error
func CreateOrder(c *gin.Context, order *models.CreateOrder) error {
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return errors.New("uživatel není v kontextu")
	}
	// Kontrola jestli adresa patří uživateli
	address, err := database.GetAddressById(order.AddressId)
	if err != nil {
		return err
	}
	if address.UserID != u.(*models.User).ID {
		return errors.New("adresa nepatří uživateli")
	}
	// TODO vymyslet nějaké podmínky pro slevy a kromě frontendu je validovat i zde
	// Vytvoří objednávku
	err = database.CreateOrder(u.(*models.User).ID, order.AddressId, order.DiscountId)
	if err != nil {
		return err
	}
	return nil
}

// Vrátí všechny objednávky uživatele
//
//	@param c
//	@return error
func GetAllOrders(c *gin.Context) error {
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return errors.New("uživatel není v kontextu")
	}
	orders, err := database.GetAllOrders(u.(*models.User).ID)
	if err != nil {
		return err
	}
	c.JSON(200, orders)
	return nil
}
