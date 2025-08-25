package main

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()

	repo := repository.New(ctx, os.Getenv("DATABASE_URL"), os.Getenv("CACHE_URL"))
}
