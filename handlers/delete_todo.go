package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"log"
	"net/http"
)

func DeleteTodo(context *gin.Context) {
	userID := context.GetInt("userID")
	todoID := context.GetInt("todoID")

	result, deleteErr := db.Database.Exec(
		"DELETE FROM todo WHERE id = ? AND userID = ?",
		todoID,
		userID,
	)

	if deleteErr != nil {
		log.Print("Can't delete todo: ", deleteErr)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	affectedRows, affectedRowsErr := result.RowsAffected()

	if affectedRowsErr != nil {
		log.Print("Can't get number of deleted rows: ", affectedRowsErr)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if affectedRows == 0 {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "this todo doesn't exist or is already deleted"},
		)

		return
	}

	if affectedRows > 1 {
		log.Print("deleted more than one todo")
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": fmt.Sprintf("todo with id %d deleted", todoID)},
	)
}
