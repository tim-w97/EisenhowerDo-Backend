package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tim-w97/my-awesome-Todo-API/data"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"net/http"
	"os"
	"time"
)

func searchUser(user types.User) (types.User, error) {
	for _, dummyUser := range data.DummyUsers {
		if dummyUser.Username == user.Username {
			return dummyUser, nil
		}
	}

	return user, errors.New("this user doesn't exist")
}

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

	// search user in database and return it if found
	user, searchError := searchUser(requestedUser)

	if searchError != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": searchError.Error()},
		)

		return
	}

	// TODO: Use password hashes
	if user.Password != requestedUser.Password {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Your password is incorrect"},
		)

		return
	}

	// Generate the JSON Web Token and add the username to the claims
	// The token expires after 8 hours
	claims := jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(time.Hour * 8).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("SECRET"))

	tokenString, signError := token.SignedString(secret)

	if signError != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Can't generate token"},
		)

		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"jwt": tokenString},
	)
}
