package main

import (
	"context"
	"log"
	"os"

	"github.com/M-Haruki/fsledger/api/internal/server"
)

func main() {
	ctx := context.Background()

	cfg := server.Config{
		IsDev: os.Getenv("APP_ENV") == "dev",
	}

	server, err := server.New(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(":1323"); err != nil {
		log.Fatal(err)
	}
}
