package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"github.com/tim-w97/my-awesome-Todo-API/util"
	"log"
	"net/http"
)

func validateTodo(todo types.Todo, context *gin.Context) (isValid bool) {
	if len(todo.Title) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please add a title"},
		)

		isValid = false
		return
	}

	if len(todo.Text) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please add a text"},
		)

		isValid = false
		return
	}

	if todo.CategoryID == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please add a category ID"},
		)

		isValid = false
		return
	}

	isValid = true
	return
}

func AddTodo(context *gin.Context) {
	var newTodo types.Todo

	// convert received json to a new Todo
	if bindErr := context.BindJSON(&newTodo); bindErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "can't read todo from body"},
		)

		log.Print(bindErr)
		return
	}

	if isValid := validateTodo(newTodo, context); !isValid {
		return
	}

	newTodo.UserID = context.GetInt("userID")

	sql, err := util.ReadSQLFile("add_todo.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return
	}

	_, insertErr := db.Database.Exec(
		sql,
		newTodo.UserID,
		newTodo.Title,
		newTodo.Text,
		newTodo.CategoryID,
	)

	if insertErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't insert todo"},
		)

		log.Print(insertErr)
		return
	}

	context.IndentedJSON(
		http.StatusCreated,
		gin.H{"message": "created todo successfully"},
	)
}
