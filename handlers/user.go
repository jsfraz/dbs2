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

// Vytvoření uživatele.
//
//	@param c
//	@param request
//	@return error
func CreateUser(c *gin.Context, request *models.CreateUser) error {
	// Kontrola zda uživatel existuje
	exists, err := database.UserExistsByMail(request.Mail)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	if exists {
		c.AbortWithStatus(409)
		return errors.New("uživatel s tímto e-mailem již existuje")
	}
	// Vytvoření uživatele
	u, err := models.NewUser(request.FirstName, request.LastName, request.Mail, request.Role, request.Password)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	err = database.CreateUser(u)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}
