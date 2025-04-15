package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load environment variables")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging DB: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")

	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal("AAAAAAAAAAA")
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string

		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal("A")
		}

		fmt.Printf("ID %d Name %s Email %s\n", id, name, email)
	}
}
