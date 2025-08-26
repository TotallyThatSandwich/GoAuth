
package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func New(address string, password string) *redis.Client {
    return redis.NewClient(&redis.Options{
		Addr: address,
        Password: password,
        DB: 0,
        Protocol: 2,
	})
}
