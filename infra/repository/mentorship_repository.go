package repository

import (
	"context"
	"study_event_go/domain/entity"
	"study_event_go/domain/interfaces"
	"study_event_go/ent"
	"study_event_go/types"

	"github.com/juju/errors"
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
		// logger
		return nil, errors.New("internal server error")
	}

	entModel, err := m.conn.MentorshipSystem.Get(ctx, id)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	mentorship.ID = entModel.ID
	mentorship.Name = entModel.Name
	mentorship.CreatedAt = entModel.CreatedAt
	mentorship.UpdatedAt = entModel.UpdatedAt

	return mentorship, nil
}

func (m *mentorshipRepository) Get(ctx context.Context, id types.MentorshipSystemID) (*entity.MentorshipSystem, error) {

	entModel, err := m.conn.MentorshipSystem.Get(ctx, id)
	if err != nil {
		// logger
		return nil, errors.NotFoundf("id[%d]", id)
	}

	return &entity.MentorshipSystem{
		ID:        entModel.ID,
		CreatedAt: entModel.CreatedAt,
		UpdatedAt: entModel.UpdatedAt,
		DeletedAt: entModel.DeletedAt,
		Name:      entModel.Name,
	}, nil
}
