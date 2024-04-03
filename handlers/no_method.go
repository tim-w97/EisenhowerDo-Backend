package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoMethod(context *gin.Context) {
	context.IndentedJSON(
		http.StatusMethodNotAllowed,
		gin.H{"message": "this method is not allowed"},
	)
}
