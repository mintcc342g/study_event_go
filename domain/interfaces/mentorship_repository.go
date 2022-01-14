package interfaces

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/types"
)

// MentorshipRepository ...
type MentorshipRepository interface {
	Save(ctx context.Context, mentorship *entity.Mentorship) (*entity.Mentorship, error)
	Mentorship(ctx context.Context, id types.MentorshipID) (*entity.Mentorship, error)
	MentorshipByName(ctx context.Context, name string) (*entity.Mentorship, error)
	Mentorships(ctx context.Context, offset uint32) ([]*entity.Mentorship, error)
	Update(ctx context.Context, mentorship *entity.Mentorship) (*entity.Mentorship, error)
}
