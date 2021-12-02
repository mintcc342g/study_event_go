package entity

import (
	"study_event_go/types"

	"github.com/juju/errors"
)

// Garden ...
type Garden struct {
	ID               uint64
	Name             string
	Location         string
	MentorshipSystem *MentorshipSystem
	Legions          []*Legion
	Lilies           []*Lily
}

// NewTempleLegion ...
func (g *Garden) NewTempleLegion(lilies []*Lily) (*Legion, error) {

	if !g.IsLudovico() {
		return nil, errors.BadRequestf("invalid garden")
	}

	legion := &Legion{
		Members: lilies,
	}

	for _, lily := range lilies {
		if lily.rank == types.TempleLegionLeaderLank {
			legion.LeaderID = lily.id
		}
	}

	return legion, nil
}

// EqualID ...
func (g *Garden) EqualID(id uint64) bool {
	return g.ID == id
}

// IsLudovico ...
func (g *Garden) IsLudovico() bool {
	return g.Name == types.LudovicoMissionSchool
}

// // Location ...
// type Location struct {
// 	id      uint64
// 	country string
// 	city    string
// }
