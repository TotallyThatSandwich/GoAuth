package repository

import (
	"context"
	"log/slog"

	"github.com/TotallyThatSandwich/GoAuth/internal/sqlc"
)

func (r *UserRepository) CreateUser(ctx context.Context, username string, hashedPass string) (*sqlc.User, error) { 

	user, err := r.db.CreateUser(ctx, r.dbtx, sqlc.CreateUserParams{Username: username, HashedPassword: hashedPass})
	if err != nil {
		slog.Error("Failed to create user", "err", err)
		return nil, err
	} else {
		slog.Info("User Created", "username", username, "password", hashedPass, "token", user.UserToken)

		return &user, nil
	}
}
