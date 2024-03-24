package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tim-w97/my-awesome-Todo-API/data"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"net/http"
	"time"
)

func Login(context *gin.Context) {
	var requestedUser types.User

	// convert received json to a user
	if err := context.BindJSON(&requestedUser); err != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Can't convert body to user"},
		)

		return
	}

	var foundUser types.User

	for _, dummyUser := range data.DummyUsers {
		if dummyUser.Username == requestedUser.Username {
			foundUser = dummyUser
			break
		}
	}

	if foundUser == (types.User{}) {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "This user doesn't exist"},
		)

		return
	}

	// TODO: Use password hashes
	if foundUser.Password != requestedUser.Password {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Your password is incorrect"},
		)

		return
	}

	// Generate the JSON Web Token and add the username to the claims
	// The token expires after 8 hours
	claims := jwt.MapClaims{
		"sub": foundUser.Username,
		"exp": time.Now().Add(time.Hour * 8).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"jwt": token},
	)
}
