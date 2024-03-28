package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func InitEndpointsAndRun() {
	// Set up router
	router := gin.Default()
	initEndpoints(router)

	port := os.Getenv("PORT")
	address := fmt.Sprintf("localhost:%s", port)

	// Start the router
	err := router.Run(address)

	if err != nil {
		log.Fatal("Can't run awesome Todo API:", err)
	}
}
