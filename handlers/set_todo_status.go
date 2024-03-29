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

	result, err := db.Database.Exec(
		"UPDATE todo SET isCompleted = ? WHERE id = ? AND userID = ?",
		context.GetInt("todoID"),
		context.GetInt("userID"),
	)

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "insert failed, can't update status of todo"},
		)

		log.Print(err)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't count affected rows"},
		)
	}

	if rowsAffected == 0 {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "updated no rows"},
		)

		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "updated status of todo successfully"},
	)
}
