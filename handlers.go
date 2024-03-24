package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(context *gin.Context) {
	var requestedUser user

	// convert received json to a user
	if err := context.BindJSON(&requestedUser); err != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Can't convert body to user"},
		)

		return
	}

	var foundUser user

	for _, dummyUser := range dummyUsers {
		if dummyUser.Username == requestedUser.Username {
			foundUser = dummyUser
			break
		}
	}

	if foundUser == (user{}) {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "This user doesn't exist"},
		)

		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "User exists"},
	)
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, dummyTodos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	// convert received json to a new Todo
	if err := context.BindJSON(&newTodo); err != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Can't convert body to Todo item"},
		)

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
