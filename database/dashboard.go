package database

import (
	"dbs2/models"
	"dbs2/utils"
)

// Zavolá pohled book_popularity_stats
//
//	@return []models.BookPopularityStat
//	@return error
func BookPopularityStats() ([]models.BookPopularityStat, error) {
	var stats []models.BookPopularityStat = []models.BookPopularityStat{}
	err := utils.GetSingleton().PostgresDb.Raw("SELECT * FROM book_popularity_stats").Scan(&stats).Error
	if err != nil {
		return nil, err
	}
	return stats, nil
}

// Zavolá pohled book_stats
//
//	@return []models.BookStat
//	@return error
func BookStats() ([]models.BookStat, error) {
	var stats []models.BookStat = []models.BookStat{}
	err := utils.GetSingleton().PostgresDb.Raw("SELECT * FROM book_stats").Scan(&stats).Error
	if err != nil {
		return nil, err
	}
	return stats, nil
}

// Zavolá pohled customer_activity
//
//	@return []models.CustomerActivity
//	@return error
func CustomerActivity() ([]models.CustomerActivity, error) {
	var activities []models.CustomerActivity = []models.CustomerActivity{}
	err := utils.GetSingleton().PostgresDb.Raw("SELECT * FROM customer_activity").Scan(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}
