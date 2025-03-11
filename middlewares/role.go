package middlewares

import (
	"dbs2/models"
	"errors"
	"fmt"
	"slices"

	"github.com/gin-gonic/gin"
)

// Middleware vracející status 401 pokud uživatel nemá danou roli/role.
//
//	@param c
//	@param role
func Role(c *gin.Context, roles []models.Role) {
	// Načtení uživatele
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		c.Error(errors.New("uživatel není v kontextu"))
		return
	}
	// Kontrola role
	if !slices.Contains(roles, u.(*models.User).Role) {
		c.AbortWithStatus(401)
		c.Error(fmt.Errorf("nesprávná uživatelská role: %s, vyžadováno jedno z: %s", u.(*models.User).Role, roles))
		return
	}
	c.Next()
}
