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

func getUserByID(userID int, context *gin.Context) (types.User, error) {
	var queriedUser types.User

	row := db.Database.QueryRow(
		"SELECT * FROM user WHERE id = ?",
		userID,
	)

	scanErr := row.Scan(&queriedUser.ID, &queriedUser.Username, &queriedUser.Password)

	if scanErr == nil {
		return queriedUser, nil
	}

	if errors.Is(scanErr, sql.ErrNoRows) {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "there is no user with this ID"},
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

		log.Print(cookieError.Error())
		context.Abort()
		return
	}

	token, parseError := jwt.Parse(tokenString, getSecret)

	if parseError != nil {
		context.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{"message": "can't parse token"},
		)

		log.Print(parseError.Error())
		context.Abort()
		return
	}

	if !token.Valid {
		context.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{"message": "token is invalid"},
		)

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

	// TODO: check if that works correctly!
	// Check if the token is expired
	currentTime := float64(time.Now().Unix())
	tokenExpireTime := claims["exp"].(float64)

	if currentTime > tokenExpireTime {
		context.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{"message": "token is expired"},
		)

		context.Abort()
		return
	}

	// Get the user id from the tokens subject and search the corresponding user
	// JSON package parses numbers as float64, so we need to convert the subject back to int
	subject := claims["sub"].(float64)
	userID := int(subject)

	user, searchError := getUserByID(userID, context)

	if searchError != nil {
		// TODO: check if I call .Error() everywhere
		log.Print(searchError.Error())
		context.Abort()
		return
	}

	context.Set("userID", user.ID)

	// Continue with the request if everything goes well
	context.Next()
}
