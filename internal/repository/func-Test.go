package repository

import (
	"context"
	"log/slog"
)

func (r *UserRepository) TestRepo(ctx context.Context) {
	
	slog.Info("testing Redis connection...")

	// write "alive":"true" to cache
	err := r.cache.Set(ctx, "alive", "true", 0).Err()
	if err != nil {
		slog.Error("Redis connection test failed. Exiting app")
		panic(err)
	}
	// read "alive" key from cache
	_, err = r.cache.Get(ctx, "alive").Result()
	if err != nil {
		slog.Error("Redis connection test failed. Exiting app")
		panic(err)
	} 
	
	slog.Info("Redis test sucsessfull")
	slog.Info("testing Postgres connection...")

	// run healtcheck on DB 
	// sql:'SELECT 1 AS alive;' 
	_, err = r.db.HealthCheck(context.Background(), r.dbtx)
	if err != nil {
		slog.Error("Postgress connection test failed. Exiting app")
		panic(err)
	}
	slog.Info("Postgres test sucsessfull")
}
