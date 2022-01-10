package entity

import (
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

// NewMentorship ...
func NewMentorship(mentorshipDTO *dto.Mentorship) (*Mentorship, error) {
	if mentorshipDTO.Name == "" {
		return nil, errors.BadRequestf("invalid name")
	}

	return &Mentorship{
		Name: mentorshipDTO.Name,
	}, nil
}

// isDuplicatedName ...
func (m *Mentorship) isDuplicatedName(comparable *Mentorship) error {
	if comparable != nil && m.ID != comparable.ID && m.Name == comparable.Name {
		return errors.AlreadyExistsf("The name [%s]", comparable.Name)
	}

	return nil
}

// Update ...
func (m *Mentorship) Update(mentorshipDTO *dto.Mentorship, comparable *Mentorship) (err error) {
	if err = m.isDuplicatedName(comparable); err != nil {
		return
	}

	if mentorshipDTO.Name != "" {
		m.Name = mentorshipDTO.Name
	}

	return
}

// Delete ...
func (m *Mentorship) Delete() {
	now := time.Now().UTC()
	m.DeletedAt = &now
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
