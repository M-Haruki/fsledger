package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/M-Haruki/fsledger/api/internal/db/migrations"
	"github.com/pressly/goose/v3"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	var err error

	// get args
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: migrate <status|up|down>")
	}
	command := os.Args[1]
	if command != "status" && command != "up" && command != "down" {
		return fmt.Errorf("usage: migrate <status|up|down>")
	}

	// connect db
	var db *sql.DB
	databaseUrl := os.Getenv("DATABASE_URL")
	db, err = sql.Open("pgx", databaseUrl)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}

	// setting goose
	goose.SetBaseFS(migrations.FS)
	err = goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	// run command
	switch command {
	case "status":
		err = goose.Status(db, ".")
	case "up":
		err = goose.Up(db, ".")
	case "down":
		err = goose.Down(db, ".")
	}
	return err
}
