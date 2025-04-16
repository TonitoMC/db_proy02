package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Cargar .env para las credenciales de la DB
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load environment variables")
	}

	// Variables para credenciales DB
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Construir string de conexion a la DB
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	// Realizar conexion a la DB
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	defer db.Close()

	// Ping a la DB solo para verificar conexion
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging DB: %v", err)
	}

	// Se se pingea a la DB indicar que la conexion fue realizada
	fmt.Println("Connected to PostgreSQL!")

	isolationLevels := []string{"READ COMMITTED", "REPEATABLE READ", "SERIALIZABLE"}
	concurrentUsers := []int{5, 10, 20, 30}
	eventID := 1
	seatID := 1

	for _, isolation := range isolationLevels {
		fmt.Printf("Running simulation with isolation level: %s\n", isolation)
		for _, usersNumber := range concurrentUsers {
			runSim(db, usersNumber, isolation, seatID, eventID)
			seatID += 1
		}
	}

	runSim(db, 30, "REPEATABLE READ", 30, 1)
}

// Hashmap de utilidad para el nivel de aislamiento
func mapIsolationLevel(level string) sql.IsolationLevel {
	switch level {
	case "READ COMMITTED":
		return sql.LevelReadCommitted
	case "REPEATABLE READ":
		return sql.LevelRepeatableRead
	case "SERIALIZABLE":
		return sql.LevelSerializable
	default:
		return sql.LevelDefault
	}
}

// Funcion de utilidad para correr una simulacion, se le da una
// cantidad de usuarios y nivel de aislamiento. Imprime los resultados
// de cada ejecucion.
func runSim(db *sql.DB, users int, isolation string, seatID int, eventID int) {
	// Waitgroup para esperar la ejecucion de todas las goroutines
	// y canal de resultados + duracion
	var wg sync.WaitGroup
	resultChan := make(chan bool, users)
	durationChan := make(chan time.Duration, users)
	startSignal := make(chan struct{})

	// Crear las Goroutines
	for i := 0; i < users; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-startSignal
			start := time.Now()

			reserveSeat(db, i+1, eventID, seatID, isolation, &wg, resultChan)
			durationChan <- time.Since(start)
		}(i)
	}
	close(startSignal)
	wg.Wait()
	close(resultChan)
	close(durationChan)

	// Impresion de resultados
	var success, fail int
	for result := range resultChan {
		if result {
			success++
		} else {
			fail++
		}
	}

	var totalDuration int64
	var avgDuration int64
	var maxDuration int64

	for d := range durationChan {
		ms := d.Milliseconds()
		if ms > maxDuration {
			maxDuration = ms
		}
		totalDuration += ms
	}

	avgDuration = totalDuration / int64(users)

	fmt.Printf("Isolation Level: %s | Users: %d | Success: %d | Fail: %d | Avg. Time: %d | Max. Time: %d\n", isolation, users, success, fail, avgDuration, maxDuration)
}

// Funcion para reservar un asiento en la DB
// 1. Leer si el asiento esta reservado o no
// 2. Crear una reservacion para el usuario
// 3. Marcar el asiento como reservado
func reserveSeat(db *sql.DB, userID, eventID, seatID int, isolationLevel string, wg *sync.WaitGroup, resultChan chan<- bool) {
	// Marca Goroutine como terminada al terminar la ejecucion de la funcion

	// Iniciamos la transaccion con el nivel de aislamiento especificado
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: mapIsolationLevel(isolationLevel),
	})
	if err != nil {
		log.Println("BeginTx failed:", err)
		resultChan <- false
		return
	}

	// Verificar que el asiento no este reservado
	var reserved bool
	err = tx.QueryRow(`SELECT reserved FROM event_seats WHERE event_id = $1 AND seat_id = $2`, eventID, seatID).Scan(&reserved)
	if err != nil {
		log.Println("Error checking seat:", err)
		tx.Rollback()
		resultChan <- false
		return
	}

	// Rollback si esta reservado
	if reserved {
		tx.Rollback()
		resultChan <- false
		return
	}

	// Crear reservacion
	var reservationID int
	err = tx.QueryRow(`INSERT INTO reservations (user_id, event_id) VALUES ($1, $2) RETURNING id`, userID, eventID).Scan(&reservationID)
	if err != nil {
		log.Println("Error inserting reservation:", err)
		tx.Rollback()
		resultChan <- false
		return
	}

	// Creat sets_reservations para el asiento
	_, err = tx.Exec(`INSERT INTO seats_reservations (reservation_id, event_id, event_seat_id)
										VALUES ($1, $2, (SELECT id FROM event_seats WHERE event_id = $2 AND seat_id = $3))`,
		reservationID, eventID, seatID)
	if err != nil {
		log.Println("Error inserting into seats_reservations:", err)
		resultChan <- false
		return
	}

	// Marcar asiento como reservado
	_, err = tx.Exec(`UPDATE event_seats SET reserved = true WHERE event_id = $1 AND seat_id = $2`, eventID, seatID)
	if err != nil {
		log.Println("Error updating seat:", err)
		tx.Rollback()
		resultChan <- false
		return
	}

	if err = tx.Commit(); err != nil {
		log.Println("Commit failed:", err)
		resultChan <- false
		return
	}

	resultChan <- true
}
