package entity

import (
	"strings"
	"study-event-go/application/dto"
	"study-event-go/types"
	"time"
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
func NewMentorship(req *dto.Mentorship) *MentorshipSystem {
	return &MentorshipSystem{
		Name: strings.TrimSpace(strings.ToLower(req.Name)),
	}
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
