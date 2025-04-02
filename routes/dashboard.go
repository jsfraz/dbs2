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

func DashboardRoute(g *fizz.RouterGroup) {
	grp := g.Group("dashboard", "Dashboard", "Dashboard")

	// Autentifikační middleware
	grp.Use(middlewares.Auth)
	// Middleware povolující pouze admina
	grp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleAdmin})
	})

	// Popularita knih
	grp.GET("bookPopularityStats", utils.CreateOperationOption("Popularita knih", true), tonic.Handler(handlers.BookPopularityStats, 200))
	// Statistiky knih
	grp.GET("bookStats", utils.CreateOperationOption("Statistiky knih", true), tonic.Handler(handlers.BookStats, 200))
	// Aktivita zákazníků
	grp.GET("customerActivity", utils.CreateOperationOption("Aktivita zákazníků", true), tonic.Handler(handlers.CustomerActivity, 200))
}
