package container

import (
	"study_event_go/domain/interfaces"
	"study_event_go/infra/repository"

	"github.com/go-redis/redis/v8"
)

// RepositoryContainer ...
type RepositoryContainer struct {
	RedisRepo  interfaces.RedisRepository
	LilyRepo   interfaces.LilyRepository
	GardenRepo interfaces.GardenRepository
	EventRepo  interfaces.EventRepository
}

func newRepositoryContainer(redis redis.Cmdable) *RepositoryContainer {
	return &RepositoryContainer{
		RedisRepo:  repository.NewRedisRepository(redis),
		LilyRepo:   repository.NewLilyRepository(),
		GardenRepo: repository.NewGardenRepository(),
		EventRepo:  repository.NewMachineryRepository(),
	}
}

// InitRepositoryContainer ...
func InitRepositoryContainer(redis redis.Cmdable) *RepositoryContainer {
	return newRepositoryContainer(redis)
}
