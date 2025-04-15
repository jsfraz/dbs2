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
func OrderRoute(g *fizz.RouterGroup) {
	grp := g.Group("order", "Order", "Objednávky")

	// Autentifikační middleware
	grp.Use(middlewares.Auth)

	// Middleware povolující pouze zákazníka
	grp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleCustomer})
	})

	// Vytvoření objednávky
	grp.POST("", utils.CreateOperationOption("Vytvoření objednávky", true), tonic.Handler(handlers.CreateOrder, 204))
	// Vrátí všechny objednávky uživatele
	grp.GET("all", utils.CreateOperationOption("Vrátí všechny objednávky uživatele", true), tonic.Handler(handlers.GetAllOrders, 200))
}
