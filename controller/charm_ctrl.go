package controller

import (
	"net/http"
	"strings"
	"study-event-go/application"
	"study-event-go/application/dto"
	"study-event-go/types"

	"github.com/labstack/echo/v4"
)

// CharmController ...
type CharmController struct {
	charmSvc *application.CharmService
}

// NewCharmController ...
func NewCharmController(charmSvc *application.CharmService) *CharmController {
	return &CharmController{
		charmSvc: charmSvc,
	}
}

// NewCharmCreator ...
func (h *CharmController) NewCharmCreator(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		Name string            `json:"name"`
		Type types.ArsenalType `json:"type"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("CharmController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	charmCreatorDTO := &dto.CharmCreator{
		Name: strings.TrimSpace(strings.ToLower(request.Name)),
		Type: request.Type,
	}

	charmCreatorDTO, err = h.charmSvc.NewCharmCreator(ctx, charmCreatorDTO)
	if err != nil {
		c.Logger().Error("CharmController NewCharmCreator", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "New CharmCreator OK", charmCreatorDTO)
}

// NewCharmModel ...
func (h *CharmController) NewCharmModel(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		CreatorID  types.CharmCreatorID       `json:"creator_id"`
		Name       string                     `json:"name"`
		Type       types.CharmModelType       `type:"type"`
		Generation types.CharmModelGeneration `json:"generation"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("CharmController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	charmModelDTO := &dto.CharmModel{
		CreatorID:  request.CreatorID,
		Name:       strings.TrimSpace(strings.ToLower(request.Name)),
		Type:       request.Type,
		Generation: request.Generation,
	}

	charmModelDTO, err = h.charmSvc.NewCharmModel(ctx, charmModelDTO)
	if err != nil {
		c.Logger().Error("CharmController NewCharmModel", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "New CharmModel OK", charmModelDTO)
}
