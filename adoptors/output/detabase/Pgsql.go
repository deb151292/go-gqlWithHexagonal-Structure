package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func DatabaseConn() {
	// Retrieve PostgreSQL connection details from environment variables
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	sslMode := os.Getenv("POSTGRES_SSLMODE")

	log.Println(host, port, user, password, dbName, sslMode)

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslMode)

	// Create a connection pool with additional configurations
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("Error parsing database URL: %v", err)
	}

	// Set connection timeout and reconnection attempts
	config.ConnConfig.ConnectTimeout = 5 * time.Second
	config.MaxConnLifetime = 10 * time.Minute // Set 0 For Infinite Time
	config.MaxConns = 5

	poolConn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}
	pool = poolConn
	defer pool.Close()

	// Example query
	var result string
	err = pool.QueryRow(context.Background(), "SELECT $1::text", "PostgreSQL Connected").Scan(&result)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}

	fmt.Printf("Database : %s\n", result)
}
