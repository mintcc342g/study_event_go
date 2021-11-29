package entity

import (
	"study_event_go/domain/vo"
)

// Lily ...
type Lily struct {
	id        uint64
	name      *vo.Name
	rareSkill *Skill
	subSkill  *Skill
	mainCharm *Charm
	subCharm  *Charm
	gardenID  uint64
	rank      uint32
	legionID  uint64
}

// Skill ...
type Skill struct {
	id   uint64
	name string
}

// Charm ...
type Charm struct {
	id      uint64
	name    string
	creator string
}

// Legion ...
type Legion struct {
	ID       uint64
	GardenID uint64
	Name     string
	LeaderID uint64
	Members  []*Lily
}
