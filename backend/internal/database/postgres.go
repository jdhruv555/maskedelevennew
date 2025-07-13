package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PostgresPool *pgxpool.Pool

func InitPostgres() {
	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		log.Fatal("POSTGRES_URL not set in environment")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Failed to create PostgreSQL pool: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	log.Println("PostgreSQL connected")
	PostgresPool = pool
}
