package validation

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ValidateSQLResult(result sql.Result, context *gin.Context) (ok bool) {
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't count affected rows"},
		)

		log.Print(err)
		return false
	}

	if rowsAffected == 0 {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "no changes were made"},
		)

		return false
	}

	return true
}
