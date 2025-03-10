package routes

import (
	"dbs2/handlers"
	"dbs2/utils"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

// Přihlašovací routa.
//
//	@param grp
func AuthRoute(g *fizz.RouterGroup) {
	grp := g.Group("auth", "Authentication", "Přihlášení uživatele")

	// Přihlášení
	grp.POST("",
		utils.CreateOperationOption(
			"Přihlášení",
			false),
		tonic.Handler(handlers.Login, 200))

	// Registrace zákazníka
	grp.POST("register",
		utils.CreateOperationOption(
			"Registrace zákazníka",
			false),
		tonic.Handler(handlers.Register, 200))
}
