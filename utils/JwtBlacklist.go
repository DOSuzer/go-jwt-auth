package utils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func AddToBlacklist(token string, expiry time.Duration) error {
	err := rdb.Set(ctx, token, "revoked", expiry).Err()
	if err != nil {
		return err
	}
	return nil
}

func IsBlacklisted(token string) (bool, error) {
	result, err := rdb.Get(ctx, token).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	if result == "revoked" {
		return true, nil
	}

	return false, nil
}
