package main

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
