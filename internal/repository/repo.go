package repository

import (
	"context"
	
	"github.com/jackc/pgx/v5"

	"github.com/TotallyThatSandwich/GoAuth/internal/cache"
	"github.com/TotallyThatSandwich/GoAuth/internal/sqlc"
	"github.com/redis/go-redis/v9"
)

func New(ctx context.Context, conn *pgx.Conn, cache_addr string, cache_paswd string) *UserRepository {
	tx, _ := conn.Begin(ctx)

	return &UserRepository{db: sqlc.New(tx), cache: cache.New(cache_addr, cache_paswd)}
}

// UserRepository manages db and cache
type UserRepository struct {
    db *sqlc.Queries
    cache *redis.Client
}

