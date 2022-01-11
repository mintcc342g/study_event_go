package application

import (
	"context"
	"study-event-go/application/dto"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"

	"github.com/juju/errors"
)

// SkillService ...
type SkillService struct {
	skillRepo interfaces.SkillRepository
}

// NewSkillService ...
func NewSkillService(
	skillRepo interfaces.SkillRepository,
) *SkillService {
	return &SkillService{
		skillRepo: skillRepo,
	}
}

// TODO: logger

// New ...
func (g *SkillService) New(ctx context.Context, skillDTO *dto.Skill) (*dto.Skill, error) {

	_, err := g.skillRepo.SkillByName(ctx, skillDTO.Name)
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	skill, err := entity.NewSkill(skillDTO)
	if err != nil {
		return nil, err
	}

	if skill, err = g.skillRepo.New(ctx, skill); err != nil {
		return nil, err
	}

	return skill.DTO(), nil
}
