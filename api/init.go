package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/handlers"
	"github.com/tim-w97/my-awesome-Todo-API/middleware"
	"log"
	"os"
)

func InitRoutesAndRun() {
	// Set up router and routes
	router := gin.Default()

	// routes for registration and login
	router.POST("/login", handlers.Login)

	// routes for handling Todo items
	router.GET("/todos", middleware.RequireAuth, handlers.GetTodos)
	router.GET("/todos/:id", middleware.RequireAuth, handlers.GetTodoByID)
	router.POST("/todos", middleware.RequireAuth, handlers.AddTodo)

	port := os.Getenv("PORT")
	address := fmt.Sprintf("localhost:%s", port)

	// Start the router
	err := router.Run(address)

	if err != nil {
		log.Fatal("Can't run awesome Todo API:", err)
	}
}
