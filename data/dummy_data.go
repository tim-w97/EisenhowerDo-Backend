package data

import "github.com/tim-w97/my-awesome-Todo-API/types"

var DummyTodos = []types.Todo{
	{ID: "1", Title: "Einkaufen", Text: "Ich brauch noch Toastbrot und Nutella"},
	{ID: "2", Title: "Geschenk für Oma kaufen", Text: "Ideen: Orchidee, Pralinen, Käsekuchen"},
	{ID: "3", Title: "Putzen", Text: "Staubsaugen, Kleiderschrank ausmisten, Schuhe putzen"},
}

var DummyUsers = []types.User{
	{Username: "tim", Password: "test"},
	{Username: "celine", Password: "mcdonalds"},
	{Username: "ronny", Password: "benz"},
}
