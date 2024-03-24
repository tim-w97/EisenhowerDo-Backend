package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
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

	// TODO: Use password hashes
	if foundUser.Password != requestedUser.Password {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Your password is incorrect"},
		)

		return
	}

	// Generate the JSON Web Token and add the username to the claims
	// The token expires after 8 hours
	claims := jwt.MapClaims{
		"sub": foundUser.Username,
		"exp": time.Now().Add(time.Hour * 8).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"jwt": token},
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
