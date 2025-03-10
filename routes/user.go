package routes

import (
	"dbs2/handlers"
	"dbs2/middlewares"
	"dbs2/models"
	"dbs2/utils"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

// Uživatelská routa.
//
//	@param g
func UserRoute(g *fizz.RouterGroup) {
	grp := g.Group("user", "User", "Uživatelé")

	// Autentifikační middleware
	grp.Use(middlewares.Auth)

	// WhoAmI
	grp.GET("whoami", utils.CreateOperationOption("Kdo jsem?", true), tonic.Handler(handlers.WhoAmI, 200))

	// Routa pro management uživatelů
	mgmtGrp := grp.Group("management", "User management", "Management uživatelů - operace pro admina.")

	// Middleware povolující pouze admina
	mgmtGrp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleAdmin})
	})

	// Všichni uživatelé kteří nejsou zákaznící
	mgmtGrp.GET("byRoles", utils.CreateOperationOption("Všichni uživatelé kteří nejsou zákaznící", true), tonic.Handler(handlers.GetUsersByRoles, 200))
	// Vytvoření uživatele s rolí databaseManager nebo reviewApprover
	mgmtGrp.POST("user", utils.CreateOperationOption("Vytvoření uživatele s rolí databaseManager nebo reviewApprover", true), tonic.Handler(handlers.CreateUser, 204))
}
