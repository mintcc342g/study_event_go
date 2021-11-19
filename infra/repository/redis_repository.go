package repository

import (
	"study_event_go/domain/interfaces"

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
