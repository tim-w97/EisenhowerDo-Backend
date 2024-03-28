package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TodoIDExists(context *gin.Context) {
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
			gin.H{"message": "please send a numeric todo ID"},
		)

		context.Abort()
		return
	}

	context.Set("todoID", id)
}
