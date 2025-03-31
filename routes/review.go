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

// Routa pro recenze.
//
//	@param g
func ReviewRoute(g *fizz.RouterGroup) {
	grp := g.Group("review", "Review", "Recenze")

	// Recenze podle knihy
	grp.GET("all", utils.CreateOperationOption("Všechny schválené recenze podle knihy.", true), tonic.Handler(handlers.GetApprovedReviewsByBookId, 200))

	// Routa pro zákaznické recenze
	userGrp := grp.Group("customer", "Customer review", "Recenze zákazníka")

	// Autentifikační middleware
	userGrp.Use(middlewares.Auth)

	// Middleware povolující pouze zákazníka
	userGrp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleCustomer})
	})

	// Vytvoření recenze
	userGrp.POST("", utils.CreateOperationOption("Vytvoření recenze", true), tonic.Handler(handlers.CreateReview, 204))
	// Zjištění zda se uživatelova recenze schvaluje
	userGrp.GET("isBeingApproved", utils.CreateOperationOption("Zjištění zda se uživatelova recenze schvaluje", true), tonic.Handler(handlers.IsUserReviewBeingApproved, 200))
	// TODO Aktualizace recenze
	// TODO Odstranění recenze

	// Routa pro management recenzí
	mgmtGrp := grp.Group("management", "Review management", "Management recenzí - operace pro admina a review approvera.")

	// Autentifikační middleware
	mgmtGrp.Use(middlewares.Auth)

	// Middleware povolující pouze admina a review approvera
	mgmtGrp.Use(func(c *gin.Context) {
		middlewares.Role(c, []models.Role{models.RoleAdmin, models.RoleReview})
	})

	// Schvalování/smazání recenzí
	mgmtGrp.POST("approve", utils.CreateOperationOption("Schvalování/smazání recenzí", true), tonic.Handler(handlers.ApproveReview, 204))
	// Vrátí recenze ke schválení
	mgmtGrp.GET("toApprove", utils.CreateOperationOption("Vrátí recenze ke schválení", true), tonic.Handler(handlers.GetReviewsToApprove, 200))
}
