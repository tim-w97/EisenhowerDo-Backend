package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/Todo24-API/db"
	"github.com/tim-w97/Todo24-API/util"
	"github.com/tim-w97/Todo24-API/validation"
	"log"
	"net/http"
)

func DeleteTodo(context *gin.Context) {
	userID := context.GetInt("userID")
	todoID := context.GetInt("todoID")

	sql, err := util.ReadSQLFile("delete_todo.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return
	}

	result, deleteErr := db.Database.Exec(
		sql,
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

	if ok := validation.ValidateSQLResult(result, context); !ok {
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "todo deleted successfully"},
	)
}
