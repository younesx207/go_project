package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver for database/sql, using blank identifier as it's only used for side-effects
)

// DB is a global variable that will hold the database connection pool
var DB *sql.DB

// Connect establishes a connection to the MySQL database and assigns it to the global DB variable
func Connect() error {
	// Data Source Name (DSN) containing the connection info for the MySQL database
	dsn := "root:root@tcp(localhost:3306)/bookstores"

	// Open a new database connection using the provided DSN
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		// Return the error if the connection could not be established
		return err
	}

	// Ping the database to verify the connection is alive
	if err := db.Ping(); err != nil {
		// Return the error if the ping failed, indicating the connection isn't working
		return err
	}

	// Assign the established database connection to the global DB variable
	DB = db

	// Print a success message indicating the connection was successful
	fmt.Println("Connected to the database successfully")

	// Return nil indicating no errors occurred
	return nil
}
