package database

import (
	"dbs2/models"
	"dbs2/utils"
)

// Vytvoří objednávku.
//
//	@param userId
//	@param addressId
//	@param discountId
//	@return error
func CreateOrder(userId uint, addressId uint, discountId *uint) error {
	tx := utils.GetSingleton().PostgresDb.Begin()
	// Uživatele s jeho košíkem (books)
	var user models.User
	if err := tx.Preload("Cart").Where("id = ?", userId).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Cena
	var totalPrice uint64
	for _, book := range user.Cart {
		totalPrice += book.Price
	}
	// Sleva
	if discountId != nil {
		discount, err := GetDiscount(*discountId, userId)
		if err != nil {
			tx.Rollback()
			return err
		}
		totalPrice -= discount.Price
	}
	// Označení slevy jako použité
	if discountId != nil {
		if err := tx.Model(&models.Discount{}).Where("id = ?", *discountId).Update("used", true).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	// Vytvoření objednávky
	order := models.Order{
		UserID:       userId,
		AddressID:    addressId,
		DiscountID:   discountId,
		TotalPrice:   totalPrice,
		Status:       "pending",
		OrderedBooks: user.Cart,
	}
	// Objednávka včetně vazeb na knihy
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Smazání košíku
	if err := tx.Model(&user).Association("Cart").Clear(); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
