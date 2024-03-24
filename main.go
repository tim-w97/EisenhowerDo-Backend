package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	router := gin.Default()

	// routes
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)
	router.POST("/todos", addTodo)

	// load port from environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can't load the .env file")
	}

	port := os.Getenv("PORT")
	hostnameAndPort := fmt.Sprintf("localhost:%s", port)

	// run the awesome Todo API
	if err := router.Run(hostnameAndPort); err != nil {
		log.Fatal("Can't run awesome Todo API:", err)
	}
}
