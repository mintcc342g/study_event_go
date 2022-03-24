package entity

import (
	"study-event-go/application/dto"
	"study-event-go/types"
	"time"

	"github.com/juju/errors"
)

// Garden ...
type Garden struct {
	ID           types.GardenID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string
	Location     string
	MentorshipID types.MentorshipID
	LegionSystem types.LegionSystem
	Legions      []*Legion
	Lilies       []*Lily
}

// NewGarden ...
func NewGarden(gardenDTO *dto.Garden) (*Garden, error) {
	if gardenDTO.Name == "" {
		return nil, errors.BadRequestf("invalid name")
	}

	if gardenDTO.Location == "" {
		return nil, errors.BadRequestf("invalid location")
	}

	if gardenDTO.LegionSystem == types.NoneLegionSystem {
		return nil, errors.BadRequestf("invalid legion system")
	}

	return &Garden{
		Name:         gardenDTO.Name,
		Location:     gardenDTO.Location,
		MentorshipID: gardenDTO.MentorshipID,
		LegionSystem: gardenDTO.LegionSystem,
	}, nil
}

// NewTopLegion ...
func (g *Garden) NewTopLegion(lilies []*Lily) (*Legion, error) {
	switch g.Name {
	case types.LudovicMissionSchool:
		return g.newTempleLegion(lilies)
	default:
		return nil, errors.BadRequestf("invalid garden")
	}
}

func (g *Garden) newTempleLegion(lilies []*Lily) (*Legion, error) {
	legion := &Legion{
		Members: lilies,
	}

	for _, lily := range lilies {
		if lily.IsFirstPlace() {
			legion.LeaderID = lily.ID
		}
	}

	return legion, nil
}

// EqualID ...
func (g *Garden) EqualID(id uint64) bool {
	return g.ID == types.GardenID(id)
}

// isDuplicatedName ...
func (g *Garden) isDuplicatedName(comparable *Garden) error {
	if comparable != nil && g.ID != comparable.ID && g.Name == comparable.Name {
		return errors.AlreadyExistsf("The name [%s]", comparable.Name)
	}

	return nil
}

// IsTopLegionSystem ...
func (g *Garden) IsTopLegionSystem() bool {
	return g.LegionSystem == types.TopLegionSystem
}

// IsAutonomicLegionSystem ...
func (g *Garden) IsAutonomicLegionSystem() bool {
	return g.LegionSystem == types.AutonomicLegionSystem
}

// Update ...
func (g *Garden) Update(req *dto.Garden, comparable *Garden) (err error) {
	if err = g.isDuplicatedName(comparable); err != nil {
		return
	}

	if req.Name != "" {
		g.Name = req.Name
	}

	if req.Location != "" {
		g.Location = req.Location
	}

	if req.MentorshipID != 0 {
		g.MentorshipID = req.MentorshipID
	}

	if req.LegionSystem != types.NoneLegionSystem {
		g.LegionSystem = req.LegionSystem
	}

	return
}

// Delete ...
func (g *Garden) Delete() {
	now := time.Now().UTC()
	g.DeletedAt = &now
}

// DTO ...
func (g *Garden) DTO() *dto.Garden {
	return &dto.Garden{
		ID:           g.ID,
		CreatedAt:    g.CreatedAt,
		UpdatedAt:    g.UpdatedAt,
		DeletedAt:    g.DeletedAt,
		Name:         g.Name,
		Location:     g.Location,
		MentorshipID: g.MentorshipID,
	}
}
