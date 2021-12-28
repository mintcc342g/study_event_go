package container

import (
	"study-event-go/application"
)

// ServiceContainer ...
type ServiceContainer struct {
	ProtectionSvc *application.ProtectionService
	EventSvc      *application.EventService
	MentorshipSvc *application.MentorshipService
}

func newServiceContainer(repo *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		ProtectionSvc: application.NewProtectionService(repo.LilyRepo, repo.GardenRepo, repo.EventRepo),
		EventSvc:      application.NewEventService(repo.RedisRepo),
		MentorshipSvc: application.NewMentorshipService(repo.MentorshipRepo),
	}
}

// InitServiceContainer ...
func InitServiceContainer(repo *RepositoryContainer) *ServiceContainer {
	return newServiceContainer(repo)
}
