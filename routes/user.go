package routes

import (
	"dbs2/handlers"
	"dbs2/middlewares"
	"dbs2/utils"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

// Uživatelská routa.
//
//	@param g
func UserRoute(g *fizz.RouterGroup) {
	grp := g.Group("user", "User", "Uživatelé")

	// Autentifikační middleware
	grp.Use(middlewares.Auth)

	// WhoAmI
	grp.GET("whoami", utils.CreateOperationOption("Kdo jsem?", true), tonic.Handler(handlers.WhoAmI, 200))
}
