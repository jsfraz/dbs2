package database

import (
	"dbs2/models"
	"dbs2/utils"
)

// Zjistí zda recenze už existuje podle knihy a uživatele.
//
//	@param bookId
//	@param userId
//	@return bool
//	@return error
func ReviewExistsByBookIdAndUserId(bookId uint, userId uint) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.Review{}).Where("book_id = ? AND user_id = ?", bookId, userId).Count(&count).Error
	return count == 1, err
}

// Vytvoří novou recenzi.
//
//	@param review
//	@return error
func CreateReview(review *models.Review) error {
	return utils.GetSingleton().PostgresDb.Create(review).Error
}

// Vrátí všechny schválené recenze podle knihy.
//
//	@param bookId
//	@return []*models.Review
//	@return error
func GetApprovedReviewsByBookId(bookId uint) ([]*models.Review, error) {
	var reviews []*models.Review
	err := utils.GetSingleton().PostgresDb.Where("book_id = ? AND approved = ?", bookId, true).Preload("User").Preload("Book").Order("id ASC").Find(&reviews).Error
	return reviews, err
}

// Schválí recenzi.
//
//	@param reviewId
//	@param approved
//	@return error
func ApproveReview(reviewId uint, approved bool) error {
	// Shválení recenze
	if approved {
		return utils.GetSingleton().PostgresDb.Model(&models.Review{}).Where("id = ?", reviewId).Update("approved", approved).Error
	}
	// Smazání recenze
	return utils.GetSingleton().PostgresDb.Where("id = ?", reviewId).Delete(&models.Review{}).Error
}

// Vrátí recenze ke schválení.
//
//	@return []*models.Review
//	@return error
func GetReviewsToApprove() ([]*models.Review, error) {
	var reviews []*models.Review
	err := utils.GetSingleton().PostgresDb.Where("approved = ?", false).Preload("User").Preload("Book").Order("id ASC").Find(&reviews).Error
	return reviews, err
}

// Zjistí zda recenze existuje podle ID.
//
//	@param reviewId
//	@return bool
//	@return error
func ReviewExistsById(reviewId uint) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.Review{}).Where("id = ?", reviewId).Count(&count).Error
	return count == 1, err
}

// IsReviewBeingApproved Zjištění zda se recenze schvaluje.
//
//	@param bookId
//	@param userId
//	@return bool
//	@return error
func IsReviewBeingApproved(bookId uint, userId uint) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.Review{}).Where("book_id = ? AND user_id = ? AND approved = ?", bookId, userId, false).Count(&count).Error
	return count == 1, err
}
