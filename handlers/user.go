package handlers

import (
	"dbs2/database"
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

// Vrátí uživtele podle role.
//
//	@param c
//	@param request
//	@return *[]models.User
//	@return error
func GetUsersByRoles(c *gin.Context, request *models.RolesRequest) (*[]models.User, error) {
	users, err := database.GetUsersByRole(request.Roles)
	if err != nil {
		return nil, err
	}
	return users, nil
}
