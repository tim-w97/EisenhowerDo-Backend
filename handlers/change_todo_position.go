package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
)

func getPositionFromTodo(context *gin.Context) (int, error) {
	var currentPosition int

	row := db.Database.QueryRow(
		"SELECT position FROM todo WHERE id = ? AND userID = ?",
		context.GetInt("todoID"),
		context.GetInt("userID"),
	)

	if scanErr := row.Scan(&currentPosition); scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			context.IndentedJSON(
				http.StatusNotFound,
				gin.H{"message": "todo id doesn't exist or you aren't the creator"},
			)
		} else {
			context.IndentedJSON(
				http.StatusInternalServerError,
				gin.H{"message": "can't convert position of todo to move"},
			)
		}

		return 0, scanErr
	}

	return currentPosition, nil
}

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

func shiftOtherTodos(currentPosition, desiredPosition int, context *gin.Context) error {
	var updateOtherTodosSQL string

	if desiredPosition > currentPosition {
		// The user want's to move the Todo item "down",
		// so I have to decrement the positions of all todos between current position and desired position
		updateOtherTodosSQL =
			"UPDATE todo SET position = position - 1 WHERE position > ? AND position <= ? AND userID = ?"
	} else {
		// The user want's to move the Todo item "up",
		// so I have to increment the positions of all todos between current position and desired position
		updateOtherTodosSQL =
			"UPDATE todo SET position = position + 1 WHERE position < ? AND position >= ? AND userID = ?"
	}

	_, updateErr := db.Database.Exec(
		updateOtherTodosSQL,
		currentPosition,
		desiredPosition,
		context.GetInt("userID"),
	)

	if updateErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't shift other todos"},
		)

		return updateErr
	}

	return nil
}

func updatePosition(desiredPosition int, context *gin.Context) error {
	_, updateErr := db.Database.Exec(
		"UPDATE todo SET position = ? WHERE id = ? AND userID = ?",
		desiredPosition,
		context.GetInt("todoID"),
		context.GetInt("userID"),
	)

	if updateErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't update position of todo"},
		)

		return updateErr
	}

	return nil
}

func ChangeTodoPosition(context *gin.Context) {
	currentPosition, scanErr := getPositionFromTodo(context)

	if scanErr != nil {
		log.Print(scanErr)
		return
	}

	desiredPosition, bindErr := readDesiredPositionFromBody(context)

	if bindErr != nil {
		log.Print(bindErr)
		return
	}

	if currentPosition == desiredPosition {
		context.IndentedJSON(
			http.StatusOK,
			gin.H{"message": "this todo is already at the desired position"},
		)

		return
	}

	if err := shiftOtherTodos(currentPosition, desiredPosition, context); err != nil {
		log.Print(err)
		return
	}

	if err := updatePosition(desiredPosition, context); err != nil {
		log.Print(err)
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "successfully moved todo to desired position"},
	)
}
