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
	// Vyhledávání knih
	grp.GET("search", utils.CreateOperationOption("Vyhledávání knih", true), tonic.Handler(handlers.SearchBooks, 200))
	// Vrátí knihu podle ID
	grp.GET("", utils.CreateOperationOption("Vrátí knihu podle ID", true), tonic.Handler(handlers.GetBookById, 200))

	// Routa pro management knih
	mgmtGrp := grp.Group("management", "Book management", "Management knih - operace pro admina a database managera.")

	// Autentifikační middleware
	mgmtGrp.Use(middlewares.Auth)
	// Middleware povolující pouze admina a database managera
	mgmtGrp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleAdmin, models.RoleDbManager})
	})

	// Vytvoření nové knihy
	mgmtGrp.POST("book", utils.CreateOperationOption("Vytvoření knihy", true), tonic.Handler(handlers.CreateBook, 200))
	// Aktualizace knihy
	mgmtGrp.PATCH("book", utils.CreateOperationOption("Aktualizace knihy", true), tonic.Handler(handlers.UpdateBook, 204))
	// Nahrání obrázku knihy
	mgmtGrp.GinRouterGroup().POST("bookImage", handlers.UploadBookImage)
	// Odstranění obrázku knihy
	mgmtGrp.DELETE("bookImage", utils.CreateOperationOption("Odstranění obrázku knihy", true), tonic.Handler(handlers.DeleteBookImage, 204))
	// Mazání knihy
	mgmtGrp.DELETE("book", utils.CreateOperationOption("Mazání knihy", true), tonic.Handler(handlers.DeleteBook, 204))
}
