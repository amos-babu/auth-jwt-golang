package database

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB(dbUsername, dbName, dbPassword string) (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUsername, dbName, url.QueryEscape(dbPassword))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createUserDB(db); err != nil {
		return nil, err
	}
	fmt.Println("Connected to DB succesfully")

	return db, nil

}

func createUserDB(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL UNIQUE,
            email VARCHAR(255) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("Error creating user table: %w", err)
	}

	fmt.Println("Table Created Succesfully")

	return nil
}
