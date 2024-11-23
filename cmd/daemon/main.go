package main

import (
	"fmt"
	"log"

	"github.com/myrat012/test-work-song-lib/db"
	"github.com/myrat012/test-work-song-lib/internal/migration"
	"github.com/myrat012/test-work-song-lib/pkg/config"
)

func main() {
	// Load .env
	conf, err := config.LoadEnv(".env")
	if err != nil {
		return
	}

	// Create PostgreSQL connection string for pgx
	pool, err := db.NewPool(conf)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	fmt.Println("END")

	if err := migration.RunMigrations(conf.ConnString); err != nil {
		log.Fatalf("Unable to run migrations: %v", err)
	}

}
