package middleware

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
	"strconv"
)

func getUserByID(userID int, context *gin.Context) (types.User, error) {
	var queriedUser types.User

	sqlString, err := util.ReadSQLFile("get_user_by_id.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return types.User{}, err
	}

	row := db.Database.QueryRow(
		sqlString,
		userID,
	)

	scanErr := row.Scan(&queriedUser.ID, &queriedUser.Username, &queriedUser.Password)

	if scanErr == nil {
		return queriedUser, nil
	}

	if errors.Is(scanErr, sql.ErrNoRows) {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "user associated with token doesn't exist"},
		)
	} else {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't assign user row to user struct"},
		)
	}

	return types.User{}, scanErr
}

func getSecret(_ *jwt.Token) (interface{}, error) {
	// I could use the token here to check if the used algorithm is the one I expect
	secret := os.Getenv("SECRET")
	return []byte(secret), nil
}

func JWTAuth(context *gin.Context) {
	tokenString, cookieError := context.Cookie("Authorization")

	if cookieError != nil {
		context.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{"message": "can't get token from cookie"},
		)

		log.Print(cookieError)
		context.Abort()
		return
	}

	token, parseError := jwt.Parse(tokenString, getSecret)

	if errors.Is(parseError, jwt.ErrTokenExpired) {
		context.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{"message": "token is expired"},
		)

		context.Abort()
		return
	}

	if parseError != nil {
		context.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{"message": "can't parse token"},
		)

		log.Print(parseError)
		context.Abort()
		return
	}

	// Get the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		context.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{"message": "can't get claims from token"},
		)

		context.Abort()
		return
	}

	// the subject holds the id of the logged-in user
	subject, err := claims.GetSubject()

	if err != nil {
		context.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{"message": "can't get subject from token"},
		)

		context.Abort()
		return
	}

	userID, err := strconv.Atoi(subject)

	if err != nil {
		context.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{"message": "can't convert subject to user ID"},
		)

		context.Abort()
		return
	}

	user, searchError := getUserByID(userID, context)

	if searchError != nil {
		log.Print(searchError)
		context.Abort()
		return
	}

	context.Set("userID", user.ID)

	// Continue with the request if everything goes well
	context.Next()
}
