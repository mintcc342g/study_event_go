package controller

import (
	"net/http"
	"strings"
	"study-event-go/application"
	"study-event-go/application/dto"
	"study-event-go/types"

	"github.com/labstack/echo/v4"
)

// GardenController ...
type GardenController struct {
	gardenSvc *application.GardenService
}

// NewGardenController ...
func NewGardenController(gardenSvc *application.GardenService) *GardenController {
	return &GardenController{
		gardenSvc: gardenSvc,
	}
}

// New ...
func (g *GardenController) New(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		Name         string             `json:"name"`
		Location     string             `json:"location"`
		MentorshipID types.MentorshipID `json:"mentorship_id"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("GardenController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	gardenDTO := &dto.Garden{
		Name:         strings.TrimSpace(strings.ToLower(request.Name)),
		Location:     strings.TrimSpace(strings.ToLower(request.Location)),
		MentorshipID: request.MentorshipID,
	}

	gardenDTO, err = g.gardenSvc.New(ctx, gardenDTO)
	if err != nil {
		c.Logger().Error("GardenController New", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "New Garden OK", gardenDTO)
}

// Get ...
func (g *GardenController) Get(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		ID types.GardenID `param:"id"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("GardenController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	gardenDTO, err := g.gardenSvc.Get(ctx, request.ID)
	if err != nil {
		c.Logger().Error("GardenController Get", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "Get Garden OK", gardenDTO)
}

// List ...
func (g *GardenController) List(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		Offset uint32 `query:"offset"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("GardenController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	gardenDTO, err := g.gardenSvc.List(ctx, request.Offset)
	if err != nil {
		c.Logger().Error("GardenController List", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "List Gardens OK", gardenDTO)
}

// Update ...
func (g *GardenController) Update(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		ID           types.GardenID     `param:"id"`
		Name         string             `json:"name"`
		Location     string             `json:"location"`
		MentorshipID types.MentorshipID `json:"mentorship_id"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("GardenController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	gardenDTO := &dto.Garden{
		ID:           request.ID,
		Name:         strings.TrimSpace(strings.ToLower(request.Name)),
		Location:     strings.TrimSpace(strings.ToLower(request.Location)),
		MentorshipID: request.MentorshipID,
	}

	gardenDTO, err = g.gardenSvc.Update(ctx, gardenDTO)
	if err != nil {
		c.Logger().Error("GardenController Update", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "Update Garden OK", gardenDTO)
}
