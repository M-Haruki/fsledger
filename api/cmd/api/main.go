package main

import (
	"context"
	"log"

	"github.com/M-Haruki/fsledger/api/internal/server"
)

func main() {

	// e := echo.New()
	// // e.GET("/", func(c *echo.Context) error {
	// // 	return c.JSON(200, map[string]string{"message": "Hello, World!"})
	// // })

	// e.Start(":1323")

	ctx := context.Background()

	server, err := server.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(":1323"); err != nil {
		log.Fatal(err)
	}
}
