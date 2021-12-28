package entity

import (
	"study-event-go/types"

	"github.com/juju/errors"
)

// Garden ...
type Garden struct {
	ID               types.GardenID
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
		if lily.IsFirstPlace() {
			legion.LeaderID = lily.id
		}
	}

	return legion, nil
}

// EqualID ...
func (g *Garden) EqualID(id uint64) bool {
	return g.ID == types.GardenID(id)
}

// IsLudovico ...
func (g *Garden) IsLudovico() bool {
	return g.Name == types.LudovicoMissionSchool
}
