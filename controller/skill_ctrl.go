package controller

import (
	"net/http"
	"strings"
	"study-event-go/application"
	"study-event-go/application/dto"
	"study-event-go/types"

	"github.com/labstack/echo/v4"
)

// SkillController ...
type SkillController struct {
	skillSvc *application.SkillService
}

// NewSkillController ...
func NewSkillController(skillSvc *application.SkillService) *SkillController {
	return &SkillController{
		skillSvc: skillSvc,
	}
}

// New ...
func (s *SkillController) New(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		Name string          `json:"name"`
		Type types.SkillType `json:"type"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("SkillController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	skillDTO := &dto.Skill{
		Name: strings.TrimSpace(strings.ToLower(request.Name)),
		Type: request.Type,
	}

	skillDTO, err = s.skillSvc.New(ctx, skillDTO)
	if err != nil {
		c.Logger().Error("SkillController New", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "New Skill OK", skillDTO)
}
