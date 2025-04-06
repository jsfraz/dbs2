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

// Routa košíku.
//
//	@param g
func CartRoute(g *fizz.RouterGroup) {
	grp := g.Group("cart", "Cart", "Košík")
	// Autentifikační middleware
	grp.Use(middlewares.Auth)
	// Middleware povolující pouze zákazníka
	grp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleCustomer})
	})

	// Přidání knihy do košíku
	grp.POST("", utils.CreateOperationOption("Přidání knihy do košíku", true), tonic.Handler(handlers.AddBookToCart, 204))
	// Odstranění knihy z košíku
	grp.DELETE("", utils.CreateOperationOption("Odstranění knihy z košíku", true), tonic.Handler(handlers.RemoveBookFromCart, 204))
	// Vrácení všech knih v košíku
	grp.GET("all", utils.CreateOperationOption("Vrácení všech knih v košíku", true), tonic.Handler(handlers.GetAllBooksInCart, 200))
	// Vrátí zda je kniha v košíku
	grp.GET("exists", utils.CreateOperationOption("Vrátí zda je kniha v košíku", true), tonic.Handler(handlers.IsBookInCart, 200))
	// Vrátí počet knih v košíku
	grp.GET("count", utils.CreateOperationOption("Vrátí počet knih v košíku", true), tonic.Handler(handlers.GetCartCount, 200))
}
