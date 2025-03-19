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

// Routa knih.
//
//	@param g
func BookRoute(g *fizz.RouterGroup) {
	grp := g.Group("book", "Book", "Knihy")

	// Autentifikační middleware
	grp.Use(middlewares.Auth)

	// Routa pro management knih
	mgmtGrp := grp.Group("management", "Book management", "Management knih - operace pro admina.")

	// Middleware povolující pouze admina
	mgmtGrp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleAdmin})
	})

	// Vytvoření nové knihy
	mgmtGrp.POST("book", utils.CreateOperationOption("Vytvoření knihy", true), tonic.Handler(handlers.CreateBook, 200))
	// Nahrání obrázku knihy
	grp.GinRouterGroup().POST("bookImage", handlers.UploadBookImage)
}
