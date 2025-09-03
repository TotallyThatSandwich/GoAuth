package repository

import (
	"context"
	"log/slog"
	"fmt"

	"encoding/json"

	"github.com/TotallyThatSandwich/GoAuth/internal/sqlc"
)

func (r *UserRepository) GetUser(ctx context.Context, username string, hashedPass string) (*sqlc.User, error) { 

	slog.Info("Checking redis cache for user.", "username", username, "password", hashedPass)

	cache_val, err := r.cache.Get(ctx, fmt.Sprintf("user:%s:%s", username, hashedPass)).Result()
	if err == nil {
		slog.Info("Found User In Cache.", "username", username, "password", hashedPass)
		var user sqlc.User
		err := json.Unmarshal([]byte(cache_val), &user)
		if err != nil {
        	return nil, err
    	}
    	return &user, nil
	} else {
		slog.Error("could not find user in cache", "username", username, "password", hashedPass)
	}
	
	user, err := r.db.GetUser(ctx, r.dbtx, sqlc.GetUserParams{Username: username, HashedPassword: hashedPass})
	if err == nil {
		slog.Info("Found User In db.", "username", username, "password", hashedPass)
		slog.Info("Writing user to cache.", "username", username, "password", hashedPass)
		userJSON, err := json.Marshal(user)
		if err != nil {
			slog.Error("Failed to marshal user")
		}
		err = r.cache.Set(ctx, fmt.Sprintf("user:%s:%s", username, hashedPass), userJSON, 0).Err()
		if err != nil {
			slog.Error("Failed to write to cache")
		}

		return &user, nil

	} else {
		slog.Error("could not find user in db", "username", username, "password", hashedPass)
		return nil, err
	}
}

