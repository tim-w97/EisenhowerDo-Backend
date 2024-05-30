package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/Todo24-API/db"
	"github.com/tim-w97/Todo24-API/types"
	"github.com/tim-w97/Todo24-API/util"
	"log"
	"net/http"
)

func GetSharedTodos(context *gin.Context) {
	sharedTodos := make([]types.Todo, 0)

	userID := context.GetInt("userID")

	sql, err := util.ReadSQLFile("get_shared_todos.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return
	}

	rows, queryErr := db.Database.Query(
		sql,
		userID,
	)

	if queryErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't query shared todos from database"},
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
			&todo.IsImportant,
			&todo.IsUrgent,
		); scanErr != nil {
			context.IndentedJSON(
				http.StatusInternalServerError,
				gin.H{"message": "can't assign todo row to todo struct"},
			)

			log.Print(scanErr)
			return
		}

		sharedTodos = append(sharedTodos, todo)
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
			gin.H{"message": "the query for shared todo rows threw an error"},
		)

		log.Print(rowsErr)
		return
	}

	context.IndentedJSON(http.StatusOK, sharedTodos)
}
