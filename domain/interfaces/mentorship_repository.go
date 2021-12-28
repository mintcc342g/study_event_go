package interfaces

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/types"
)

// MentorshipRepository ...
type MentorshipRepository interface {
	New(ctx context.Context, mentorship *entity.MentorshipSystem) (*entity.MentorshipSystem, error)
	Get(ctx context.Context, id types.MentorshipSystemID) (*entity.MentorshipSystem, error)
}
