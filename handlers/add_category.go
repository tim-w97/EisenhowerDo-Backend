package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"github.com/tim-w97/my-awesome-Todo-API/util"
	"github.com/tim-w97/my-awesome-Todo-API/validation"
	"log"
	"net/http"
)

func AddCategory(context *gin.Context) {
	var newCategory types.Category

	if bindErr := context.BindJSON(&newCategory); bindErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "can't read category from body"},
		)

		log.Print(bindErr)
		return
	}

	if newCategory.Category == "" {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "please add a valid category"},
		)

		return
	}

	sql, err := util.ReadSQLFile("add_category.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return
	}

	result, insertErr := db.Database.Exec(
		sql,
		newCategory.Category,
	)

	if insertErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't insert category"},
		)

		log.Print(insertErr)
		return
	}

	if ok := validation.ValidateSQLResult(result, context); !ok {
		return
	}

	context.IndentedJSON(
		http.StatusCreated,
		gin.H{"message": "created category successfully"},
	)
}
