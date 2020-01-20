package db

import (
	"log"

	"github.com/jackc/pgx"
)

var (
	pgPool *pgx.ConnPool
)

// Connect is
func Connect() {
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			User:     "postgres",
			Password: "1",
			Database: "quote.alif.db",
		},
	}
	var err error
	pgPool, err = pgx.NewConnPool(connPoolConfig)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}
}

// Close is
func Close() {
	pgPool.Close()
}
