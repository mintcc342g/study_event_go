package entity

import (
	"strings"
	"study-event-go/application/dto"
	"study-event-go/types"
	"time"

	"github.com/juju/errors"
)

// Mentorship ...
type Mentorship struct {
	ID        types.MentorshipID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
}

// GardenMentorship ...
type GardenMentorship struct {
	garden     *Garden
	mentorship *Mentorship
	createdAt  time.Time
	updatedAt  time.Time
	deletedAt  time.Time
	mentor     *Lily
	mentee     *Lily
}

// NewMentorship ...
func NewMentorship(req *dto.Mentorship) (*Mentorship, error) {
	if err := validateMentorshipDTO(req); err != nil {
		return nil, err
	}

	return &Mentorship{
		Name: req.Name,
	}, nil
}

// Update ...
func (m *Mentorship) Update(req *dto.Mentorship) error {
	if err := validateMentorshipDTO(req); err != nil {
		return err
	}

	m.Name = req.Name

	return nil
}

// Delete ...
func (m *Mentorship) Delete() {
	now := time.Now().UTC()
	m.DeletedAt = &now
}

func validateMentorshipDTO(req *dto.Mentorship) error {
	req.Name = strings.TrimSpace(strings.ToLower(req.Name))
	if req.Name == "" {
		return errors.BadRequestf("invalid name")
	}

	return nil
}

// DTO ...
func (m *Mentorship) DTO() *dto.Mentorship {
	return &dto.Mentorship{
		ID:        m.ID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: m.DeletedAt,
		Name:      m.Name,
	}
}
