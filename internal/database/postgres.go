package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}
	if err := migrate(db); err != nil {
		log.Fatalf("migration error: %v", err)
	}
	log.Println("database connected")
	return db
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id           SERIAL PRIMARY KEY,
			email        VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			created_at   TIMESTAMP DEFAULT NOW()
		);

		CREATE TABLE IF NOT EXISTS events (
			id          SERIAL PRIMARY KEY,
			title       VARCHAR(255) NOT NULL,
			description TEXT,
			location    VARCHAR(255),
			event_date  DATE NOT NULL,
			event_time  TIME NOT NULL,
			event_type  VARCHAR(100),
			created_at  TIMESTAMP DEFAULT NOW(),
			updated_at  TIMESTAMP DEFAULT NOW()
		);
	`)
	return err
}
