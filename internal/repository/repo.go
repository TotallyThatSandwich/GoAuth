package repository

import (
	"context"

	"github.com/TotallyThatSandwich/GoAuth/internal/cache"
	"github.com/TotallyThatSandwich/GoAuth/internal/sqlc"
)

func New(ctx context.Context, db_url string, cache_url string) *UserRepository {
	return UserRepository{db: sqlc.New(pgx.Connect(ctx, db_url)), cache: cache.New(cache_url)}
}

type UserRepository struct {
    db    *sqlc.Queries
    cache *redis.Client
}

func (r *UserRepository) CheckUserAuth(ctx context.Context, username string, password_hash string) (*sqlc.User, error) {
    // check redis first
	key := fmt.Sprintf("user:%s:%s", username, password_hash)
    if val, err := r.cache.Get(ctx, key).Result(); err == nil {
        var u sqlc.User
        _ = json.Unmarshal([]byte(val), &u)
		return &u, nil
    }

    user, err := r.db.CheckUserAuth(ctx, sqlc.CheckUserAuthParams{Username username, HashedPassword password_hash})
    if err != nil {
        return nil, err
    }

    // update cache
    data, _ := json.Marshal(user)
	key := fmt.Sprintf("user:%s:%s", username, password_hash)
    r.cache.Set(ctx, key, data, time.Hour)
    return user, nil
}


