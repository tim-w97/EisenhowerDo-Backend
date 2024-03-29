package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/handlers"
	"github.com/tim-w97/my-awesome-Todo-API/middleware"
)

func initLoginAndRegistration(router *gin.Engine) {
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
}

func initAuthorizedEndpoints(router *gin.Engine) {
	// all routes beginning with /todos are secured with JSON Web Token Authorization
	authorized := router.Group("/todos", middleware.JWTAuth)

	// all routes with a Todo ID parameter require parsing of the todo ID
	withTodoID := authorized.Group("/:id", middleware.ParseTodoID)

	// Get all Todos of a user
	authorized.GET("/", handlers.GetTodos)

	// Get a Todo by ID
	withTodoID.GET("/", handlers.GetTodoByID)

	// Add a new Todo
	authorized.POST("/", handlers.AddTodo)

	// Share a Todo with another user
	withTodoID.POST("/share", handlers.ShareTodo)

	// Update an existing Todo
	withTodoID.PUT("/", handlers.UpdateTodo)

	// Change the list position of a Todo
	withTodoID.PUT("/position", handlers.ChangeTodoPosition)

	// Toggle a Todo as completed or uncompleted
	withTodoID.PUT("/status", handlers.SetTodoStatus)

	// Delete a Todo
	withTodoID.DELETE("/", handlers.DeleteTodo)
}

func initEndpoints(router *gin.Engine) {
	initLoginAndRegistration(router)
	initAuthorizedEndpoints(router)
}
