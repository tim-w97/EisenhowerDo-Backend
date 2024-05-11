package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/Todo24-API/db"
	"github.com/tim-w97/Todo24-API/types"
	"github.com/tim-w97/Todo24-API/util"
	"log"
	"net/http"
)

func GetTodoByID(context *gin.Context) {
	var todo types.Todo

	userID := context.GetInt("userID")
	todoID := context.GetInt("todoID")

	sqlStr, err := util.ReadSQLFile("get_todo_by_id.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return
	}

	row := db.Database.QueryRow(
		sqlStr,
		todoID,
		userID,
	)

	scanErr := row.Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Title,
		&todo.Text,
		&todo.IsImportant,
		&todo.IsUrgent,
		&todo.CategoryID,
	)

	if scanErr == nil {
		context.IndentedJSON(http.StatusOK, todo)
		return
	}

	if errors.Is(scanErr, sql.ErrNoRows) {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "can't find a todo for this ID"},
		)

		return
	}

	context.IndentedJSON(
		http.StatusInternalServerError,
		gin.H{"message": "can't convert todo from database"},
	)

	log.Print(scanErr)
}
