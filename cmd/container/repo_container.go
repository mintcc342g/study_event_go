package container

import (
	"study_event_go/domain/interfaces"
	"study_event_go/infra/repository"

	"github.com/go-redis/redis/v8"
)

type RepositoryContainer struct {
	RedisRepo   interfaces.RedisRepository
	StudentRepo interfaces.StudentRepository
}

func newRepositoryContainer(redis redis.Cmdable) *RepositoryContainer {
	return &RepositoryContainer{
		RedisRepo:   repository.NewRedisRepository(redis),
		StudentRepo: repository.NewStudentRepository(),
	}
}

func InitRepositoryContainer(redis redis.Cmdable) *RepositoryContainer {
	return newRepositoryContainer(redis)
}
