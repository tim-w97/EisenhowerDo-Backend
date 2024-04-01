package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
)

func SetHeaders(context *gin.Context) {
	context.Header(
		"Access-Control-Allow-Origin",
		os.Getenv("WEBSITE"),
	)

	context.Header(
		"Access-Control-Allow-Credentials",
		"true",
	)

	// Continue with the request if everything goes well
	context.Next()
}
