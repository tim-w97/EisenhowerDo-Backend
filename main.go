package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/tim-w97/my-awesome-Todo-API/handlers"
	"github.com/tim-w97/my-awesome-Todo-API/middleware"
	"log"
	"os"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can't load the .env file: ", err)
	}

	// Create connection to MySQL Database
	mySQLHost := os.Getenv("MYSQL_HOST")
	mySQLPort := os.Getenv("MYSQL_PORT")
	mySQLAddress := fmt.Sprintf("%s:%s", mySQLHost, mySQLPort)

	mySQLConfig := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASS"),
		Net:                  "tcp",
		Addr:                 mySQLAddress,
		DBName:               os.Getenv("MYSQL_DB"),
		AllowNativePasswords: true,
	}

	db, dbError := sql.Open("mysql", mySQLConfig.FormatDSN())

	if dbError != nil {
		log.Fatal("Can't connect to the MySQL Database")
	}

	// Test the database connection
	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal("Can't ping the MySQL Database: ", pingErr)
	}

	// Set up routes
	router := gin.Default()

	// routes for registration and login
	router.POST("/login", handlers.Login)

	// routes for handling Todo items
	router.GET("/todos", middleware.RequireAuth, handlers.GetTodos)
	router.GET("/todos/:id", middleware.RequireAuth, handlers.GetTodoByID)
	router.POST("/todos", middleware.RequireAuth, handlers.AddTodo)

	// run the awesome Todo API
	port := os.Getenv("PORT")
	address := fmt.Sprintf("localhost:%s", port)

	if err := router.Run(address); err != nil {
		log.Fatal("Can't run awesome Todo API:", err)
	}
}
