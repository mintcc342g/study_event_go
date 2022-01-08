package entity

import (
	"strings"
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
func NewGarden(req *dto.Garden) (*Garden, error) {
	if err := validateGardenDTO(req); err != nil {
		return nil, err
	}

	return &Garden{
		Name:         req.Name,
		Location:     req.Location,
		MentorshipID: req.MentorshipID,
	}, nil
}

func validateGardenDTO(req *dto.Garden) error {
	req.Name = strings.TrimSpace(strings.ToLower(req.Name))
	if req.Name == "" {
		return errors.BadRequestf("invalid name")
	}

	req.Location = strings.TrimSpace(strings.ToLower(req.Location))
	if req.Location == "" {
		return errors.BadRequestf("invalid location")
	}

	return nil
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
