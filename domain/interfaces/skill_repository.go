package interfaces

import (
	"context"
	"study-event-go/domain/entity"
)

// SkillRepository ...
type SkillRepository interface {
	New(ctx context.Context, skill *entity.Skill) (*entity.Skill, error)
	GetByName(ctx context.Context, name string) (*entity.Skill, error)
}
