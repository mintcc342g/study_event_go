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
	garden    *Garden
	grade     uint32
	rank      uint32
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
