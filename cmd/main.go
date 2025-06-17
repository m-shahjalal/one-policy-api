package main

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/config"
	"github.com/m-shahjalal/onepolicy-api/internal/middleware"
	"github.com/m-shahjalal/onepolicy-api/internal/router"
	"github.com/m-shahjalal/onepolicy-api/internal/seed"
	"os"
)

func init() {
	config.LoadEnvVariables()
	config.DatabaseConnection()
	config.InitMigration(config.DB)
	seed.SeedData(config.DB)
}

func main() {
	app := gin.Default()
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
