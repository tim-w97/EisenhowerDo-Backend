package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
)

func SetTodoStatus(context *gin.Context) {
	var todoStatus types.TodoStatus

	if bindErr := context.BindJSON(&todoStatus); bindErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "can't read todo status from body"},
		)

		log.Print(bindErr)
		return
	}

	_, updateErr := db.Database.Exec(
		"UPDATE todo SET isCompleted = ? WHERE id = ? AND userID = ?",
		todoStatus.IsCompleted,
		context.GetInt("todoID"),
		context.GetInt("userID"),
	)

	if updateErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "insert failed, can't update status of todo"},
		)

		log.Print(updateErr)
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "updated status of todo successfully"},
	)
}
