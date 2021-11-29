package controller

import (
	"net/http"
	"study_event_go/application"
	"study_event_go/application/dto"
	"study_event_go/types"

	"github.com/jinzhu/copier"
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
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		CaveLocation string `json:"caveLocation"`
		Huges        []struct {
			HugeClass types.HugeClass `json:"hugeClass"`
			HugeType  types.HugeType  `json:"hugeType"`
		} `json:"huges"`
		TotalCount uint32           `json:"totalCount"`
		AlertLevel types.AlertLevel `json:"alertLevel"`
	}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error("ProtectionController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	warningDTO := &dto.Warning{}
	if err := copier.Copy(warningDTO, request); err != nil {
		c.Logger().Error("ProtectionController Copy", "err", err)
		return response(c, http.StatusInternalServerError, "internal server error", nil)
	}

	if err := p.battleSvc.Warning(ctx, warningDTO); err != nil {
		c.Logger().Error("ProtectionController Warning", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "Warning OK", nil)
}
