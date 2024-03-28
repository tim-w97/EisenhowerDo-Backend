package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
)

func UpdateTodo(context *gin.Context) {
	var updatedTodo types.Todo

	// TODO: only update given values

	if bindErr := context.BindJSON(&updatedTodo); bindErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Can't convert body to Todo item"},
		)

		log.Print(bindErr)
		return
	}

	updatedTodo.ID = context.GetInt("todoID")
	updatedTodo.UserID = context.GetInt("userID")

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

	// TODO: Ensure if all three values can be updated
	result, updateErr := db.Database.Exec(
		"UPDATE todo SET title = ?, text = ?, isCompleted = ? WHERE id = ? AND userID = ?",
		updatedTodo.Title,
		updatedTodo.Text,
		updatedTodo.IsCompleted,
		updatedTodo.ID,
		updatedTodo.UserID,
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
			gin.H{"message": "there is no todo associated with this id or you aren't the creator of this todo"},
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
