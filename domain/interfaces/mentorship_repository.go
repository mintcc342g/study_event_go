package interfaces

import (
	"context"
	"study_event_go/domain/entity"
)

// MentorshipRepository ...
type MentorshipRepository interface {
	New(ctx context.Context, mentorship *entity.MentorshipSystem) (*entity.MentorshipSystem, error)
}
