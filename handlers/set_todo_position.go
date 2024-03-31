package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"github.com/tim-w97/my-awesome-Todo-API/util"
	"log"
	"net/http"
)

func readDesiredPositionFromBody(context *gin.Context) (int, error) {
	var desiredPosition types.TodoPosition

	if bindErr := context.BindJSON(&desiredPosition); bindErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "can't read desired todo position from body"},
		)

		return 0, bindErr
	}

	return desiredPosition.Position, nil
}

func SetTodoPosition(context *gin.Context) {
	desiredPosition, err := readDesiredPositionFromBody(context)

	if err != nil {
		log.Print(err)
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

	_, updateErr := db.Database.Exec(
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

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "successfully moved todo to desired position"},
	)
}
