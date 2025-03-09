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
	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return models.NewLoginResponse(accessToken), nil
}
