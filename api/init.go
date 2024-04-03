package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/middleware"
	"log"
	"os"
)

func InitEndpointsAndRun() {
	// Set up router
	router := gin.Default()

	router.Use(middleware.ConfigureCORS)
	initEndpoints(router)

	port := os.Getenv("PORT")
	address := fmt.Sprintf("localhost:%s", port)

	// Start the router
	err := router.Run(address)

	if err != nil {
		log.Fatal("Can't run awesome Todo API:", err)
	}
}
