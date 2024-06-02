package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/Todo24-API/db"
	"github.com/tim-w97/Todo24-API/types"
	"github.com/tim-w97/Todo24-API/util"
	"github.com/tim-w97/Todo24-API/validation"
	"log"
	"net/http"
)

func getUserID(context *gin.Context, username string) (userID int, err error) {
	sqlString, readSqlErr := util.ReadSQLFile("get_user_by_username.sql")

	if readSqlErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		err = readSqlErr
		return
	}

	var user types.User

	row := db.Database.QueryRow(
		sqlString,
		username,
	)

	scanErr := row.Scan(&user.ID, &user.Username, &user.Password)

	if scanErr == nil {
		userID = user.ID
		return
	}

	err = scanErr

	if errors.Is(scanErr, sql.ErrNoRows) {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "Diesen Benutzer gibt es nicht"},
		)
	} else {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't assign user row to user struct"},
		)
	}

	return
}

func ShareTodo(context *gin.Context) {
	var sharedTodo types.SharedTodo

	if bindErr := context.BindJSON(&sharedTodo); bindErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "can't convert body to share todo"},
		)

		log.Print(bindErr)
		return
	}

	userID, err := getUserID(context, sharedTodo.Username)

	if err != nil {
		log.Print(err)
		return
	}

	if context.GetInt("userID") == userID {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "you can't share a todo with yourself"},
		)

		return
	}

	sqlString, err := util.ReadSQLFile("share_todo.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return
	}

	result, insertErr := db.Database.Exec(
		sqlString,
		context.GetInt("todoID"),
		userID,
	)

	if insertErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't insert shared todo"},
		)

		log.Print(insertErr)
		return
	}

	if ok := validation.ValidateSQLResult(result, context); !ok {
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "shared todo successfully"},
	)
}
