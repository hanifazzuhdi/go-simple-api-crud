package main

import (
	"context"
	"log"
	"os"
	"simple-api/internal/app"
)

func main() {

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres://root:password@localhost:5432/go_simple_user_api"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ctx := context.Background()
	application := app.New(ctx, dbUrl)

	defer application.Shutdown(ctx)

	err := application.Run(port)
	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}
