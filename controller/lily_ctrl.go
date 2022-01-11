package controller

import (
	"net/http"
	"strings"
	"study-event-go/application"
	"study-event-go/application/dto"
	"study-event-go/types"
	"time"

	"github.com/labstack/echo/v4"
)

// LilyController ...
type LilyController struct {
	lilySvc *application.LilyService
}

// NewLilyController ...
func NewLilyController(lilySvc *application.LilyService) *LilyController {
	return &LilyController{
		lilySvc: lilySvc,
	}
}

// New ...
func (l *LilyController) New(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		FirstName  string         `json:"first_name"`
		MiddleName string         `json:"middle_name"`
		LastName   string         `json:"last_name"`
		Rank       uint32         `json:"rank"`
		Birth      time.Time      `json:"birth"`
		GardenID   types.GardenID `json:"garden_id"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("LilyController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	lilyDTO := &dto.Lily{
		FirstName:  strings.TrimSpace(strings.ToLower(request.FirstName)),
		MiddleName: strings.TrimSpace(strings.ToLower(request.MiddleName)),
		LastName:   strings.TrimSpace(strings.ToLower(request.LastName)),
		Rank:       request.Rank,
		Birth:      request.Birth,
		GardenID:   request.GardenID,
	}

	lilyDTO, err = l.lilySvc.New(ctx, lilyDTO)
	if err != nil {
		c.Logger().Error("LilyController New", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "New Lily OK", lilyDTO)
}

// NewCharms ...
func (l *LilyController) NewCharms(c echo.Context) error {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		LilyID        types.LilyID         `param:"id"`
		CharmModelIDs []types.CharmModelID `json:"charm_model_ids"`
	}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error("LilyController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	lilyDTO, err := l.lilySvc.NewCharms(ctx, request.LilyID, request.CharmModelIDs)
	if err != nil {
		c.Logger().Error("LilyController NewCharms", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "New Lily's Charms OK", lilyDTO)
}
