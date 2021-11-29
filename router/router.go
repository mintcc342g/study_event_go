package router

import (
	"study_event_go/cmd/container"

	"github.com/labstack/echo/v4"
)

// Route ...
func Route(sys *echo.Group, ctrlContainer *container.ControllerContainer) {
	sys.POST("/alarm", ctrlContainer.Protection.Alarm)
}
