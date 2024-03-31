package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func ParseTodoID(context *gin.Context) {
	idString := context.Param("id")

	if idString == "" {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please provide a todo ID"},
		)

		context.Abort()
		return
	}

	id, convertErr := strconv.Atoi(idString)

	if convertErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please send a valid todo ID in integer format"},
		)

		log.Print(convertErr)
		context.Abort()
		return
	}

	context.Set("todoID", id)
}
