package entity

import (
	"strings"
	"study-event-go/application/dto"
	"study-event-go/types"
	"time"

	"github.com/juju/errors"
)

// MentorshipSystem ...
type MentorshipSystem struct {
	ID        types.MentorshipSystemID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
}

// GardenMentorship ...
type GardenMentorship struct {
	garden           *Garden
	mentorshipSystem *MentorshipSystem
	createdAt        time.Time
	updatedAt        time.Time
	deletedAt        time.Time
	mentor           *Lily
	mentee           *Lily
}

// NewMentorship ...
func NewMentorship(req *dto.Mentorship) (*MentorshipSystem, error) {
	if err := validateMentorshipDTO(req); err != nil {
		return nil, err
	}

	return &MentorshipSystem{
		Name: req.Name,
	}, nil
}

// Update ...
func (m *MentorshipSystem) Update(req *dto.Mentorship) error {
	if err := validateMentorshipDTO(req); err != nil {
		return err
	}

	m.Name = req.Name

	return nil
}

func validateMentorshipDTO(req *dto.Mentorship) error {
	req.Name = strings.TrimSpace(strings.ToLower(req.Name))
	if req.Name == "" {
		return errors.BadRequestf("invalid name")
	}

	return nil
}

// DTO ...
func (m *MentorshipSystem) DTO() *dto.Mentorship {
	return &dto.Mentorship{
		ID:        m.ID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: m.DeletedAt,
		Name:      m.Name,
	}
}
