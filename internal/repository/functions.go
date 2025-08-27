package repository

import (
	"context"
)

func (r *UserRepository) TestRepo(ctx context.Context) (string, string) {
	
	status := [2]bool{true, true}	//1: cache 2: DB

	// write "alive":"true" to cache
	err := r.cache.Set(ctx, "alive", "true", 0).Err()
	if err != nil {
    	status[1] = false
	}

	// read "alive" key from cache
	_, err = r.cache.Get(ctx, "alive").Result()
	if err != nil {
    	status[1] = false
	} 

	// run healtcheck on DB 
	// sql:'SELECT 1 AS alive;' 
	_, err = r.db.HealthCheck(context.Background())
	if err != nil {
		status[1] = false
	} 

	switch status {
	case [2]bool{true, true}:
		return "Cache: alive", "Database: alive"
	case [2]bool{true, false}:
		return "Cache: alive", "Database: dead"
	case [2]bool{false, true}:
		return "Cache: dead", "Database: alive"
	default:
		return "Cache: dead", "Database: dead"
	}	
}

