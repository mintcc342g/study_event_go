package entity

import (
	"study_event_go/types"
	"time"
)

// MentorshipSystem ...
type MentorshipSystem struct {
	ID   types.MentorshipSystemID
	Name string
}

// Mentorship ...
type Mentorship struct {
	garden    *Garden
	mentor    *Lily
	mentee    *Lily
	createdAt time.Time
	deletedAt time.Time
}
