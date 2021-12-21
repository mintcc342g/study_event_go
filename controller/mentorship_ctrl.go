package controller

import (
	"net/http"
	"study_event_go/application"
	"study_event_go/application/dto"
	"study_event_go/types"

	"github.com/jinzhu/copier"
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

	mentorshipDTO := &dto.Mentorship{}
	if err = copier.Copy(mentorshipDTO, request); err != nil {
		c.Logger().Error("MentorshipController Copy", "err", err)
		return response(c, http.StatusInternalServerError, "internal server error", nil)
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
		ID types.MentorshipSystemID `param:"id"`
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
