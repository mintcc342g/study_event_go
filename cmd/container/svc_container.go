package container

import (
	"study_event_go/application"
)

// ServiceContainer ...
type ServiceContainer struct {
	BattleSvc *application.BattleService
	EventSvc  *application.EventService
}

func newServiceContainer(repo *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		BattleSvc: application.NewBattleService(repo.StudentRepo),
		EventSvc:  application.NewEventService(repo.RedisRepo),
	}
}

// InitServiceContainer ...
func InitServiceContainer(repo *RepositoryContainer) *ServiceContainer {
	return newServiceContainer(repo)
}
