package handlers

import (
	"dbs2/database"
	"dbs2/models"
	"errors"

	"github.com/gin-gonic/gin"
)

// Vrátí všechny adresy zákazníka.
//
//	@param c
//	@return *[]models.Address
//	@return error
func GetAllCustomerAddresses(c *gin.Context) (*[]models.Address, error) {
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return nil, errors.New("uživatel není v kontextu")
	}
	addresses, err := database.GetAllCustomerAddresses(u.(*models.User).ID)
	if err != nil {
		return nil, err
	}
	return &addresses, nil
}

// Vytvoří novou adresu.
//
//	@param c
//	@param address
//	@return error
func CreateAddress(c *gin.Context, address *models.CreateAddress) error {
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return errors.New("uživatel není v kontextu")
	}
	newAddress := models.NewAddress(address.City, address.Street, address.PostCode, u.(*models.User).ID)
	err := database.CreateAddress(newAddress)
	if err != nil {
		return err
	}
	return nil
}
