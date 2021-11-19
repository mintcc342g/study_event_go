package controller

import (
	"net/http"
	"study_event_go/service"

	"github.com/labstack/echo/v4"
)

// ProtectionController ...
type ProtectionController struct {
	battleSvc *service.BattleService
}

// NewProtectionController ...
func NewProtectionController(battleSvc *service.BattleService) *ProtectionController {
	return &ProtectionController{
		battleSvc: battleSvc,
	}
}

func (p *ProtectionController) Warning(c echo.Context) (err error) {
	return response(c, http.StatusOK, "Warning OK", nil)
}
