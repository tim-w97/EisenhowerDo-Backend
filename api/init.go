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

	authorized := router.Group("/")

	authorized.Use(middleware.JWTAuth())
	{
		// routes secured with JSON Web Token
		authorized.GET("/todos", handlers.GetTodos)
		authorized.GET("/todos/:id", handlers.GetTodoByID)

		authorized.POST("/todos", handlers.AddTodo)
	}

	port := os.Getenv("PORT")
	address := fmt.Sprintf("localhost:%s", port)

	// Start the router
	err := router.Run(address)

	if err != nil {
		log.Fatal("Can't run awesome Todo API:", err)
	}
}
