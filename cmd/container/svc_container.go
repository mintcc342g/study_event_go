package container

import (
	"study_event_go/application"
)

// ServiceContainer ...
type ServiceContainer struct {
	ProtectionSvc *application.ProtectionService
	EventSvc      *application.EventService
}

func newServiceContainer(repo *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		ProtectionSvc: application.NewProtectionService(repo.LilyRepo, repo.GardenRepo, repo.EventRepo),
		EventSvc:      application.NewEventService(repo.RedisRepo),
	}
}

// InitServiceContainer ...
func InitServiceContainer(repo *RepositoryContainer) *ServiceContainer {
	return newServiceContainer(repo)
}
