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

// Routa žándrů.
//
//	@param g
func GenreRoute(g *fizz.RouterGroup) {
	grp := g.Group("genre", "Genre", "Žánry")

	// Autentifikační middleware
	grp.Use(middlewares.Auth)

	// Všechny žánry
	grp.POST("all", utils.CreateOperationOption("Všechny žánry", true), tonic.Handler(handlers.GetAllGenres, 200))

	// Routa pro management žánrů
	mgmtGrp := grp.Group("management", "Genre management", "Management žánrů - operace pro admina.")

	// Middleware povolující pouze admina
	mgmtGrp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleAdmin})
	})

	// Vytvoření žánru
	mgmtGrp.POST("genre", utils.CreateOperationOption("Vytvoření žánru", true), tonic.Handler(handlers.CreateGenre, 204))
	// TODO aktualiazce žánru
	// TODO odstarnění žánru
}
