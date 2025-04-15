package handlers

import (
	"dbs2/database"
	"dbs2/models"
	"errors"
	"slices"

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
		return u.(*models.User), nil
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

// Aktualizace uživatele.
//
//	@param c
//	@param request
//	@return error
func UpdateUser(c *gin.Context, request *models.UpdateUser) error {
	// Kontrola zda neupravuje sebe
	u, _ := c.Get("user")
	if u.(*models.User).ID == request.Id {
		c.AbortWithStatus(401)
		return errors.New("nelze editovat sebe")
	}
	// Kontrola zda existuje
	exists, err := database.UserExistsById(request.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	if !exists {
		c.AbortWithStatus(404)
		return errors.New("uživatel neexistuje")
	}
	// Aktualizace
	user, err := database.GetUserById(request.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	user.Update(request)
	err = database.UpdateUser(user)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Mazání uživatelů
//
//	@param c
//	@param request
//	@return error
func DeleteUsers(c *gin.Context, request *models.Ids) error {
	// Kontrola zda nemaže sebe
	u, _ := c.Get("user")
	if slices.Contains(request.Ids, u.(*models.User).ID) {
		c.AbortWithStatus(400)
		return errors.New("nelze mazat sebe")
	}
	// Smazání
	err := database.DeleteUsers(request.Ids)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Smazání uživatele.
//
//	@param c
//	@param request
//	@return error
func DeleteUser(c *gin.Context, request *models.Id) error {
	// Kontrola zda nemaže sebe
	u, _ := c.Get("user")
	if u.(*models.User).ID == request.Id {
		c.AbortWithStatus(400)
		return errors.New("nelze mazat sebe")
	}
	// Kontrola zda maže customera nebo reviewApprovera
	user, err := database.GetUserById(request.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	if user.Role != models.RoleCustomer && user.Role != models.RoleReview {
		c.AbortWithStatus(400)
		return errors.New("nelze mazat admina nebo databaseManager")
	}
	// Smazání
	err = database.DeleteUser(request.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}
