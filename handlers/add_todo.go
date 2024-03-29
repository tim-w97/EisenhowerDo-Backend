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
			gin.H{"message": "can't read todo from body"},
		)

		log.Print(bindErr)
		return
	}

	if len(newTodo.Title) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please add a title"},
		)

		return
	}

	if len(newTodo.Text) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please add a text"},
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
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't insert todo"},
		)

		log.Print(insertErr)
		return
	}

	insertedID, idErr := result.LastInsertId()

	if idErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't get id of inserted todo"},
		)

		log.Print(idErr)
		return
	}

	newTodo.ID = int(insertedID)

	context.IndentedJSON(http.StatusCreated, newTodo)
}
