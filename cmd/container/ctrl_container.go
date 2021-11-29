package container

import (
	"study_event_go/controller"
)

// ControllerContainer ...
type ControllerContainer struct {
	Protection *controller.ProtectionController
}

func newControllerContainer(svcContainer *ServiceContainer, repoContainer *RepositoryContainer) *ControllerContainer {
	return &ControllerContainer{
		Protection: controller.NewProtectionController(svcContainer.ProtectionSvc),
	}
}

// InitControllerContainer ...
func InitControllerContainer(svcContainer *ServiceContainer, repoContainer *RepositoryContainer) *ControllerContainer {
	return newControllerContainer(svcContainer, repoContainer)
}
