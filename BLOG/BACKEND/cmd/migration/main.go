package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/yafiakmal/golang-mini-project/blog/config"
	"github.com/yafiakmal/golang-mini-project/blog/setup"
)

func main() {
	setup.LoadEnv()

	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go [up|down|drop]")
	}
	action := os.Args[1]

	db, err := sql.Open("postgres", config.GetPostgresURL())
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Migration driver setup failed: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("Migration initialization failed: %v", err)
	}

	switch action {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("migrate up failed: %v", err)
		}
		fmt.Println("✅ Migrasi berhasil dijalankan!")

	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatalf("migrate down failed: %v", err)
		}
		fmt.Println("↩️  Migrasi berhasil di-rollback satu langkah!")

	case "drop":
		if err := m.Drop(); err != nil {
			log.Fatalf("migrate drop failed: %v", err)
		}
		fmt.Println("🗑️  Semua tabel migrasi dihapus!")

	default:
		log.Fatalf("Unknown command: %s (gunakan up|down|drop)", action)
	}
}
