package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/Todo24-API/db"
	"github.com/tim-w97/Todo24-API/types"
	"github.com/tim-w97/Todo24-API/util"
	"log"
	"net/http"
)

func GetCategories(context *gin.Context) {
	categories := make([]types.Category, 0)

	sql, err := util.ReadSQLFile("get_categories.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return
	}

	rows, queryErr := db.Database.Query(sql)

	if queryErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't query categories from database"},
		)

		log.Print(queryErr)
		return
	}

	for rows.Next() {
		var category types.Category

		if scanErr := rows.Scan(
			&category.ID,
			&category.Category,
		); scanErr != nil {
			context.IndentedJSON(
				http.StatusInternalServerError,
				gin.H{"message": "can't assign category row to category struct"},
			)

			log.Print(scanErr)
			return
		}

		categories = append(categories, category)
	}

	if closeErr := rows.Close(); closeErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't close database category rows"},
		)

		log.Print(closeErr)
		return
	}

	// Check for an error from the overall query
	if rowsErr := rows.Err(); rowsErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "the query for category rows threw an error"},
		)

		log.Print(rowsErr)
		return
	}

	context.IndentedJSON(http.StatusOK, categories)
}
