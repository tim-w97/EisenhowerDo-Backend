package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoRoute(context *gin.Context) {
	context.IndentedJSON(
		http.StatusNotFound,
		gin.H{"message": "this endpoint doesn't exist"},
	)
}
