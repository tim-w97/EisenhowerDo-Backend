package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/data"
	"net/http"
)

func GetTodoByID(context *gin.Context) {
	id := context.Param("id")

	for _, todo := range data.DummyTodos {
		if fmt.Sprint(todo.ID) == id {
			context.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	context.IndentedJSON(
		http.StatusNotFound,
		gin.H{"message": "todo not found"},
	)
}
