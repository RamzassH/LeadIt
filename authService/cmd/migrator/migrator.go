package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string
	var forceVersion int

	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.IntVar(&forceVersion, "force-version", 0, "force version")
	flag.Parse()

	if storagePath == "" || migrationsPath == "" {
		panic("storage or migrations path is empty")
	}

	dsn := "postgres://postgres:postgres@localhost:5434/authdb?sslmode=disable"

	migrator, err := migrate.New("file://"+migrationsPath, dsn)
	if err != nil {
		panic(err)
	}

	if err := migrator.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations found")
			return
		}
		panic(err)
	}

	fmt.Println("Migrations applied successfully")
}
