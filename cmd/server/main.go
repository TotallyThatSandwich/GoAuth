package main

import (
	"context"
	"os"
	"fmt"
	"log/slog"
    "os"
    "path/filepath"

	"github.com/TotallyThatSandwich/GoAuth/internal/repository"
	
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	ctx := context.Background()

	logDir := "logs"
    if err := os.MkdirAll(logDir, 0755); err != nil {
        panic(err)
    }

	// Create/open the log file
    logFile := filepath.Join(logDir, "app.log")
    f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer f.Close()

	// Update defult logger slog logger writing to file
    slog.SetDefault(slog.New(slog.NewJSONHandler(f)))
	
	slog.Info("Application started", "version", "1.0")


	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close(ctx)

	repo := repository.New(ctx, conn, os.Getenv("CACHE_ADR"), os.Getenv("CACHE_PASWD"))
	

	cache_status, db_status := repo.TestRepo(ctx) 
	fmt.Println(cache_status)
	fmt.Println(db_status)
}
