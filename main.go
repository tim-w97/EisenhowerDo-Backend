package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tim-w97/my-awesome-Todo-API/handlers"
	"log"
	"os"
)

func main() {
	router := gin.Default()

	// routes for registration and login
	router.POST("/login", handlers.Login)

	// routes for handling Todo items
	router.GET("/todos", handlers.GetTodos)
	router.GET("/todos/:id", handlers.GetTodoByID)
	router.POST("/todos", handlers.AddTodo)

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
