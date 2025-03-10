package middlewares

import (
	"dbs2/database"
	"dbs2/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

// Middleware pro ověřování uživatelů.
// Pokud má uživatel platný přístupový token, nastaví jeho ID v kontextu.
// Pokud není platný, vrátí stav 401.
//
//	@param c Gin context
func Auth(c *gin.Context) {
	// Token z kontextu
	userId, err := utils.TokenValid(utils.ExtractTokenFromContext(c), utils.GetSingleton().Config.AccessTokenSecret)
	// Kontrola tokenu
	if err != nil {
		c.AbortWithStatus(401)
		c.Error(err)
		return
	}
	exists, err := database.UserExistsById(userId)
	if err != nil {
		c.AbortWithStatus(500)
		c.Error(err)
		return
	}
	if !exists {
		c.AbortWithStatus(401)
		c.Error(errors.New("uživatel neexistuje"))
		return
	}
	// Vrátí uživatele
	user, err := database.GetUserById(userId)
	if err != nil {
		c.AbortWithStatus(500)
		c.Error(err)
		return
	}
	// Nastavení uživatele do kontextu
	c.Set("user", user)
	c.Next()
}
