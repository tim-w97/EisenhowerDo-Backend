package main

import (
	"github.com/joho/godotenv"
	"github.com/tim-w97/Todo24-API/api"
	"github.com/tim-w97/Todo24-API/db"
	"log"
)

// TODO: add documentation to all functions
func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can't load the .env file: ", err)
	}

	// Create connection to MySQL Database
	db.ConnectToDatabase()

	// Let's run this thing!
	api.InitEndpointsAndRun()
}
