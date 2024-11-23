package migration

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(connStr string) (err error) {
	migrationPath := filepath.Join("db", "migrations")

	sqlDB, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer sqlDB.Close()

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	driver, err := pgx.WithInstance(sqlDB, &pgx.Config{})
	if err != nil {
		log.Fatalf("Failed to create pgx driver: %v", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationPath),
		"pgx",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
		return
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
		return
	}

	log.Println("Migrations applied successfully!")
	return nil
}
