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

// AuthorRoute
//
//	@param g
func AuthorRoute(g *fizz.RouterGroup) {
	grp := g.Group("author", "Author", "Autoři")

	// Autentifikační middleware
	grp.Use(middlewares.Auth)

	// Všichni autoři
	grp.POST("all", utils.CreateOperationOption("Všichni autoři", true), tonic.Handler(handlers.GetAllAuthors, 200))

	// Routa pro management autorů
	mgmtGrp := grp.Group("management", "Author management", "Management autorů - operace pro admina.")

	// Middleware povolující pouze admina
	mgmtGrp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleAdmin})
	})

	// Vytvoření autora
	mgmtGrp.POST("author", utils.CreateOperationOption("Vytvoření autora", true), tonic.Handler(handlers.CreateAuthor, 204))
	// TODO aktualiazce autora
	// TODO odstarnění autora
}
