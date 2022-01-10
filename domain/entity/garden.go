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

	return &Garden{
		Name:         gardenDTO.Name,
		Location:     gardenDTO.Location,
		MentorshipID: gardenDTO.MentorshipID,
	}, nil
}

// NewTempleLegion ...
func (g *Garden) NewTempleLegion(lilies []*Lily) (*Legion, error) {

	if !g.IsLudovic() {
		return nil, errors.BadRequestf("invalid garden")
	}

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

// IsLudovic ...
func (g *Garden) IsLudovic() bool {
	return g.Name == types.LudovicMissionSchool
}

// isDuplicatedName ...
func (g *Garden) isDuplicatedName(comparable *Garden) error {
	if comparable != nil && g.ID != comparable.ID && g.Name == comparable.Name {
		return errors.AlreadyExistsf("The name [%s]", comparable.Name)
	}

	return nil
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
