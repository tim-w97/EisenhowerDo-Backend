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
	// Set up router
	router := gin.Default()

	// Routes for Registration and Login
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	authorized := router.Group("/")

	// Routes secured with JSON Web Token
	authorized.Use(middleware.JWTAuth())
	{
		authorized.GET("/todos", handlers.GetTodos)
		authorized.GET("/todos/:id", handlers.GetTodoByID)

		// TODO: Use middleware for /todos/:id - for repetitive tasks like checking if the id is present and so on
		authorized.POST("/todos", handlers.AddTodo)
		authorized.POST("/todos/:id/share", handlers.ShareTodo)

		authorized.PUT("/todos/:id", handlers.UpdateTodo)
		authorized.PUT("/todos/:id/position", handlers.ChangeTodoPosition)
		authorized.PUT("/todos/:id/status", handlers.SetTodoStatus)

		authorized.DELETE("/todos/:id", handlers.DeleteTodo)
	}

	port := os.Getenv("PORT")
	address := fmt.Sprintf("localhost:%s", port)

	// Start the router
	err := router.Run(address)

	if err != nil {
		log.Fatal("Can't run awesome Todo API:", err)
	}
}
