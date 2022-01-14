package interfaces

import (
	"context"
	"study-event-go/domain/entity"
)

// SkillRepository ...
type SkillRepository interface {
	Save(ctx context.Context, skill *entity.Skill) (*entity.Skill, error)
	SkillByName(ctx context.Context, name string) (*entity.Skill, error)
}
