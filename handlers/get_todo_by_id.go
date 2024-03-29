package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
)

func GetTodoByID(context *gin.Context) {
	var todo types.Todo

	userID := context.GetInt("userID")
	todoID := context.GetInt("todoID")

	row := db.Database.QueryRow(
		"SELECT * FROM todo WHERE id = ? AND userID = ?",
		todoID,
		userID,
	)

	scanErr := row.Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Title,
		&todo.Text,
		&todo.Position,
		&todo.IsCompleted,
	)

	if scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			context.IndentedJSON(
				http.StatusNotFound,
				gin.H{"message": "Can't find a todo for this ID"},
			)

			return
		}

		log.Print("Can't assign todo row to todo struct: ", scanErr)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}
