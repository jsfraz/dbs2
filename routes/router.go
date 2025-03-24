package routes

import (
	"dbs2/utils"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

// Výchozí router.
//
//	@return *fizz.Fizz
//	@return error
func NewRouter() (*fizz.Fizz, error) {
	// Instance Ginu
	engine := gin.Default()
	// Defaultní cors config, Allow Origin, Authorization header
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	engine.Use(cors.New(config))
	// fizz instance
	fizz := fizz.NewFromEngine(engine)
	// security
	fizz.Generator().SetSecuritySchemes(map[string]*openapi.SecuritySchemeOrRef{
		"bearerAuth": {
			SecurityScheme: &openapi.SecurityScheme{
				Type:         "http",
				Scheme:       "bearer",
				BearerFormat: "JWT",
			},
		},
	})

	// Servery
	fizz.Generator().SetServers([]*openapi.Server{
		{
			Description: utils.GetSingleton().Config.AppUrl,
			URL:         utils.GetSingleton().Config.AppUrl,
		},
	})

	// OpenApi info
	infos := &openapi.Info{
		Title:       "dbs2",
		Description: "Backend pro projekt pro KIKM/DBS2 a KIT/TNPW2.",
	}

	// Základní API routa
	grp := fizz.Group("api", "", "")

	// OpenAPI
	if utils.GetSingleton().Config.Swagger {
		grp.GET("openapi.json", nil, fizz.OpenAPI(infos, "json"))
		// Swagger UI (https://github.com/swagger-api/swagger-ui/blob/HEAD/docs/usage/installation.md#unpkg)
		engine.LoadHTMLGlob("html/*.html")
		engine.GET("/swagger", func(c *gin.Context) {
			c.HTML(200, "swagger.html", gin.H{
				"url": fmt.Sprintf("%s/api/openapi.json", utils.GetSingleton().Config.AppUrl),
			})
		})
		/*
			engine.GET("/", func(c *gin.Context) {
				c.Redirect(301, fmt.Sprintf("%s/swagger", utils.GetSingleton().Config.AppUrl))
			})
		*/
	}

	// Ostatní routy
	AuthRoute(grp)
	UserRoute(grp)
	AuthorRoute(grp)
	GenreRoute(grp)
	BookRoute(grp)
	CartRoute(grp)
	WishlistRoute(grp)

	if len(fizz.Errors()) != 0 {
		return nil, fmt.Errorf("errors: %v", fizz.Errors())
	}
	return fizz, nil
}
