package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/Todo24-API/handlers"
	"github.com/tim-w97/Todo24-API/middleware"
)

func initEndpoints(router *gin.Engine) {
	// Handle not existing routes
	router.NoRoute(handlers.NoRoute)

	// Handle not allowed methods
	router.NoMethod(handlers.NoMethod)

	// I don't group my routes because there is an issue with grouping routes and middleware
	// https://github.com/gin-gonic/gin/issues/531

	// Login and Registration

	router.POST(
		"/register",
		handlers.Register,
	)

	router.POST(
		"/login",
		handlers.Login,
	)

	router.POST(
		"/logout",
		handlers.Logout,
	)

	// Simple Todo Operations

	router.GET(
		"/todos",
		middleware.JWTAuth,
		handlers.GetTodos,
	)

	router.POST(
		"/todos",
		middleware.JWTAuth,
		handlers.AddTodo,
	)

	router.GET(
		"/todos/:id",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.GetTodoByID,
	)

	router.PUT(
		"/todos/:id",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.UpdateTodo,
	)

	router.DELETE(
		"/todos/:id",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.DeleteTodo,
	)

	// Advanced Todo Operations

	router.PUT(
		"/todos/:id/status",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.SetTodoStatus,
	)

	router.POST(
		"/todos/:id/share",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.ShareTodo,
	)

	router.PUT(
		"/todos/:id/position",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.SetTodoPosition,
	)

	// Shared Todos

	router.GET(
		"/todos/shared",
		middleware.JWTAuth,
		handlers.GetSharedTodos,
	)

	// Todo Categories

	router.GET(
		"/categories",
		middleware.JWTAuth,
		handlers.GetCategories,
	)

	router.POST(
		"/categories",
		middleware.JWTAuth,
		handlers.AddCategory,
	)

	// Don't allow deleting categories for now
	// because they are public and usable for every user
}
