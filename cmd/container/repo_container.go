package container

import (
	"context"
	"study-event-go/domain/interfaces"
	"study-event-go/ent"
	"study-event-go/ent/migrate"
	"study-event-go/infra/repository"

	"github.com/RichardKnop/machinery/v1"
	"github.com/go-redis/redis/v8"
)

// RepositoryContainer ...
type RepositoryContainer struct {
	EventRepo      interfaces.EventRepository
	GardenRepo     interfaces.GardenRepository
	LilyRepo       interfaces.LilyRepository
	MentorshipRepo interfaces.MentorshipRepository
	RedisRepo      interfaces.RedisRepository
	SkillRepo      interfaces.SkillRepository
}

func newRepositoryContainer(db *ent.Client, redis redis.Cmdable, machineryServer *machinery.Server) *RepositoryContainer {
	return &RepositoryContainer{
		EventRepo:      repository.NewMachineryRepository(machineryServer),
		GardenRepo:     repository.NewGardenRepository(db),
		LilyRepo:       repository.NewLilyRepository(db),
		MentorshipRepo: repository.NewMentorshipRepository(db),
		RedisRepo:      repository.NewRedisRepository(redis),
		SkillRepo:      repository.NewSkillRepository(db),
	}
}

// InitRepositoryContainer ...
func InitRepositoryContainer(db *ent.Client, redis redis.Cmdable, machineryServer *machinery.Server) (*RepositoryContainer, error) {
	ctx := context.Background()
	if err := db.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		println("InitRepositoryContainer", "Create", "error", err) // TODO: logger
		return nil, err
	}

	return newRepositoryContainer(db, redis, machineryServer), nil
}
