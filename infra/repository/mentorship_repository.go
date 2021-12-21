package repository

import (
	"context"
	"study_event_go/domain/entity"
	"study_event_go/domain/interfaces"
	"study_event_go/ent"
)

type mentorshipRepository struct {
	conn *ent.Client
}

// NewMentorshipRepository ...
func NewMentorshipRepository(conn *ent.Client) interfaces.MentorshipRepository {
	return &mentorshipRepository{
		conn: conn,
	}
}

// TODO: logger

func (m *mentorshipRepository) New(ctx context.Context, mentorship *entity.MentorshipSystem) (*entity.MentorshipSystem, error) {

	id, err := m.conn.MentorshipSystem.
		Create().
		SetName(mentorship.Name).
		SetNillableDeletedAt(nil).
		OnConflict().
		UpdateUpdatedAt().
		UpdateDeletedAt().
		ID(ctx)
	if err != nil {
		return nil, err
	}

	entMode, err := m.conn.MentorshipSystem.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	mentorship.ID = entMode.ID
	mentorship.Name = entMode.Name
	mentorship.CreatedAt = entMode.CreatedAt
	mentorship.UpdatedAt = entMode.UpdatedAt

	return mentorship, nil
}
