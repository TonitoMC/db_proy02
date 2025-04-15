package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable"

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
