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
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't query todos from database"},
		)

		log.Print(queryErr)
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
			context.IndentedJSON(
				http.StatusInternalServerError,
				gin.H{"message": "can't assign todo row to todo struct"},
			)

			log.Print(scanErr)
			return
		}

		todos = append(todos, todo)
	}

	if closeErr := rows.Close(); closeErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't close database todo rows"},
		)

		log.Print(closeErr)
		return
	}

	// Check for an error from the overall query
	if rowsErr := rows.Err(); rowsErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "the query for todo rows threw an error"},
		)

		log.Print(rowsErr)
		return
	}

	context.IndentedJSON(http.StatusOK, todos)
}
