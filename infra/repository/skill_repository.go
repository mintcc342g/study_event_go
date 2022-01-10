package repository

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/ent"
	entSkill "study-event-go/ent/skill"

	"github.com/juju/errors"
)

type skillRepository struct {
	conn *ent.Client
}

// NewSkillRepository ...
func NewSkillRepository(conn *ent.Client) interfaces.SkillRepository {
	return &skillRepository{
		conn: conn,
	}
}

func (s *skillRepository) New(ctx context.Context, skill *entity.Skill) (*entity.Skill, error) {

	id, err := s.conn.Skill.
		Create().
		SetName(skill.Name).
		SetType(skill.Type).
		OnConflict().
		UpdateUpdatedAt().
		ClearDeletedAt().
		ID(ctx)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	entModel, err := s.conn.Skill.Get(ctx, id)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	skill.ID = entModel.ID
	skill.CreatedAt = entModel.CreatedAt
	skill.UpdatedAt = entModel.UpdatedAt
	skill.DeletedAt = entModel.DeletedAt
	skill.Name = entModel.Name
	skill.Type = entModel.Type

	return skill, nil
}

func (s *skillRepository) GetByName(ctx context.Context, name string) (*entity.Skill, error) {

	entModel, err := s.conn.Skill.
		Query().
		Where(
			entSkill.Name(name),
			entSkill.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		// logger
		if ent.IsNotFound(err) { // TODO: converter (ent error -> error)
			return nil, errors.NotFoundf("The name[%s]", name)
		}
		return nil, errors.New("internal server error")
	}

	return &entity.Skill{
		ID:        entModel.ID,
		CreatedAt: entModel.CreatedAt,
		UpdatedAt: entModel.UpdatedAt,
		DeletedAt: entModel.DeletedAt,
		Name:      entModel.Name,
		Type:      entModel.Type,
	}, nil
}
