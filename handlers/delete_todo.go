package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"log"
	"net/http"
)

func DeleteTodo(context *gin.Context) {
	userID := context.GetInt("userID")
	todoID := context.GetInt("todoID")

	_, deleteErr := db.Database.Exec(
		"DELETE FROM todo WHERE id = ? AND userID = ?",
		todoID,
		userID,
	)

	if deleteErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't delete todo"},
		)

		log.Print(deleteErr)
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "todo deleted successfully"},
	)
}
