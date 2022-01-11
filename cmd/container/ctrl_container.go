package container

import (
	"study-event-go/controller"
)

// ControllerContainer ...
type ControllerContainer struct {
	Charm      *controller.CharmController
	Garden     *controller.GardenController
	Lily       *controller.LilyController
	Mentorship *controller.MentorshipController
	Protection *controller.ProtectionController
	Skill      *controller.SkillController
}

func newControllerContainer(svcContainer *ServiceContainer, repoContainer *RepositoryContainer) *ControllerContainer {
	return &ControllerContainer{
		Charm:      controller.NewCharmController(svcContainer.CharmSvc),
		Garden:     controller.NewGardenController(svcContainer.GardenSvc),
		Lily:       controller.NewLilyController(svcContainer.LilySvc),
		Mentorship: controller.NewMentorshipController(svcContainer.MentorshipSvc),
		Protection: controller.NewProtectionController(svcContainer.ProtectionSvc),
		Skill:      controller.NewSkillController(svcContainer.SkillSvc),
	}
}

// InitControllerContainer ...
func InitControllerContainer(svcContainer *ServiceContainer, repoContainer *RepositoryContainer) *ControllerContainer {
	return newControllerContainer(svcContainer, repoContainer)
}
