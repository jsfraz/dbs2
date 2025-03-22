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

	// Všechny knihy
	grp.GET("all", utils.CreateOperationOption("Všechny knihy", true), tonic.Handler(handlers.GetAllBooks, 200))
	// Obrázek knihy
	grp.GinRouterGroup().GET("/image/:id", handlers.GetBookImage)

	// Routa pro management knih
	mgmtGrp := grp.Group("management", "Book management", "Management knih - operace pro admina.")

	// Autentifikační middleware
	mgmtGrp.Use(middlewares.Auth)
	// Middleware povolující pouze admina
	mgmtGrp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleAdmin})
	})

	// Vytvoření nové knihy
	mgmtGrp.POST("book", utils.CreateOperationOption("Vytvoření knihy", true), tonic.Handler(handlers.CreateBook, 200))
	// Nahrání obrázku knihy
	mgmtGrp.GinRouterGroup().POST("bookImage", handlers.UploadBookImage)
}
