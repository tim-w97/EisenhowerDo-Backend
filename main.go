package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	// routes
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)
	router.POST("/todos", addTodo)

	err := router.Run("localhost:8080")

	if err != nil {
		log.Fatal("Can't start the server:", err)
	}
}
