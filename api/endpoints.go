package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/handlers"
	"github.com/tim-w97/my-awesome-Todo-API/middleware"
)

func initEndpoints(router *gin.Engine) {
	// I don't use grouping here because there is an issue with grouping routes and middleware
	// https://github.com/gin-gonic/gin/issues/531

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

	router.PUT(
		"/todos/:id/status",
		middleware.JWTAuth,
		middleware.ParseTodoID,
		handlers.SetTodoStatus,
	)

}
