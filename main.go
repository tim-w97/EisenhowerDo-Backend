package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Structure of a Todo item
type todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Some dummy Todo items
var dummyTodos = []todo{
	{ID: "1", Title: "Einkaufen", Text: "Ich brauch noch Toastbrot und Nutella"},
	{ID: "2", Title: "Geschenk für Oma kaufen", Text: "Ideen: Orchidee, Pralinen, Käsekuchen"},
	{ID: "3", Title: "Putzen", Text: "Staubsaugen, Kleiderschrank ausmisten, Schuhe putzen"},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)

	err := router.Run("localhost:8080")

	if err != nil {
		log.Fatal("Can't start the server:", err)
	}
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, dummyTodos)
}
