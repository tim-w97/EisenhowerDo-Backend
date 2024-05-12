package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/Todo24-API/db"
	"github.com/tim-w97/Todo24-API/types"
	"github.com/tim-w97/Todo24-API/util"
	"log"
	"net/http"
)

func searchUser(user types.User, context *gin.Context) (types.User, bool) {
	var queriedUser types.User

	passwordHash := util.GetPasswordHash(user.Password)

	sqlString, err := util.ReadSQLFile("login_user.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return user, false
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
		return queriedUser, true
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
	return user, false
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
	user, ok := searchUser(requestedUser, context)

	if !ok {
		return
	}

	token, signError := util.GenerateToken(user.ID)

	if signError != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't generate token"},
		)

		log.Print(signError)
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"token": token},
	)
}
