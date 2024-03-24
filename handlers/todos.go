package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/data"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"net/http"
)

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, data.DummyTodos)
}

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

func GetTodoByID(context *gin.Context) {
	id := context.Param("id")

	for _, todo := range data.DummyTodos {
		if fmt.Sprint(todo.ID) == id {
			context.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	context.IndentedJSON(
		http.StatusNotFound,
		gin.H{"message": "todo not found"},
	)
}
