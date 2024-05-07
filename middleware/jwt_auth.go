package middleware

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tim-w97/Todo24-API/db"
	"github.com/tim-w97/Todo24-API/types"
	"github.com/tim-w97/Todo24-API/util"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getUserByID(userID int, context *gin.Context) (user types.User, ok bool) {
	var queriedUser types.User

	sqlString, err := util.ReadSQLFile("get_user_by_id.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		context.Abort()
		log.Print(err)

		ok = false
		return
	}

	row := db.Database.QueryRow(
		sqlString,
		userID,
	)

	scanErr := row.Scan(&queriedUser.ID, &queriedUser.Username, &queriedUser.Password)

	if scanErr == nil {
		ok = true
		user = queriedUser
		return
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

	log.Print(scanErr)
	ok = false
	return
}

func getSecret(token *jwt.Token) (secret interface{}, error error) {
	// check if the signing method is the right one
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		error = errors.New("unexpected signing method")
		return
	}

	secret = []byte(os.Getenv("SECRET"))
	error = nil
	return
}

func getTokenFromHeader(header string) (string, error) {
	if header == "" {
		return "", errors.New("empty header")
	}

	jwtToken := strings.Split(header, " ")

	if len(jwtToken) != 2 {
		return "", errors.New("invalid header")
	}

	return jwtToken[1], nil
}

func JWTAuth(context *gin.Context) {
	tokenString, headerError := getTokenFromHeader(
		context.GetHeader("Authorization"),
	)

	if headerError != nil {
		context.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{"message": headerError.Error()},
		)

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

	user, ok := getUserByID(userID, context)

	if !ok {
		return
	}

	context.Set("userID", user.ID)

	// Continue with the request if everything goes well
	context.Next()
}
