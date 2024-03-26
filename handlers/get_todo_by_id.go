package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
	"strconv"
)

func GetTodoByID(context *gin.Context) {
	var todo types.Todo

	idString := context.Param("id")

	if idString == "" {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please send a todo ID"},
		)

		return
	}

	id, convertErr := strconv.Atoi(idString)

	if convertErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please send a numeric todo ID (examples: 1, 2, 3, 42)"},
		)

		return
	}

	row := db.Database.QueryRow("SELECT * FROM todo WHERE id = ?", id)

	scanErr := row.Scan(&todo.ID, &todo.Title, &todo.Text)

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
