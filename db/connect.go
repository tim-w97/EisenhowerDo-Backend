package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

// Database variable to access the database from other packages
var Database *sql.DB

func ConnectToDatabase() {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)

	var dbError error

	Database, dbError = sql.Open(
		"mysql",
		connectionString,
	)

	if dbError != nil {
		log.Fatal("Can't connect to the MySQL Database", dbError)
	}

	pingErr := Database.Ping()

	if pingErr != nil {
		log.Fatal("Can't ping the MySQL Database: ", pingErr)
	}
}
