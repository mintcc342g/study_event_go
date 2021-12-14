package entity

import (
	"study_event_go/domain/vo"
	"study_event_go/types"
)

// Lily ...
type Lily struct {
	id        types.LilyID
	name      *vo.Name
	rareSkill *Skill
	subSkill  *Skill
	mainCharm *Charm
	subCharm  *Charm
	gardenID  types.GardenID
	rank      uint32
	legionID  types.LegionID
}

// IsFirstPlace ...
func (l *Lily) IsFirstPlace() bool {
	return l.rank == types.FirstPlace
}

// Skill ...
type Skill struct {
	id   types.SkillID
	name string
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
