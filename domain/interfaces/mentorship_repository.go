package interfaces

import (
	"context"
	"study_event_go/domain/entity"
	"study_event_go/types"
)

// MentorshipRepository ...
type MentorshipRepository interface {
	New(ctx context.Context, mentorship *entity.MentorshipSystem) (*entity.MentorshipSystem, error)
	Get(ctx context.Context, id types.MentorshipSystemID) (*entity.MentorshipSystem, error)
}
