package repository

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/ent"
	"study-event-go/ent/mentorshipsystem"
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

func (m *mentorshipRepository) GetByName(ctx context.Context, name string) (*entity.MentorshipSystem, error) {

	entModel, err := m.conn.MentorshipSystem.
		Query().
		Where(mentorshipsystem.Name(name)).
		Only(ctx)
	if err != nil {
		// logger
		if ent.IsNotFound(err) { // TODO: converter (ent error -> error)
			return nil, errors.NotFoundf("The name[%s]", name)
		}
		return nil, errors.New("internal server error")
	}

	return &entity.MentorshipSystem{
		ID:        entModel.ID,
		CreatedAt: entModel.CreatedAt,
		UpdatedAt: entModel.UpdatedAt,
		DeletedAt: entModel.DeletedAt,
		Name:      entModel.Name,
	}, nil
}

func (m *mentorshipRepository) List(ctx context.Context, offset uint32) ([]*entity.MentorshipSystem, error) {

	entModels, err := m.conn.MentorshipSystem.
		Query().
		Limit(10).
		Offset(int(offset)).
		All(ctx)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	mentorships := make([]*entity.MentorshipSystem, len(entModels))
	for i, v := range entModels {
		mentorships[i] = &entity.MentorshipSystem{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
			Name:      v.Name,
		}
	}

	return mentorships, nil
}

func (m *mentorshipRepository) Update(ctx context.Context, mentorship *entity.MentorshipSystem) (*entity.MentorshipSystem, error) {

	entModel, err := m.conn.MentorshipSystem.
		UpdateOneID(mentorship.ID).
		SetName(mentorship.Name).
		Save(ctx)
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
