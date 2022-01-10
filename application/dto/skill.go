package dto

import (
	"study-event-go/types"
	"time"
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
