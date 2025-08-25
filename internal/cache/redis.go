
package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func New(db string) *redis.Client {
	opt, err := redis.ParseURL(db)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}
