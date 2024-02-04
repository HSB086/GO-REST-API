package database

/* import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx"
)*/

/*var DB *sql.DB
var err error

func InitDB() {
	connectionString := "user=postgres password=root host=localhost port=5432 database=events_db sslmode=disable"
	DB, err = sql.Open("pgx", connectionString)

	if err != nil {
		fmt.Printf("cause: %v\n", err)
		panic("Could not connect to DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
} */

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	connString := "postgresql://postgres:root@localhost:5432/events_db"

	config, err := pgx.ParseConfig(connString)
	if err != nil {
		panic(err)
	}
	DB = stdlib.OpenDB(*config)

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	if err != nil {
		panic(err)
	}

	createTables()
}

func createTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		fmt.Printf("cause: %v", err)
		panic("Could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL, 
		location TEXT NOT NULL,
		date_time TIMESTAMPTZ NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		fmt.Printf("cause: %v", err)
		panic("Could not create events table")
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations(
		id SERIAL PRIMARY KEY,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		fmt.Printf("cause: %v", err)
		panic("Could not create registration table")
	}

}
