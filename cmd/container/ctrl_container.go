package container

import (
	"study_event_go/controller"
)

type ControllerContainer struct {
	Protection *controller.ProtectionController
}

func newControllerContainer(svcContainer *ServiceContainer, repoContainer *RepositoryContainer) *ControllerContainer {
	return &ControllerContainer{
		Protection: controller.NewProtectionController(svcContainer.BattleSvc),
	}
}

func InitControllerContainer(svcContainer *ServiceContainer, repoContainer *RepositoryContainer) *ControllerContainer {
	return newControllerContainer(svcContainer, repoContainer)
}
