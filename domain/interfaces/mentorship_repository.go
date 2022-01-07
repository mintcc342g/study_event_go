package interfaces

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/types"
)

// MentorshipRepository ...
type MentorshipRepository interface {
	New(ctx context.Context, mentorship *entity.Mentorship) (*entity.Mentorship, error)
	Get(ctx context.Context, id types.MentorshipID) (*entity.Mentorship, error)
	GetByName(ctx context.Context, name string) (*entity.Mentorship, error)
	List(ctx context.Context, offset uint32) ([]*entity.Mentorship, error)
	Update(ctx context.Context, mentorship *entity.Mentorship) (*entity.Mentorship, error)
}
