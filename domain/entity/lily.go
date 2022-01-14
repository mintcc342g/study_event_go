package entity

import (
	"study-event-go/application/dto"
	"study-event-go/domain/vo"
	"study-event-go/types"
	"time"

	"github.com/juju/errors"
)

// Lily ...
type Lily struct {
	ID              types.LilyID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
	CauseOfDeletion types.CauseOfDeletion
	Name            *vo.Name
	Birth           time.Time
	Rank            uint32
	GardenID        types.GardenID
	LegionID        types.LegionID
	Charms          []*Charm
	Skills          []*Skill
	Relations       []*LilysRelation
}

// NewLily ...
func NewLily(lilyDTO *dto.Lily) (*Lily, error) {
	if lilyDTO.FirstName == "" || lilyDTO.LastName == "" {
		return nil, errors.BadRequestf("invalid name")
	}

	if lilyDTO.GardenID == 0 {
		return nil, errors.BadRequestf("invalid garden id")
	}

	return &Lily{
		Name:     vo.NewName(lilyDTO.FirstName, lilyDTO.MiddleName, lilyDTO.LastName),
		Birth:    lilyDTO.Birth,
		Rank:     lilyDTO.Rank,
		GardenID: lilyDTO.GardenID,
	}, nil
}

// ContractWith ...
func (l *Lily) ContractWith(charmModelIDs []types.CharmModelID, models ...*CharmModel) error {
	if len(charmModelIDs) != len(models) {
		return errors.BadRequestf("invalid charm model ids")
	}

	for _, m := range models {
		l.Charms = append(l.Charms, &Charm{
			ModelID: m.ID,
			OwnerID: l.ID,
		})
	}

	return nil
}

// IsFirstPlace ...
func (l *Lily) IsFirstPlace() bool {
	return l.Rank == types.FirstPlace
}

// DTO ...
func (l *Lily) DTO() *dto.Lily {
	return &dto.Lily{
		ID:              l.ID,
		CreatedAt:       l.CreatedAt,
		UpdatedAt:       l.UpdatedAt,
		DeletedAt:       l.DeletedAt,
		CauseOfDeletion: l.CauseOfDeletion,
		FirstName:       l.Name.First,
		MiddleName:      l.Name.Middle,
		LastName:        l.Name.Last,
		Birth:           l.Birth,
		Rank:            l.Rank,
		GardenID:        l.GardenID,
		LegionID:        l.LegionID,
	}
}

// Ability ...
type Ability struct {
	LilyID    types.LilyID
	SkillID   types.SkillID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Legion ...
type Legion struct {
	ID       types.LegionID
	GardenID types.GardenID
	Name     string
	LeaderID types.LilyID
	Members  []*Lily
}

// LilysRelation ...
type LilysRelation struct {
	GardenID     types.GardenID
	MentorshipID types.MentorshipID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	MentorID     *Lily
	MenteeID     *Lily
}
