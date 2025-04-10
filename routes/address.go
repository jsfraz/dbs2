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

// Routa pro adresy.
//
//	@param g
func AddressRoute(g *fizz.RouterGroup) {
	grp := g.Group("address", "Address", "Adresy")

	// Autentifikační middleware
	grp.Use(middlewares.Auth)

	// Middleware povolující pouze zákazníka
	grp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleCustomer})
	})

	// Všechny adresy zákazníka
	grp.GET("", utils.CreateOperationOption("Všechny adresy zákazníka", true), tonic.Handler(handlers.GetAllCustomerAddresses, 200))
	// Vytvoření adresy
	grp.POST("", utils.CreateOperationOption("Vytvoření adresy", true), tonic.Handler(handlers.CreateAddress, 200))
}
