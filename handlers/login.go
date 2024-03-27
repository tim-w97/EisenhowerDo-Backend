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

func searchUser(user types.User) (types.User, error) {
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
			return user, errors.New("incorrect username of password")
		}

		// TODO: This is a internal server error, not Bad Request
		log.Print("can't assign user row to user struct: ", scanErr)
		return user, errors.New("can't assign user row to user struct")
	}

	return queriedUser, nil
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

	// TODO: add expire time
	context.IndentedJSON(
		http.StatusOK,
		gin.H{"jwt": tokenString},
	)
}
