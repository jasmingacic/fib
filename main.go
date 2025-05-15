package main

import (
	"fib/controllers"
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var defaultPort = 8080

func main() {

	port := flag.Int("port", defaultPort, "port to run the server on")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	registerRoutes(router)

	// Start the server
	fmt.Printf("Server is running on port %d...\n", *port)
	log.Fatal(router.Run(fmt.Sprintf(":%d", *port)))
}

// registerRoutes registers all API routes
func registerRoutes(router *gin.Engine) {
	// Group API routes
	api := router.Group("/api")
	{
		api.GET("/fibonacci/:n", controllers.GetFibonacci)
	}
}
