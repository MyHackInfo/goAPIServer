package sqlDbConnection

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Initialize and open the database connection
func initDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/godb?parseTime=true")
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// SQLDBInit initializes the database connection and returns it
func SQLDBInit() *sql.DB {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
