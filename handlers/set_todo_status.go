package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/Todo24-API/db"
	"github.com/tim-w97/Todo24-API/types"
	"github.com/tim-w97/Todo24-API/util"
	"github.com/tim-w97/Todo24-API/validation"
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

	sql, err := util.ReadSQLFile("set_todo_status.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return
	}

	result, updateErr := db.Database.Exec(
		sql,
		todoStatus.IsCompleted,
		context.GetInt("todoID"),
		context.GetInt("userID"),
	)

	if updateErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "insert failed, can't update status of todo"},
		)

		log.Print(updateErr)
		return
	}

	if ok := validation.ValidateSQLResult(result, context); !ok {
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "updated status of todo successfully"},
	)
}
