package container

import (
	"study_event_go/service"
)

type ServiceContainer struct {
	BattleSvc *service.BattleService
	EventSvc  *service.EventService
}

func newServiceContainer(repo *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		BattleSvc: service.NewBattleService(repo.StudentRepo),
		EventSvc:  service.NewEventService(repo.RedisRepo),
	}
}

func InitServiceContainer(repo *RepositoryContainer) *ServiceContainer {
	return newServiceContainer(repo)
}
