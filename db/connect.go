package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

// Database variable to access the database from other packages
var Database *sql.DB

func ConnectToDatabase() {
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

	// I need to declare dbError outside to avoid a scope issue
	// on this way, the global Database variable gets assigned
	var dbError error
	Database, dbError = sql.Open("mysql", mySQLConfig.FormatDSN())

	if dbError != nil {
		log.Fatal("Can't connect to the MySQL Database")
	}

	pingErr := Database.Ping()

	if pingErr != nil {
		log.Fatal("Can't ping the MySQL Database: ", pingErr)
	}
}
