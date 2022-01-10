package entity

import (
	"study-event-go/application/dto"
	"study-event-go/types"
	"time"

	"github.com/juju/errors"
)

// Skill ...
type Skill struct {
	ID        types.SkillID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Type      types.SkillType
}

// NewSkill ...
func NewSkill(skillDTO *dto.Skill) (*Skill, error) {
	if skillDTO.Name == "" {
		return nil, errors.BadRequestf("invalid name")
	}

	if skillDTO.Type == types.NoneSkillType {
		return nil, errors.BadRequestf("invalid type")
	}

	return &Skill{
		Name: skillDTO.Name,
		Type: skillDTO.Type,
	}, nil
}

// DTO ...
func (s *Skill) DTO() *dto.Skill {
	return &dto.Skill{
		ID:        s.ID,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		DeletedAt: s.DeletedAt,
		Name:      s.Name,
		Type:      s.Type,
	}
}
