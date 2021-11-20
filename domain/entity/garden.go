package entity

// Garden ...
type Garden struct {
	id               uint64
	name             string
	location         *Location
	mentorshipSystem *MentorshipSystem
}

// Location ...
type Location struct {
	id      uint64
	country string
	city    string
}
