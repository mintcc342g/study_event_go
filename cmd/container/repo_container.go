package container

import (
	"study_event_go/domain/interfaces"
	"study_event_go/infra/repository"

	"github.com/RichardKnop/machinery/v1"
	"github.com/go-redis/redis/v8"
)

// RepositoryContainer ...
type RepositoryContainer struct {
	RedisRepo  interfaces.RedisRepository
	LilyRepo   interfaces.LilyRepository
	GardenRepo interfaces.GardenRepository
	EventRepo  interfaces.EventRepository
}

func newRepositoryContainer(redis redis.Cmdable, machineryServer *machinery.Server) *RepositoryContainer {
	return &RepositoryContainer{
		RedisRepo:  repository.NewRedisRepository(redis),
		LilyRepo:   repository.NewLilyRepository(),
		GardenRepo: repository.NewGardenRepository(),
		EventRepo:  repository.NewMachineryRepository(machineryServer),
	}
}

// InitRepositoryContainer ...
func InitRepositoryContainer(redis redis.Cmdable, machineryServer *machinery.Server) *RepositoryContainer {
	return newRepositoryContainer(redis, machineryServer)
}
