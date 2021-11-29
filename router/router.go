package router

import (
	"study_event_go/cmd/container"

	"github.com/labstack/echo/v4"
)

// GardenRoute ...
func GardenRoute(sys *echo.Group, ctrlContainer *container.ControllerContainer) {
	sys.POST("/:id/alarm", ctrlContainer.Protection.Alarm)
}
