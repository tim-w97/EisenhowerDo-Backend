package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"github.com/tim-w97/my-awesome-Todo-API/util"
	"github.com/tim-w97/my-awesome-Todo-API/validation"
	"log"
	"net/http"
)

func readDesiredPositionFromBody(context *gin.Context) (position int, ok bool) {
	var desiredPosition types.TodoPosition

	if bindErr := context.BindJSON(&desiredPosition); bindErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "can't read desired todo position from body"},
		)

		log.Print(bindErr)
		ok = false
		return
	}

	ok = true
	position = desiredPosition.Position
	return
}

func SetTodoPosition(context *gin.Context) {
	desiredPosition, ok := readDesiredPositionFromBody(context)

	if !ok {
		return
	}

	if desiredPosition == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please provide your desired position"},
		)

		return
	}

	sql, err := util.ReadSQLFile("set_todo_position.sql")

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
		context.GetInt("todoID"),
		desiredPosition,
	)

	if updateErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't move todo to desired position"},
		)

		log.Print(updateErr)
		return
	}

	if ok := validation.ValidateSQLResult(result, context); !ok {
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "successfully moved todo to desired position"},
	)
}
