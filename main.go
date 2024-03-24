package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tim-w97/my-awesome-Todo-API/handlers"
	"github.com/tim-w97/my-awesome-Todo-API/middleware"
	"log"
	"os"
)

func main() {
	router := gin.Default()

	// routes for registration and login
	router.POST("/login", handlers.Login)

	// routes for handling Todo items
	router.GET("/todos", middleware.RequireAuth, handlers.GetTodos)
	router.GET("/todos/:id", middleware.RequireAuth, handlers.GetTodoByID)
	router.POST("/todos", middleware.RequireAuth, handlers.AddTodo)

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
