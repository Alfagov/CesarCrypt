package server

import (
	"github.com/Alfagov/CesarCrypt/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/unseal")
	router.POST("/seal")

	storeRouter := router.Group("/:storeBackend/:backendName")
	{
		// create a new storeBackend
		storeRouter.POST("", handlers.NewStoreBackend(nil))
		// modify storeBackend info
		storeRouter.PUT("", nil)
		// get storeBackend info
		storeRouter.GET("", nil)

		secretRouter := router.Group("/secrets/:name")
		{
			// get store secrets
			secretRouter.GET("", nil)
			// create store secret
			secretRouter.POST("", nil)
			// delete store secret
			secretRouter.DELETE("/:id", nil)

			secretRouter.POST("/:command", nil)
			secretRouter.GET("/:command", nil)

			// allow user to access store secret
			secretRouter.POST("/:id/access", nil)
			// revoke user access to store secret
			secretRouter.DELETE("/:id/access", nil)
		}
	}

	generalAuthenticationRouter := router.Group("/auth")
	{
		// Enable/Disable password login
		generalAuthenticationRouter.POST("/passwordLogin", nil)
		// Get password login status
		generalAuthenticationRouter.GET("/passwordLogin", nil)
		// Login with method
		generalAuthenticationRouter.POST("/:authMethod/login", nil)

	}

	storeHandlerRouter := router.Group("/storeHandlers")
	{
		storeHandlerRouter.POST("/:name", nil)
		storeHandlerRouter.GET("/:name", nil)
	}

	return router
}
