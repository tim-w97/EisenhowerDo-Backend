package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
	"os"
	"time"
)

func searchUser(user types.User) (foundUser types.User, error error, httpStatusCode int) {
	// TODO (maybe): Check if the username exists

	var queriedUser types.User

	row := db.Database.QueryRow(
		"SELECT * FROM user WHERE username = ? AND password = ?",
		user.Username,
		user.Password,
	)

	scanErr := row.Scan(&queriedUser.ID, &queriedUser.Username, &queriedUser.Password)

	if scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			foundUser = user
			error = errors.New("incorrect username or password")
			httpStatusCode = http.StatusNotFound

			return
		}

		log.Print("can't assign user row to user struct: ", scanErr)

		foundUser = user
		error = errors.New("can't assign user row to user struct")
		httpStatusCode = http.StatusInternalServerError

		return
	}

	return queriedUser, nil, http.StatusOK
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

	if len(requestedUser.Username) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "please provide an username"},
		)

		return
	}

	if len(requestedUser.Password) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "please provide a password"},
		)

		return
	}

	// search user in database and return it if found
	user, searchError, httpStatus := searchUser(requestedUser)

	if searchError != nil {
		context.IndentedJSON(
			httpStatus,
			gin.H{"error": searchError.Error()},
		)

		return
	}

	// Generate the JSON Web Token and add the username to the claims
	// The token expires after 1 hour
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
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

	// Save the json web token as a http only cookie with a max age of 1 hour
	context.SetCookie(
		"Authorization",
		tokenString,
		3600,
		"",
		"",
		false,
		true,
	)

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "login successful"},
	)
}
