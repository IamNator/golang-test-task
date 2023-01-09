package redis

import "github.com/go-redis/redis"

type (
	Redis struct {
		Client *redis.Client
	}
)

func NewClient(addr string, password string, db int) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &Redis{Client: client}
}
