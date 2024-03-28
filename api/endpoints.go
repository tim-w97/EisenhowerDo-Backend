package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/handlers"
	"github.com/tim-w97/my-awesome-Todo-API/middleware"
)

func initEndpoints(router *gin.Engine) {
	// Registration and Login

	router.POST(
		"/register",
		handlers.Register,
	)

	router.POST(
		"/login",
		handlers.Login,
	)

	// Getting Todo items

	router.GET(
		"/todos",
		middleware.JWTAuth,
		handlers.GetTodos,
	)

	router.GET(
		"/todos/:id",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.GetTodoByID,
	)

	// Add and share Todo items

	router.POST(
		"/todos",
		middleware.JWTAuth,
		handlers.AddTodo,
	)

	router.POST(
		"/todos/:id/share",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.ShareTodo,
	)

	// Alter Todo items

	router.PUT(
		"/todos/:id",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.UpdateTodo,
	)

	router.PUT(
		"/todos/:id/position",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.ChangeTodoPosition,
	)

	router.PUT(
		"/todos/:id/status",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.SetTodoStatus,
	)

	// Delete Todo items

	router.DELETE(
		"/todos/:id",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.DeleteTodo,
	)
}
