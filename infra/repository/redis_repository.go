package repository

import (
	"study-event-go/domain/interfaces"

	"github.com/go-redis/redis/v8"
)

type redisRepository struct {
	rds redis.Cmdable
}

// NewRedisRepository ...
func NewRedisRepository(rds redis.Cmdable) interfaces.RedisRepository {
	return &redisRepository{
		rds: rds,
	}
}
