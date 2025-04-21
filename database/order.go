package database

import (
	"dbs2/models"
	"dbs2/utils"
	"errors"
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
	if len(user.Cart) == 0 {
		return errors.New("košík je prázdný")
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
		if discount.Price > totalPrice {
			totalPrice = 0
		} else {
			totalPrice -= discount.Price
		}
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

// Vrátí všechny objednávky uživatele
//
//	@param userId
//	@return *[]models.Order
//	@return error
func GetAllOrders(userId uint) (*[]models.Order, error) {
	var orders []models.Order
	if err := utils.GetSingleton().PostgresDb.Preload("OrderedBooks").Preload("OrderedBooks.Author").Preload("OrderedBooks.Genres").Preload("Address").Preload("Discount").Where("user_id = ?", userId).Find(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil
}
