package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetHeaders(context *gin.Context) {
	context.Header(
		"Access-Control-Allow-Origin",
		//os.Getenv("WEBSITE"),
		"*",
	)

	context.Header(
		"Access-Control-Allow-Credentials",
		"true",
	)

	context.Header(
		"Content-Type",
		"application/json; charset=utf-8",
	)

	// Continue with the request if everything goes well
	context.Next()
}
