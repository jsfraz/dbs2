package handlers

import (
	"dbs2/database"
	"dbs2/models"
	"dbs2/utils"
	"encoding/base64"
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User login.
//
//	@param c
//	@param login
//	@return *models.LoginResponse
//	@return error
func Login(c *gin.Context, login *models.Login) (*models.LoginResponse, error) {
	// Nalezení uživatele
	user, err := database.GetUserByMail(login.Mail)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	// Nenalezen
	if user.Mail == "" {
		c.AbortWithStatus(401)
		return nil, errors.New("uživatel nenalezen")
	}
	// Kontrola hesla
	hashBytes, _ := base64.StdEncoding.DecodeString(user.PasswordHash)
	err = bcrypt.CompareHashAndPassword(hashBytes, []byte(login.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.AbortWithStatus(401)
			return nil, errors.New("špatné heslo")
		} else {
			c.AbortWithStatus(500)
			return nil, err
		}
	}
	// Vygenerování tokenu
	accessToken, err := utils.GenerateAccessToken(user)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return models.NewLoginResponse(accessToken), nil
}

// Registrace zákazníka.
//
//	@param c
//	@param request
//	@return error
func Register(c *gin.Context, request *models.Register) (*models.LoginResponse, error) {
	// Kontrola zda uživatel existuje
	exists, err := database.UserExistsByMail(request.Mail)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	if exists {
		c.AbortWithStatus(409)
		return nil, errors.New("uživatel s tímto e-mailem již existuje")
	}
	// Vytvoření uživatele
	u, err := models.NewUser(request.FirstName, request.LastName, request.Mail, models.RoleCustomer, request.Password)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	err = database.CreateUser(u)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	// Vygenerování tokenu
	accessToken, err := utils.GenerateAccessToken(u)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return models.NewLoginResponse(accessToken), nil
}
