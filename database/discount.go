package database

import (
	"dbs2/models"
	"dbs2/utils"
	"time"
)

// Vytvoření slevy.
//
//	@param userId
//	@param code
//	@param price
//	@param validUntil
//	@return error
func CreateDiscount(userId uint, code string, price int, validUntil time.Time) error {
	err := utils.GetSingleton().PostgresDb.Exec("SELECT purchase_discount($1, $2, $3, $4)", userId, code, price, validUntil).Error
	if err != nil {
		return err
	}
	return nil
}

// Vrátí všechny slevy zákazníka.
//
//	@param userId
//	@return []models.Discount
//	@return error
func GetAllCustomerDiscounts(userId uint) ([]models.Discount, error) {
	var discounts []models.Discount
	err := utils.GetSingleton().PostgresDb.Model(&models.Discount{}).Where("user_id = ? AND valid_until > ? AND used = ?", userId, time.Now(), false).Order("id ASC").Find(&discounts).Error
	return discounts, err
}

// Vrátí slevu.
//
//	@param id
//	@param userId
//	@return models.Discount
//	@return error
func GetDiscount(id uint, userId uint) (models.Discount, error) {
	var discount models.Discount
	err := utils.GetSingleton().PostgresDb.Model(&models.Discount{}).Where("id = ? AND user_id = ? AND valid_until > ? AND used = ?", id, userId, time.Now(), false).First(&discount).Error
	return discount, err
}
