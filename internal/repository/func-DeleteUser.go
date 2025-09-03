package repository

import (
	"context"
	"log/slog"
	"fmt"

	"github.com/TotallyThatSandwich/GoAuth/internal/sqlc"
)

func (r *UserRepository) DeleteUser(ctx context.Context, user sqlc.User) error {

	slog.Info("Checking if user in cache", "username", user.Username, "ID", user.UserID)
	
	_, err := r.cache.Get(ctx, fmt.Sprintf("user:%s:%s", user.Username, user.HashedPassword)).Result()

	if err == nil {
		slog.Info("Deleting User from cache")
		_, err := r.cache.Del(ctx, fmt.Sprintf("user:%s:%s", user.Username, user.HashedPassword)).Result()
		if err == nil {
			slog.Info("Deleted User from cache")
		} else {
			slog.Error("Failed to delete user from cache")
			return err
		}
	} else {
		slog.Info("user not in cache")
	}	

			
	slog.Info("Deleting User from db", "username", user.Username, "ID", user.UserID)
	err = r.db.DeleteUser(ctx, r.dbtx, sqlc.DeleteUserParams{UserID: user.UserID, UserToken: user.UserToken})
	if err == nil {
		slog.Info("Deleted User from db")
		return nil
	} else {
		slog.Error("Filed to delete user from db")
		return err	
	}

} 
