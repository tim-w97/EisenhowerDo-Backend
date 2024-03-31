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

	if isValid := validation.ValidateTodo(newTodo, context); !isValid {
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

	result, insertErr := db.Database.Exec(
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

	if ok := validation.ValidateSQLResult(result, context); !ok {
		return
	}

	context.IndentedJSON(
		http.StatusCreated,
		gin.H{"message": "created todo successfully"},
	)
}
