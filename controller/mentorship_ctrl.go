package controller

import (
	"net/http"
	"strings"
	"study-event-go/application"
	"study-event-go/application/dto"
	"study-event-go/types"

	"github.com/labstack/echo/v4"
)

// MentorshipController ...
type MentorshipController struct {
	mentorshipSvc *application.MentorshipService
}

// NewMentorshipController ...
func NewMentorshipController(mentorshipSvc *application.MentorshipService) *MentorshipController {
	return &MentorshipController{
		mentorshipSvc: mentorshipSvc,
	}
}

// New ...
func (m *MentorshipController) New(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		Name string `json:"name"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("MentorshipController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	mentorshipDTO := &dto.Mentorship{
		Name: strings.TrimSpace(strings.ToLower(request.Name)),
	}

	mentorshipDTO, err = m.mentorshipSvc.New(ctx, mentorshipDTO)
	if err != nil {
		c.Logger().Error("MentorshipController New", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "New Mentorship OK", mentorshipDTO)
}

// Get ...
func (m *MentorshipController) Get(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		ID types.MentorshipID `param:"id"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("MentorshipController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	mentorshipDTO, err := m.mentorshipSvc.Get(ctx, request.ID)
	if err != nil {
		c.Logger().Error("MentorshipController Get", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "Get Mentorship OK", mentorshipDTO)
}

// List ...
func (m *MentorshipController) List(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		Offset uint32 `query:"offset"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("MentorshipController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	mentorshipDTO, err := m.mentorshipSvc.List(ctx, request.Offset)
	if err != nil {
		c.Logger().Error("MentorshipController List", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "List Mentorships OK", mentorshipDTO)
}

// Update ...
func (m *MentorshipController) Update(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		ID   types.MentorshipID `param:"id"`
		Name string             `json:"name"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("MentorshipController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	mentorshipDTO := &dto.Mentorship{
		ID:   request.ID,
		Name: strings.TrimSpace(strings.ToLower(request.Name)),
	}

	mentorshipDTO, err = m.mentorshipSvc.Update(ctx, mentorshipDTO)
	if err != nil {
		c.Logger().Error("MentorshipController Update", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "Update Mentorship OK", mentorshipDTO)
}

// Delete ...
func (m *MentorshipController) Delete(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		ID types.MentorshipID `param:"id"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("MentorshipController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	mentorshipDTO, err := m.mentorshipSvc.SoftDelete(ctx, request.ID)
	if err != nil {
		c.Logger().Error("MentorshipController Delete", "err", err)
		// TODO: error handle
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "Delete Mentorship OK", mentorshipDTO)
}
