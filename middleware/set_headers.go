package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func SetHeaders(context *gin.Context) {
	context.Header(
		"Access-Control-Allow-Origin",
		os.Getenv("WEBSITE"),
	)

	context.Header(
		"Access-Control-Allow-Methods",
		"GET, POST, PUT, DELETE, OPTIONS",
	)

	context.Header(
		"Access-Control-Allow-Credentials",
		"true",
	)

	// Handle preflight requests
	if context.Request.Method == "OPTIONS" {
		context.Status(http.StatusNoContent)
	}

	// Continue with the request if everything goes well
	context.Next()
}
