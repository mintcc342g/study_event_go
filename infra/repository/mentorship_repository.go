package repository

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/ent"
	entMentorship "study-event-go/ent/mentorship"
	"study-event-go/types"

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

func (m *mentorshipRepository) New(ctx context.Context, mentorship *entity.Mentorship) (*entity.Mentorship, error) {

	id, err := m.conn.Mentorship.
		Create().
		SetName(mentorship.Name).
		OnConflict().
		UpdateUpdatedAt().
		ClearDeletedAt().
		ID(ctx)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	entModel, err := m.conn.Mentorship.Get(ctx, id)
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

func (m *mentorshipRepository) Mentorship(ctx context.Context, id types.MentorshipID) (*entity.Mentorship, error) {

	entModel, err := m.conn.Mentorship.
		Query().
		Where(
			entMentorship.ID(id),
			entMentorship.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		// logger
		return nil, errors.NotFoundf("id[%d]", id)
	}

	return &entity.Mentorship{
		ID:        entModel.ID,
		CreatedAt: entModel.CreatedAt,
		UpdatedAt: entModel.UpdatedAt,
		DeletedAt: entModel.DeletedAt,
		Name:      entModel.Name,
	}, nil
}

func (m *mentorshipRepository) MentorshipByName(ctx context.Context, name string) (*entity.Mentorship, error) {

	entModel, err := m.conn.Mentorship.
		Query().
		Where(
			entMentorship.Name(name),
			entMentorship.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		// logger
		if ent.IsNotFound(err) { // TODO: converter (ent error -> error)
			return nil, errors.NotFoundf("The name[%s]", name)
		}
		return nil, errors.New("internal server error")
	}

	return &entity.Mentorship{
		ID:        entModel.ID,
		CreatedAt: entModel.CreatedAt,
		UpdatedAt: entModel.UpdatedAt,
		DeletedAt: entModel.DeletedAt,
		Name:      entModel.Name,
	}, nil
}

func (m *mentorshipRepository) Mentorships(ctx context.Context, offset uint32) ([]*entity.Mentorship, error) {

	entModels, err := m.conn.Mentorship.
		Query().
		Where(
			entMentorship.DeletedAtIsNil()).
		Limit(10).
		Offset(int(offset)).
		All(ctx)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	mentorships := make([]*entity.Mentorship, len(entModels))
	for i, v := range entModels {
		mentorships[i] = &entity.Mentorship{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
			Name:      v.Name,
		}
	}

	return mentorships, nil
}

func (m *mentorshipRepository) Update(ctx context.Context, mentorship *entity.Mentorship) (*entity.Mentorship, error) {

	entModel, err := m.conn.Mentorship.
		UpdateOneID(mentorship.ID).
		SetName(mentorship.Name).
		SetNillableDeletedAt(mentorship.DeletedAt).
		Save(ctx)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	mentorship.ID = entModel.ID
	mentorship.Name = entModel.Name
	mentorship.CreatedAt = entModel.CreatedAt
	mentorship.UpdatedAt = entModel.UpdatedAt
	mentorship.DeletedAt = entModel.DeletedAt

	return mentorship, nil
}
