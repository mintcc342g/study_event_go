package controller

import (
	"net/http"
	"study-event-go/application"
	"study-event-go/application/dto"
	"study-event-go/types"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

// ProtectionController ...
type ProtectionController struct {
	protectionSvc *application.ProtectionService
}

// NewProtectionController ...
func NewProtectionController(protectionSvc *application.ProtectionService) *ProtectionController {
	return &ProtectionController{
		protectionSvc: protectionSvc,
	}
}

// Alarm ...
func (p *ProtectionController) Alarm(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		GardenID     types.GardenID `param:"id"`
		CaveLocation string         `json:"cave_location"`
		Huges        []struct {
			HugeClass types.HugeClass `json:"huge_class"`
			HugeType  types.HugeType  `json:"huge_type"`
		} `json:"huges"`
		TotalHugeCount    uint32           `json:"total_huge_count"`
		AlertLevel        types.AlertLevel `json:"alert_level"`
		LegionMemberCount uint32           `json:"legion_member_count"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("ProtectionController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	alarmingDTO := &dto.Alarm{}
	if err = copier.Copy(alarmingDTO, request); err != nil {
		c.Logger().Error("ProtectionController Copy", "err", err)
		return response(c, http.StatusInternalServerError, "internal server error", nil)
	}

	if err = p.protectionSvc.Alarm(ctx, alarmingDTO); err != nil {
		c.Logger().Error("ProtectionController Alarm", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "Alarm OK", nil)
}
