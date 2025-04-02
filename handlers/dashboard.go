package handlers

import (
	"dbs2/database"
	"dbs2/models"

	"github.com/gin-gonic/gin"
)

// BookPopularityStats
//
//	@param c
//	@return *[]models.BookPopularityStat
//	@return error
func BookPopularityStats(c *gin.Context) (*[]models.BookPopularityStat, error) {
	stats, err := database.BookPopularityStats()
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

// BookStats
//
//	@param c
//	@return *[]models.BookStat
//	@return error
func BookStats(c *gin.Context) (*[]models.BookStat, error) {
	stats, err := database.BookStats()
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

// CustomerActivity
//
//	@param c
//	@return *[]models.CustomerActivity
//	@return error
func CustomerActivity(c *gin.Context) (*[]models.CustomerActivity, error) {
	activities, err := database.CustomerActivity()
	if err != nil {
		return nil, err
	}
	return &activities, nil
}
