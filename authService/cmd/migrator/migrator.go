package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var migrationsPath string
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.Parse()

	if migrationsPath == "" {
		log.Fatal("migrations path is empty")
	}

	dsn := "postgres://postgres:postgres@localhost:5434/authdb?sslmode=disable"

	migrator, err := migrate.New("file://"+migrationsPath, dsn)
	if err != nil {
		log.Fatalf("failed to create migrator: %v", err)
	}
	defer migrator.Close()

	err = migrator.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations found")
			return
		}
		log.Fatalf("failed to apply migrations: %v", err)
	}

	fmt.Println("Migrations applied successfully")
}
