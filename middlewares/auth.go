package middlewares

import (
	"dbs2/database"
	"dbs2/utils"

	"github.com/gin-gonic/gin"
)

// Middleware for user authentication.
// If the user has a valid access token, it sets its ID in the context.
// If it is not valid, it returns a status of 401.
//
//	@param c Gin context
func Auth(c *gin.Context) {
	// Token z kontextu
	userId, err := utils.TokenValid(utils.ExtractTokenFromContext(c), utils.GetSingleton().Config.AccessTokenSecret)
	// Kontrola tokenu
	if err != nil {
		c.AbortWithStatus(401)
		c.Error(err)
	}
	exists, err := database.UserExistsById(userId)
	if err != nil {
		c.AbortWithStatus(500)
		c.Error(err)
	}
	if !exists {
		c.AbortWithStatus(401)
	}
	// Vrátí uživatele
	user, err := database.GetUserById(userId)
	if err != nil {
		c.AbortWithStatus(500)
		c.Error(err)
	}
	// Nastavení uživatele do kontextu
	c.Set("user", user)
	c.Next()
}
