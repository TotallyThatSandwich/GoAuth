package repository

import (
	"context"
	"log/slog"
	"fmt"

	"encoding/json"

	"github.com/TotallyThatSandwich/GoAuth/internal/sqlc"
)

func (r *UserRepository) UpdateUser(ctx context.Context, user sqlc.User) (*sqlc.User, error) {
	
	oldUser, _ := r.db.GetUserFromID(ctx, r.dbtx, user.UserID)

	slog.Info("Checking if user in cache", "ID", user.UserID)

	_, err := r.cache.Get(ctx, fmt.Sprintf("user:%s:%s", oldUser.Username, oldUser.HashedPassword)).Result()

	if err == nil {
		slog.Info("Updating User in cache")
		userJSON, err := json.Marshal(user)
		if err != nil {
			slog.Error("Failed to marshal user")
		} else {
			_, err = r.cache.Del(ctx, fmt.Sprintf("user:%s:%s", oldUser.Username, oldUser.HashedPassword)).Result()
			if err == nil {
				err = r.cache.Set(ctx, fmt.Sprintf("user:%s:%s", user.Username, user.HashedPassword), userJSON, 0).Err()
				if err != nil {
					slog.Error("Failed to write to cache")
				} else {
					slog.Info("Updated user in cache")
				}
			} else {
				slog.Error("Failed to delete stale data from cache")
			}
		}
	} else {
		slog.Info("user not in cache")
	}
	
	slog.Info("Updating user in db", "ID", user.UserID)

	newUser, err := r.db.UpdateUser(ctx, r.dbtx, sqlc.UpdateUserParams{Username: user.Username, HashedPassword: user.HashedPassword, UserID: user.UserID})
	if err == nil {
		return &newUser, nil
	} else {
		slog.Error("Failed to write to db")
		return nil, err
	}
}
