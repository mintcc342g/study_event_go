package entity

import (
	"study-event-go/domain/vo"
	"study-event-go/types"
	"time"
)

// Lily ...
type Lily struct {
	ID             types.LilyID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	DeletionReason types.DeletionReason
	Name           *vo.Name
	Rank           uint32
	MainCharm      *Charm
	SubCharm       *Charm
	GardenID       types.GardenID
	LegionID       types.LegionID
	Skills         []*Skill
}

// IsFirstPlace ...
func (l *Lily) IsFirstPlace() bool {
	return l.Rank == types.FirstPlace
}

// LilySkill ...
type LilySkill struct {
	LilyID    types.LilyID
	SkillID   types.SkillID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Charm ...
type Charm struct {
	id      types.CharmID
	name    string
	creator string
}

// Legion ...
type Legion struct {
	ID       types.LegionID
	GardenID types.GardenID
	Name     string
	LeaderID types.LilyID
	Members  []*Lily
}
