package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"github.com/tim-w97/my-awesome-Todo-API/util"
	"log"
	"net/http"
	"os"
	"time"
)

func searchUser(user types.User, context *gin.Context) (types.User, error) {
	var queriedUser types.User

	passwordHash := util.GetPasswordHash(user.Password)

	sqlString, err := util.ReadSQLFile("login_user.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return user, err
	}

	row := db.Database.QueryRow(
		sqlString,
		user.Username,
		passwordHash,
	)

	scanErr := row.Scan(
		&queriedUser.ID,
		&queriedUser.Username,
		&queriedUser.Password,
	)

	if scanErr == nil {
		return queriedUser, nil
	}

	if errors.Is(scanErr, sql.ErrNoRows) {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "incorrect username or password"},
		)
	} else {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't assign user row to user struct"},
		)
	}

	log.Print(scanErr)
	return user, scanErr
}

func Login(context *gin.Context) {
	var requestedUser types.User

	// convert received json to a user
	if err := context.BindJSON(&requestedUser); err != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "can't convert body to user"},
		)

		log.Print(err)
		return
	}

	if len(requestedUser.Username) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please provide an username"},
		)

		return
	}

	if len(requestedUser.Password) == 0 {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please provide a password"},
		)

		return
	}

	// search user in database and return it if found
	user, searchError := searchUser(requestedUser, context)

	if searchError != nil {
		log.Print(searchError)
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
			gin.H{"message": "can't generate token"},
		)

		log.Print(signError)
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
