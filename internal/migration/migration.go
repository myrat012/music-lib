package migration

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(connStr string) (err error) {
	migrationPath := filepath.Join("db", "migrations")

	m, err := migrate.New(fmt.Sprintf("file://%s", migrationPath), connStr)
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
