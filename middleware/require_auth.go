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

func JWTAuth() gin.HandlerFunc {
	return requireAuth
}

// TODO: This is duplicated code, see login.go
func getUserByID(userID int) (types.User, error) {
	var queriedUser types.User

	row := db.Database.QueryRow(
		"SELECT * FROM user WHERE id = ?",
		userID,
	)

	scanErr := row.Scan(&queriedUser.ID, &queriedUser.Username, &queriedUser.Password)

	if scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			return types.User{}, errors.New("there is no user with this ID")
		}

		// TODO: This is a internal server error, not Bad Request
		log.Print("can't assign user row to user struct: ", scanErr)
		return types.User{}, errors.New("can't assign user row to user struct")
	}

	return queriedUser, nil
}

func getSecret(_ *jwt.Token) (interface{}, error) {
	// I could use the token here to check if the used algorithm is the one I expect,
	// but I skip it for simplicity now

	secret := os.Getenv("SECRET")
	return []byte(secret), nil
}

func requireAuth(context *gin.Context) {
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

	// TODO: Why do I get this ID as float64? Find out!
	// Get the users ID from the token and search the corresponding user
	userID := int(claims["sub"].(float64))

	user, searchError := getUserByID(userID)

	if searchError != nil {
		log.Print("Can't find user from JWT token: ", searchError)
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	context.Set("user", user)

	// Continue with the request if everything goes well
	context.Next()
}
