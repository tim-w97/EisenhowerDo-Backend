package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, dummyTodos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	// convert received json to a new Todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	dummyTodos = append(dummyTodos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(context *gin.Context) {
	id := context.Param("id")

	for _, todo := range dummyTodos {
		if todo.ID == id {
			context.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	context.IndentedJSON(
		http.StatusNotFound,
		gin.H{"message": "todo not found"},
	)
}
