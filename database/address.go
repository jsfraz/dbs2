package database

import (
	"dbs2/models"
	"dbs2/utils"
)

// Vrátí všechny adresy zákazníka.
//
//	@param userId
//	@return []models.Address
//	@return error
func GetAllCustomerAddresses(userId uint) ([]models.Address, error) {
	var addresses []models.Address
	err := utils.GetSingleton().PostgresDb.Model(&models.Address{}).Where("user_id = ?", userId).Order("id ASC").Find(&addresses).Error
	return addresses, err
}

// Vytvoří novou adresu.
//
//	@param address
//	@return error
func CreateAddress(address *models.Address) error {
	return utils.GetSingleton().PostgresDb.Create(address).Error
}

// Vrátí adresu podle ID.
//
//	@param id
//	@return models.Address
//	@return error
func GetAddressById(id uint) (models.Address, error) {
	var address models.Address
	err := utils.GetSingleton().PostgresDb.Model(&models.Address{}).Where("id = ?", id).First(&address).Error
	return address, err
}
