package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/config"
	"github.com/m-shahjalal/onepolicy-api/internal/router"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	port := os.Getenv("PORT")

	app := gin.New()

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	router.AuthRouter(app)
	router.RootRoutes(app)

	fmt.Printf("Server is running on http://localhost:%s\n", port)
	err := app.Run(":" + port)

	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
