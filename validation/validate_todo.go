package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/Todo24-API/types"
	"net/http"
)

func ValidateTodo(todo types.Todo, context *gin.Context) (isValid bool) {
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

	isValid = true
	return
}
