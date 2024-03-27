package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
	"strconv"
)

func UpdateTodo(context *gin.Context) {
	var updatedTodo types.Todo

	// TODO: only update given values

	idString := context.Param("id")

	if idString == "" {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please send a todo ID"},
		)

		return
	}

	id, convertErr := strconv.Atoi(idString)

	if convertErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please send a numeric todo ID (examples: 1, 2, 3, 42)"},
		)

		return
	}

	if bindErr := context.BindJSON(&updatedTodo); bindErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Can't convert body to Todo item"},
		)

		log.Print(bindErr)
		return
	}

	updatedTodo.ID = id

	if len(updatedTodo.Title) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "please add a title"},
		)

		return
	}

	if len(updatedTodo.Text) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "please add a text"},
		)

		return
	}

	result, updateErr := db.Database.Exec(
		"UPDATE todo SET title = ?, text = ? WHERE id = ?",
		updatedTodo.Title,
		updatedTodo.Text,
		updatedTodo.ID,
	)

	if updateErr != nil {
		log.Print("Can't update todo: ", updateErr)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	affectedRows, affectedRowsErr := result.RowsAffected()

	if affectedRowsErr != nil {
		log.Print("Can't get number of updated rows: ", affectedRowsErr)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if affectedRows == 0 {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("no update happened, there is no todo with id %d", affectedRows)},
		)

		return
	}

	if affectedRows > 1 {
		log.Print("updated more than one todo")
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.IndentedJSON(http.StatusOK, updatedTodo)
}
