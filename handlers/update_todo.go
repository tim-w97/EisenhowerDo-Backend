package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/api"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"github.com/tim-w97/my-awesome-Todo-API/util"
	"log"
	"net/http"
)

func UpdateTodo(context *gin.Context) {
	var updatedTodo types.Todo

	if bindErr := context.BindJSON(&updatedTodo); bindErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "can't convert body to todo"},
		)

		log.Print(bindErr)
		return
	}

	updatedTodo.ID = context.GetInt("todoID")
	updatedTodo.UserID = context.GetInt("userID")

	if isValid := api.ValidateTodo(updatedTodo, context); !isValid {
		return
	}

	sql, err := util.ReadSQLFile("update_todo.sql")

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
		updatedTodo.Title,
		updatedTodo.Text,
		updatedTodo.CategoryID,
		updatedTodo.ID,
		updatedTodo.UserID,
	)

	if updateErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't update todo"},
		)

		log.Print(updateErr)
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "updated todo successfully"},
	)
}
