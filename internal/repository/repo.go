package repository

import (
	"context"

	"github.com/TotallyThatSandwich/GoAuth/internal/cache"
	"github.com/TotallyThatSandwich/GoAuth/internal/sqlc"
)

func New(ctx context.Context, db_url string, cache_addr string, cache_paswd string) *UserRepository {
	return UserRepository{db: sqlc.New(pgx.Connect(ctx, db_url)), cache: cache.New(cache_addr, cache_paswd)}
}

type UserRepository struct {
    db    *sqlc.Queries
    cache *redis.Client
}

func (r *UserRepository) TestRepo(ctx context.Context) string {

	err := r.cache.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
    	panic(err)
	}

	val, err := r.cache.Get(ctx, "foo").Result()
	if err != nil {
    	panic(err)
	}
	return val
}


