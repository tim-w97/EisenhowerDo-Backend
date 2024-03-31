package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"github.com/tim-w97/my-awesome-Todo-API/util"
	"log"
	"net/http"
)

func getPositionFromTodo(transaction *sql.Tx, context *gin.Context) (int, error) {
	var currentPosition int

	sqlString, err := util.ReadSQLFile("get_todo_position.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		return 0, err
	}

	row := transaction.QueryRow(
		sqlString,
		context.GetInt("todoID"),
		context.GetInt("userID"),
	)

	scanErr := row.Scan(&currentPosition)

	if scanErr == nil {
		return currentPosition, nil
	}

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

func shiftOtherTodos(transaction *sql.Tx, currentPosition, desiredPosition int, context *gin.Context) error {
	var sqlFileName string

	if desiredPosition < currentPosition {
		sqlFileName = "incr_todo_positions.sql"
	} else {
		sqlFileName = "decr_todo_positions.sql"
	}

	sqlString, err := util.ReadSQLFile(sqlFileName)

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		return err
	}

	_, updateErr := transaction.Exec(
		sqlString,
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

func updatePosition(transaction *sql.Tx, desiredPosition int, context *gin.Context) error {
	sqlString, err := util.ReadSQLFile("update_todo_position.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		return err
	}

	_, updateErr := transaction.Exec(
		sqlString,
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

func doRollback(transaction *sql.Tx) {
	if err := transaction.Rollback(); err != nil {
		log.Print(err)
	}
}

func ChangeTodoPosition(context *gin.Context) {
	// Changing the position of a todos requires multiple operations
	// To maintain database integrity, I use a transaction here

	transaction, err := db.Database.Begin()

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't create a transaction"},
		)

		log.Print(err)
		return
	}

	currentPosition, err := getPositionFromTodo(transaction, context)

	if err != nil {
		log.Print(err)
		doRollback(transaction)
		return
	}

	desiredPosition, err := readDesiredPositionFromBody(context)

	if err != nil {
		log.Print(err)
		doRollback(transaction)
		return
	}

	if currentPosition == desiredPosition {
		context.IndentedJSON(
			http.StatusOK,
			gin.H{"message": "this todo is already at the desired position"},
		)

		doRollback(transaction)
		return
	}

	if err := shiftOtherTodos(transaction, currentPosition, desiredPosition, context); err != nil {
		log.Print(err)
		doRollback(transaction)
		return
	}

	if err := updatePosition(transaction, desiredPosition, context); err != nil {
		log.Print(err)
		doRollback(transaction)
		return
	}

	// Commit the transaction
	if err := transaction.Commit(); err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't commit transaction"},
		)

		log.Print(err)
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "successfully moved todo to desired position"},
	)
}
