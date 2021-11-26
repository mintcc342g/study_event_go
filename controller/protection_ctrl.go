package controller

import (
	"net/http"
	"study_event_go/application"

	"github.com/labstack/echo/v4"
)

// ProtectionController ...
type ProtectionController struct {
	battleSvc *application.BattleService
}

// NewProtectionController ...
func NewProtectionController(battleSvc *application.BattleService) *ProtectionController {
	return &ProtectionController{
		battleSvc: battleSvc,
	}
}

// Warning ...
func (p *ProtectionController) Warning(c echo.Context) (err error) {
	return response(c, http.StatusOK, "Warning OK", nil)
}
