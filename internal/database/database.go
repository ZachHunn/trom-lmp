package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

/*
 Create env varibles
Create func db ip
*/

func ConnectDB() error {

	var err error
	connUri := "postgres://postgres:password@192.168.1.210:5432/training_platform_db?sslmode=disable"

	db, err := sql.Open("postgres", connUri)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	DB = db

	fmt.Println("Connected to database")

	return nil
}
