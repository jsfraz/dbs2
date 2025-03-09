package handlers

import (
	"dbs2/models"
	"errors"

	"github.com/gin-gonic/gin"
)

// Vrátí profil uživatele podle ID uživatele. ID se získá z přístupového tokenu.
//
//	@param c
//	@return *models.User
//	@return error
func WhoAmI(c *gin.Context) (*models.User, error) {
	u, _ := c.Get("user")
	if u != nil {
		user := u.(*models.User)
		return user, nil
	} else {
		c.AbortWithStatus(500)
		return nil, errors.New("žádný uživatel v kontextu")
	}
}
