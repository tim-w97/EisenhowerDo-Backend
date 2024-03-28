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
		handlers.GetTodoByID,
	)

	// Add and share Todo items

	// TODO: Use middleware for /todos/:id - for repetitive tasks like checking if the id is present and so on
	router.POST(
		"/todos",
		middleware.JWTAuth,
		handlers.AddTodo,
	)

	router.POST(
		"/todos/:id/share",
		middleware.JWTAuth,
		middleware.TodoIDExists,
		handlers.ShareTodo,
	)

	// Alter Todo items

	router.PUT(
		"/todos/:id",
		middleware.JWTAuth,
		middleware.TodoIDExists,
		handlers.UpdateTodo,
	)

	router.PUT(
		"/todos/:id/position",
		middleware.JWTAuth,
		middleware.TodoIDExists,
		handlers.ChangeTodoPosition,
	)

	router.PUT(
		"/todos/:id/status",
		middleware.JWTAuth,
		middleware.TodoIDExists,
		handlers.SetTodoStatus,
	)

	// Delete Todo items

	router.DELETE(
		"/todos/:id",
		middleware.JWTAuth,
		middleware.TodoIDExists,
		handlers.DeleteTodo,
	)
}
