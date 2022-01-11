package container

import (
	"study-event-go/application"
)

// ServiceContainer ...
type ServiceContainer struct {
	CharmSvc      *application.CharmService
	EventSvc      *application.EventService
	GardenSvc     *application.GardenService
	LilySvc       *application.LilyService
	MentorshipSvc *application.MentorshipService
	ProtectionSvc *application.ProtectionService
	SkillSvc      *application.SkillService
}

func newServiceContainer(repo *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		CharmSvc:      application.NewCharmService(repo.CharmRepo),
		EventSvc:      application.NewEventService(repo.RedisRepo),
		GardenSvc:     application.NewGardenService(repo.GardenRepo, repo.MentorshipRepo),
		LilySvc:       application.NewLilyService(repo.CharmRepo, repo.GardenRepo, repo.LilyRepo),
		MentorshipSvc: application.NewMentorshipService(repo.MentorshipRepo),
		ProtectionSvc: application.NewProtectionService(repo.LilyRepo, repo.GardenRepo, repo.EventRepo),
		SkillSvc:      application.NewSkillService(repo.SkillRepo),
	}
}

// InitServiceContainer ...
func InitServiceContainer(repo *RepositoryContainer) *ServiceContainer {
	return newServiceContainer(repo)
}
