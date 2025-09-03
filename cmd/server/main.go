package main

import (
	"context"
	"os"
	"log/slog"

	"github.com/TotallyThatSandwich/GoAuth/internal/repository"
	"github.com/TotallyThatSandwich/GoAuth/internal/api"
	
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	ctx := context.Background()

	// Update defult logger slog logger writing to file
    slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
	

	slog.Info("Application started", "version", "1.0")


	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close(ctx)


	repo := repository.New(ctx, conn, os.Getenv("CACHE_ADR"), os.Getenv("CACHE_PASWD"))
	repo.TestRepo(ctx)


	apiServer := api.New("localhost:8080")
	apiServer.Run(ctx, repo, "/api/v1")
}
