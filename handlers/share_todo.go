package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
)

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

	sharedTodo.TodoID = context.GetInt("todoID")

	result, insertErr := db.Database.Exec(
		"INSERT INTO sharedTodo (todoID, otherUserID) VALUES (?, ?)",
		sharedTodo.TodoID,
		sharedTodo.OtherUserID,
	)

	if insertErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't insert shared todo"},
		)

		log.Print(insertErr)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't count affected rows"},
		)

		log.Print(err)
		return
	}

	if rowsAffected == 0 {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "updated no rows"},
		)

		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "shared todo successfully"},
	)
}
