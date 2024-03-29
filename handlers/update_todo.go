package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
)

func UpdateTodo(context *gin.Context) {
	var updatedTodo types.Todo

	if bindErr := context.BindJSON(&updatedTodo); bindErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "can't convert body to todo"},
		)

		log.Print(bindErr)
		return
	}

	updatedTodo.ID = context.GetInt("todoID")
	updatedTodo.UserID = context.GetInt("userID")

	if len(updatedTodo.Title) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please add a title"},
		)

		return
	}

	if len(updatedTodo.Text) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please add a text"},
		)

		return
	}

	_, updateErr := db.Database.Exec(
		"UPDATE todo SET title = ?, text = ? WHERE id = ? AND userID = ?",
		updatedTodo.Title,
		updatedTodo.Text,
		updatedTodo.ID,
		updatedTodo.UserID,
	)

	if updateErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't update todo"},
		)

		log.Print(updateErr)
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "updated todo successfully"},
	)
}
