package controller

import (
	"net/http"
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
		Name:         request.Name,
		Location:     request.Location,
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
