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

// Routa pro objednávky.
//
//	@param g
func DiscountRoute(g *fizz.RouterGroup) {
	grp := g.Group("discount", "Discount", "Slevy")

	// Autentifikační middleware
	grp.Use(middlewares.Auth)

	// Middleware povolující pouze zákazníka
	grp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleCustomer})
	})

	// Vytvoření slevy
	grp.POST("", utils.CreateOperationOption("Vytvoření slevy", true), tonic.Handler(handlers.CreateDiscount, 200))
	// Získání všech slev zákazníka
	grp.GET("", utils.CreateOperationOption("Získání všech slev zákazníka", true), tonic.Handler(handlers.GetAllCustomerDiscounts, 200))
}
