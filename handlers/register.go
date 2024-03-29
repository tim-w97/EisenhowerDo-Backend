package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"github.com/tim-w97/my-awesome-Todo-API/util"
	"log"
	"net/http"
)

func searchUsername(username string) (usernameExists bool, error error) {
	var usernameCount int

	row := db.Database.QueryRow(
		"SELECT COUNT(*) FROM user WHERE username = ?",
		username,
	)

	if scanErr := row.Scan(&usernameCount); scanErr != nil {
		error = scanErr
		return
	}

	usernameExists = usernameCount > 0
	error = nil
	return
}

func Register(context *gin.Context) {
	var userToRegister types.User

	// convert received json to a user
	if err := context.BindJSON(&userToRegister); err != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "Can't convert body to user"},
		)

		return
	}

	usernameExists, usernameSearchErr := searchUsername(userToRegister.Username)

	if usernameSearchErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't determine if username already exists"},
		)

		log.Print(usernameSearchErr)
		return
	}

	if usernameExists {
		context.IndentedJSON(
			http.StatusConflict,
			gin.H{"message": "this username is already taken"},
		)

		return
	}

	passwordHash := util.GetPasswordHash(userToRegister.Password)

	result, insertErr := db.Database.Exec(
		"INSERT INTO user (username, password) VALUES (?, ?)",
		userToRegister.Username,
		passwordHash,
	)

	// TODO: Ensure to handle all errors like this
	if insertErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't add new user to the database"},
		)

		log.Print("Can't insert user: ", insertErr)
		return
	}

	insertedID, idErr := result.LastInsertId()

	if idErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't get id of created user"},
		)

		log.Print("Can't get id of the inserted row: ", insertErr)
		return
	}

	userToRegister.ID = int(insertedID)

	context.IndentedJSON(
		http.StatusCreated,
		gin.H{"message": "user registered successfully"},
	)
}
