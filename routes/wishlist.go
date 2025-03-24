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

// Routa wishlistu.
//
//	@param g
func WishlistRoute(g *fizz.RouterGroup) {
	grp := g.Group("wishlist", "Wishlist", "Seznam přání")
	// Autentifikační middleware
	grp.Use(middlewares.Auth)
	// Middleware povolující pouze zákazníka
	grp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleCustomer})
	})

	// Přidání knihy do seznamu přání
	grp.POST("", utils.CreateOperationOption("Přidání knihy do seznamu přání", true), tonic.Handler(handlers.AddBookToWishlist, 204))
	// Odstranění knihy z seznamu přání
	grp.DELETE("", utils.CreateOperationOption("Odstranění knihy z seznamu přání", true), tonic.Handler(handlers.RemoveBookFromWishlist, 204))
	// Vrácení všech knih v seznamu přání
	grp.GET("all", utils.CreateOperationOption("Vrácení všech knih v seznamu přání", true), tonic.Handler(handlers.GetAllBooksInWishlist, 200))
	// Vrátí zda je kniha v seznamu přání
	grp.GET("exists", utils.CreateOperationOption("Vrátí zda je kniha v seznamu přání", true), tonic.Handler(handlers.IsBookInWishlist, 200))
}
