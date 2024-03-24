package main

type todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
