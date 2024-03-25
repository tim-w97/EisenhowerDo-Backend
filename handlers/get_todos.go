package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
)

func GetTodos(context *gin.Context) {
	var todos []types.Todo

	rows, queryErr := db.Database.Query("SELECT * FROM todo")

	if queryErr != nil {
		log.Fatal("Can't query todos from database: ", queryErr)
	}

	for rows.Next() {
		var todo types.Todo

		if scanErr := rows.Scan(&todo.ID, &todo.Title, &todo.Text); scanErr != nil {
			log.Fatal("Can't assign todo row to todo struct: ", scanErr)
		}

		todos = append(todos, todo)
	}

	if closeErr := rows.Close(); closeErr != nil {
		log.Fatal("Can't close database todo rows: ", closeErr)
	}

	// Check for an error from the overall query
	if rowsErr := rows.Err(); rowsErr != nil {
		log.Fatal("The query for todo rows threw an error: ", rows)
	}

	context.IndentedJSON(http.StatusOK, todos)
}
