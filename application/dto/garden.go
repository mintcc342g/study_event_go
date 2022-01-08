package dto

import (
	"study-event-go/types"
	"time"
)

// Garden ...
type Garden struct {
	ID           types.GardenID     `json:"id,omitempty"`
	CreatedAt    time.Time          `json:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty"`
	DeletedAt    *time.Time         `json:"deleted_at,omitempty"`
	Name         string             `json:"name,omitempty"`
	Location     string             `json:"location,omitempty"`
	MentorshipID types.MentorshipID `json:"mentorship_id,omitempty"`
}
