package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
)

func GetTodos(context *gin.Context) {
	// create an empty slice of todos
	todos := make([]types.Todo, 0)

	userID := context.GetInt("userID")

	rows, queryErr := db.Database.Query(
		"SELECT * FROM todo WHERE userID = ? ORDER BY position",
		userID,
	)

	if queryErr != nil {
		log.Print("Can't query todos from database: ", queryErr)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var todo types.Todo

		if scanErr := rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Title,
			&todo.Text,
			&todo.Position,
			&todo.IsCompleted,
		); scanErr != nil {
			log.Print("Can't assign todo row to todo struct: ", scanErr)
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		todos = append(todos, todo)
	}

	if closeErr := rows.Close(); closeErr != nil {
		log.Print("Can't close database todo rows: ", closeErr)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Check for an error from the overall query
	if rowsErr := rows.Err(); rowsErr != nil {
		log.Print("The query for todo rows threw an error: ", rowsErr)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.IndentedJSON(http.StatusOK, todos)
}
