package middleware

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

func getUserByID(userID int) (foundUser types.User, error error, httpStatusCode int) {
	var queriedUser types.User

	row := db.Database.QueryRow(
		"SELECT * FROM user WHERE id = ?",
		userID,
	)

	scanErr := row.Scan(&queriedUser.ID, &queriedUser.Username, &queriedUser.Password)

	if scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			error = errors.New("there is no user with this ID")
			httpStatusCode = http.StatusNotFound

			return
		}

		error = errors.New("can't assign user row to user struct")
		httpStatusCode = http.StatusInternalServerError

		return
	}

	return queriedUser, nil, http.StatusOK
}

func getSecret(_ *jwt.Token) (interface{}, error) {
	// I could use the token here to check if the used algorithm is the one I expect
	secret := os.Getenv("SECRET")
	return []byte(secret), nil
}

func JWTAuth(context *gin.Context) {
	tokenString, cookieError := context.Cookie("Authorization")

	if cookieError != nil {
		log.Print("Can't get cookie with JWT token: ", cookieError)
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, parseError := jwt.Parse(tokenString, getSecret)

	if parseError != nil {
		log.Print("Can't parse JWT token: ", parseError)
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		log.Print("JWT token is invalid")
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Get the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Print("Can't get claims from JWT token")
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check if the token is expired
	currentTime := float64(time.Now().Unix())
	tokenExpireTime := claims["exp"].(float64)

	if currentTime > tokenExpireTime {
		log.Print("JWT token is expired")
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Get the user id from the tokens subject and search the corresponding user
	// JSON package parses numbers as float64, so we need to convert the subject back to int
	subject := claims["sub"].(float64)
	userID := int(subject)

	user, searchError, httpStatusCode := getUserByID(userID)

	if searchError != nil {
		log.Print("Can't find user from JWT token: ", searchError.Error())
		context.AbortWithStatus(httpStatusCode)
		return
	}

	context.Set("userID", user.ID)

	// Continue with the request if everything goes well
	context.Next()
}
