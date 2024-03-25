package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

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

	db, dbError := sql.Open("mysql", mySQLConfig.FormatDSN())

	if dbError != nil {
		log.Fatal("Can't connect to the MySQL Database")
	}

	// Test the database connection
	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal("Can't ping the MySQL Database: ", pingErr)
	}
}
