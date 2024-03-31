package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"github.com/tim-w97/my-awesome-Todo-API/util"
	"log"
	"net/http"
)

func searchUsername(username string, context *gin.Context) (bool, error) {
	var usernameCount int

	sql, err := util.ReadSQLFile("count_username.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return false, err
	}

	row := db.Database.QueryRow(
		sql,
		username,
	)

	if scanErr := row.Scan(&usernameCount); scanErr != nil {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "can't convert query result to int"},
		)

		return false, scanErr
	}

	return usernameCount > 0, nil
}

func Register(context *gin.Context) {
	var userToRegister types.User

	// convert received json to a user
	if err := context.BindJSON(&userToRegister); err != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "can't convert body to user"},
		)

		log.Print(err)
		return
	}

	usernameIsTaken, err := searchUsername(
		userToRegister.Username,
		context,
	)

	if err != nil {
		log.Print(err)
		return
	}

	if usernameIsTaken {
		context.IndentedJSON(
			http.StatusConflict,
			gin.H{"message": "this username is already taken"},
		)

		return
	}

	passwordHash := util.GetPasswordHash(userToRegister.Password)

	sql, err := util.ReadSQLFile("create_user.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return
	}

	_, insertErr := db.Database.Exec(
		sql,
		userToRegister.Username,
		passwordHash,
	)

	if insertErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't add new user to the database"},
		)

		log.Print(insertErr)
		return
	}

	context.IndentedJSON(
		http.StatusCreated,
		gin.H{"message": "user registered successfully"},
	)
}
