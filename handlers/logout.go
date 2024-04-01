package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(context *gin.Context) {
	// Tell the browser to delete the HTTP Only Cookie with the JWT token
	// by setting the value to an empty string and max age to a negative number
	context.SetCookie(
		"Authorization",
		"",
		-1,
		"",
		"",
		false,
		true,
	)

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "logout successful"},
	)
}
