package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/data"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"net/http"
)

func AddTodo(context *gin.Context) {
	var newTodo types.Todo

	// convert received json to a new Todo
	if err := context.BindJSON(&newTodo); err != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Can't convert body to Todo item"},
		)

		return
	}

	data.DummyTodos = append(data.DummyTodos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}
