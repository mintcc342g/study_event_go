package container

import (
	"study-event-go/controller"
)

// ControllerContainer ...
type ControllerContainer struct {
	Garden     *controller.GardenController
	Mentorship *controller.MentorshipController
	Protection *controller.ProtectionController
}

func newControllerContainer(svcContainer *ServiceContainer, repoContainer *RepositoryContainer) *ControllerContainer {
	return &ControllerContainer{
		Garden:     controller.NewGardenController(svcContainer.GardenSvc),
		Mentorship: controller.NewMentorshipController(svcContainer.MentorshipSvc),
		Protection: controller.NewProtectionController(svcContainer.ProtectionSvc),
	}
}

// InitControllerContainer ...
func InitControllerContainer(svcContainer *ServiceContainer, repoContainer *RepositoryContainer) *ControllerContainer {
	return newControllerContainer(svcContainer, repoContainer)
}
