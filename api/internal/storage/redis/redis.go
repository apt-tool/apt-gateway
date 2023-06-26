package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type connector struct {
	connection *redis.Client
}

func (c *connector) Get(key string) (string, error) {
	ctx := context.Background()

	value, err := c.connection.Get(ctx, key).Result()
	if err != nil {
		return "", ErrSetNotFound
	}

	return value, nil
}
