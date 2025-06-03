package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/config"
	"github.com/m-shahjalal/onepolicy-api/internal/middleware"
	"github.com/m-shahjalal/onepolicy-api/internal/router"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
	config.InitMigration(config.DB)
}

func main() {
	// initialize the application
	app := gin.New()

	// middlewares
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middleware.Cors())

	// routers
	router.AuthRouter(app)
	router.RootRoutes(app)
	router.PolicyRoutes(app)

	// run the server
	if err := app.Run(":" + os.Getenv("PORT")); err != nil {
		println("Error starting server: %v", err)
	}
}
