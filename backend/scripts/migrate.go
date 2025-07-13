package scripts

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MigrateDb() {
	// Initialize database connection
	db := initPostgres()
	defer db.Close()

	// Create migrations tracking table
	if err := createMigrationsTable(db); err != nil {
		log.Fatal("Failed to create migrations table:", err)
	}

	// Run migrations from postgres folder
	if err := runMigrationsFromFolder(db, "./migrations/postgres"); err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("All migrations completed!")
}

func initPostgres() *pgxpool.Pool {
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
	return pool
}

func createMigrationsTable(db *pgxpool.Pool) error {
	ctx := context.Background()
	query := `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := db.Exec(ctx, query)
	return err
}

func runMigrationsFromFolder(db *pgxpool.Pool, folderPath string) error {
	// Get all .sql files
	files, err := findMigrationFiles(folderPath)
	if err != nil {
		return err
	}

	// Sort files by version/name
	sort.Strings(files)

	ctx := context.Background()
	for _, file := range files {
		version := extractVersion(file)
		
		// Check if already applied
		var count int
		err := db.QueryRow(ctx, "SELECT COUNT(*) FROM schema_migrations WHERE version = $1", version).Scan(&count)
		if err != nil {
			return err
		}

		if count > 0 {
			fmt.Printf("Migration %s already applied\n", version)
			continue
		}

		// Read and execute migration
		if err := executeMigrationFile(db, file, version); err != nil {
			return fmt.Errorf("failed to execute %s: %w", file, err)
		}

		fmt.Printf("Applied migration: %s\n", version)
	}

	return nil
}

func findMigrationFiles(folderPath string) ([]string, error) {
	var files []string
	
	err := filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		// Only include .up.sql files for migrations
		if !d.IsDir() && strings.HasSuffix(path, ".up.sql") {
			files = append(files, path)
		}
		
		return nil
	})
	
	return files, err
}

func extractVersion(filename string) string {
	// Extract version from filename (e.g., "001_create_orders_table.up.sql" -> "001")
	base := filepath.Base(filename)
	// Remove .up.sql suffix first
	name := strings.TrimSuffix(base, ".up.sql")
	parts := strings.Split(name, "_")
	if len(parts) > 0 {
		return parts[0]
	}
	return name
}

func executeMigrationFile(db *pgxpool.Pool, filepath, version string) error {
	// Read file content
	content, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filepath, err)
	}

	// Validate SQL content
	sqlContent := strings.TrimSpace(string(content))
	if sqlContent == "" {
		return fmt.Errorf("empty migration file: %s", filepath)
	}

	fmt.Printf("Executing migration %s...\n", version)
	fmt.Printf("SQL content preview: %s\n", sqlContent[:min(100, len(sqlContent))])

	ctx := context.Background()
	
	// Begin transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Acquire advisory lock to prevent concurrent migrations
	if _, err := tx.Exec(ctx, "SELECT pg_advisory_lock(12345)"); err != nil {
		return fmt.Errorf("failed to acquire lock: %w", err)
	}

	// Split SQL content by semicolons and execute each statement
	statements := strings.Split(sqlContent, ";")
	for i, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		
		fmt.Printf("Executing statement %d: %s\n", i+1, stmt[:min(50, len(stmt))])
		
		if _, err := tx.Exec(ctx, stmt); err != nil {
			return fmt.Errorf("migration SQL failed at statement %d (%s): %w", i+1, stmt, err)
		}
	}

	// Record migration
	if _, err := tx.Exec(ctx,
		"INSERT INTO schema_migrations (version) VALUES ($1)",
		version,
	); err != nil {
		return fmt.Errorf("failed to record migration: %w", err)
	}

	// Release lock and commit
	if _, err := tx.Exec(ctx, "SELECT pg_advisory_unlock(12345)"); err != nil {
		return fmt.Errorf("failed to release lock: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Printf("Migration %s completed successfully\n", version)
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}