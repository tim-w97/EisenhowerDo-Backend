package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"log"
	"net/http"
	"strconv"
)

func DeleteTodo(context *gin.Context) {
	// TODO: User should only delete his todos, not others!

	idString := context.Param("id")

	if idString == "" {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please send a todo ID"},
		)

		return
	}

	id, convertErr := strconv.Atoi(idString)

	if convertErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please send a numeric todo ID (examples: 1, 2, 3, 42)"},
		)

		return
	}

	// TODO: duplicated code, use a middleware or a helper function
	user, userExists := context.Get("user")

	if !userExists {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "got no user to query todos"},
		)

		context.Abort()
		return
	}

	userID := user.(types.User).ID

	result, deleteErr := db.Database.Exec(
		"DELETE FROM todo WHERE id = ? AND userID = ?",
		id,
		userID,
	)

	if deleteErr != nil {
		log.Print("Can't delete todo: ", deleteErr)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	affectedRows, affectedRowsErr := result.RowsAffected()

	if affectedRowsErr != nil {
		log.Print("Can't get number of deleted rows: ", affectedRowsErr)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if affectedRows == 0 {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "this todo doesn't exist or is already deleted"},
		)

		return
	}

	if affectedRows > 1 {
		log.Print("deleted more than one todo")
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": fmt.Sprintf("todo with id %d deleted", id)},
	)
}
