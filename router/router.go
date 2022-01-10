package router

import (
	"study-event-go/cmd/container"

	"github.com/labstack/echo/v4"
)

// AssaultLilyRoute ...
func AssaultLilyRoute(sys *echo.Group, ctrlContainer *container.ControllerContainer) {

	// TODO: check authority with jwt

	sys.POST("/mentorship", ctrlContainer.Mentorship.New)
	sys.GET("/mentorship/:id", ctrlContainer.Mentorship.Get)
	sys.GET("/mentorship", ctrlContainer.Mentorship.List)
	sys.PUT("/mentorship/:id", ctrlContainer.Mentorship.Update)
	sys.DELETE("/mentorship/:id", ctrlContainer.Mentorship.Delete)

	sys.POST("/garden", ctrlContainer.Garden.New)
	sys.GET("/garden/:id", ctrlContainer.Garden.Get)
	sys.GET("/garden", ctrlContainer.Garden.List)
	sys.PUT("/garden/:id", ctrlContainer.Garden.Update)
	sys.DELETE("/garden/:id", ctrlContainer.Garden.Delete)

	// TODO
	sys.POST("/skill", ctrlContainer.Skill.New)
	// sys.GET("/skill", ctrlContainer.Skill.List)
	// sys.GET("/skill/:id", ctrlContainer.Skill.Get)
	// sys.PUT("/skill/:id", ctrlContainer.Skill.Update)
	// sys.DELETE("/skill/:id", ctrlContainer.Skill.Delete)

	// TODO
	// sys.POST("/lily", ctrlContainer.Lily.New)
	// sys.GET("/lily", ctrlContainer.Lily.List)
	// sys.GET("/lily/:id", ctrlContainer.Lily.Get)
	// sys.PUT("/lily/:id", ctrlContainer.Lily.Update)
	// sys.DELETE("/lily/:id", ctrlContainer.Lily.Delete)

	sys.POST("/garden/:id/alarm", ctrlContainer.Protection.Alarm)
}
