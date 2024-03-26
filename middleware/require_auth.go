package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tim-w97/my-awesome-Todo-API/data"
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
func getUser(username string) (types.User, error) {
	for _, dummyUser := range data.DummyUsers {
		if dummyUser.Username == username {
			return dummyUser, nil
		}
	}

	return types.User{}, errors.New("this user doesn't exist")
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

	// Get the username from the token and search the corresponding user
	username := claims["sub"].(string)

	user, searchError := getUser(username)

	if searchError != nil {
		log.Print("Can't find user from JWT token: ", searchError)
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	context.Set("user", user)

	// Continue with the request if everything goes well
	context.Next()
}
