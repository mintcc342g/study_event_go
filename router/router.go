package router

import (
	"study-event-go/cmd/container"

	"github.com/labstack/echo/v4"
)

// AssaultLilyRoute ...
func AssaultLilyRoute(sys *echo.Group, ctrlContainer *container.ControllerContainer) {

	// TODO: check authority with jwt

	sys.POST("/mentorships", ctrlContainer.Mentorship.New)
	sys.GET("/mentorships/:id", ctrlContainer.Mentorship.Get)
	sys.GET("/mentorships", ctrlContainer.Mentorship.List)
	sys.PUT("/mentorships/:id", ctrlContainer.Mentorship.Update)
	sys.DELETE("/mentorships/:id", ctrlContainer.Mentorship.Delete)

	sys.POST("/gardens", ctrlContainer.Garden.New)
	sys.GET("/gardens/:id", ctrlContainer.Garden.Get)
	sys.GET("/gardens", ctrlContainer.Garden.List)
	sys.PUT("/gardens/:id", ctrlContainer.Garden.Update)
	// sys.PUT("/gardens/:id/lilies", ctrlContainer.Garden.UpdateLilies)
	sys.DELETE("/gardens/:id", ctrlContainer.Garden.Delete)

	// TODO
	sys.POST("/skills", ctrlContainer.Skill.New)
	// sys.GET("/skills", ctrlContainer.Skill.List)
	// sys.GET("/skills/:id", ctrlContainer.Skill.Get)
	// sys.PUT("/skills/:id", ctrlContainer.Skill.Update)
	// sys.DELETE("/skills/:id", ctrlContainer.Skill.Delete)

	// TODO
	sys.POST("/lilies", ctrlContainer.Lily.New)
	// sys.GET("/lilies", ctrlContainer.Lily.List)
	// sys.GET("/lilies/:id", ctrlContainer.Lily.Get)
	// sys.PUT("/lilies/:id", ctrlContainer.Lily.Update)
	// sys.PUT("/lilies/:id/charms", ctrlContainer.Lily.UpdateCharms)
	// sys.PUT("/lilies/:id/skills", ctrlContainer.Lily.UpdateSkills)
	// sys.PUT("/lilies/:lily_id/legions/:id", ctrlContainer.Lily.ApplyLegion)
	// sys.DELETE("/lilies/:id", ctrlContainer.Lily.Delete)

	// TODO
	// sys.POST("/legions", ctrlContainer.Legion.New)
	// sys.GET("/legions", ctrlContainer.Legion.List)
	// sys.GET("/legions/:id", ctrlContainer.Legion.Get)
	// sys.PUT("/legions/:id", ctrlContainer.Legion.Update)
	// sys.PUT("/legions/:id/lilies/:lily_id", ctrlContainer.Legion.ApproveLily)
	// sys.DELETE("/legions/:id", ctrlContainer.Legion.Delete)

	sys.POST("/gardens/:id/alarm", ctrlContainer.Protection.Alarm)
}
