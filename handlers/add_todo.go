package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
)

func AddTodo(context *gin.Context) {
	var newTodo types.Todo

	// convert received json to a new Todo
	if bindErr := context.BindJSON(&newTodo); bindErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Can't convert body to Todo item"},
		)

		log.Print(bindErr)
		return
	}

	if len(newTodo.Title) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "please add a title"},
		)

		return
	}

	if len(newTodo.Text) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "please add a text"},
		)

		return
	}

	newTodo.UserID = context.GetInt("userID")

	// TODO: Add right position
	// TODO: Add Todo categories
	result, insertErr := db.Database.Exec(
		"INSERT INTO todo (title, text, userID) VALUES (?, ?, ?)",
		newTodo.Title,
		newTodo.Text,
		newTodo.UserID,
	)

	if insertErr != nil {
		log.Print("Can't insert todo: ", insertErr)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	insertedID, idErr := result.LastInsertId()

	if idErr != nil {
		log.Print("Can't get id of the inserted row: ", insertErr)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	newTodo.ID = int(insertedID)
	context.IndentedJSON(http.StatusCreated, newTodo)
}
