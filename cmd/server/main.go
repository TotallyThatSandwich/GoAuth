package main

import (
	"context"
	"os"
	"fmt"

	"github.com/TotallyThatSandwich/GoAuth/internal/sqlc"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	fmt.Println("code running")

	ctx := context.Background()

	repo := repository.New(ctx, os.Getenv("DATABASE_URL"), os.Getenv("CACHE_ADR"), os.Getenv("CACHE_PASWD"))

	fmt.Println(repo.TestRepo(ctx))
}
