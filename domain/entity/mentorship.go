package entity

import "time"

// MentorshipSystem ...
type MentorshipSystem struct {
	id   uint64
	name string
}

// Mentorship ...
type Mentorship struct {
	garden    *Garden
	mentor    *Lily
	mentee    *Lily
	createdAt time.Time
	deletedAt time.Time
}
