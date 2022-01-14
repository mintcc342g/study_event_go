package entity

import (
	"study-event-go/application/dto"
	"study-event-go/types"
	"time"

	"github.com/juju/errors"
)

// Charm ...
type Charm struct {
	ID        types.CharmID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	ModelID   types.CharmModelID
	OwnerID   types.LilyID
}

// CharmModel ...
type CharmModel struct {
	ID         types.CharmModelID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	CreatorID  types.CharmCreatorID
	Name       string
	Type       types.CharmModelType
	Generation types.CharmModelGeneration
}

// DTO ...
func (c *CharmModel) DTO() *dto.CharmModel {
	return &dto.CharmModel{
		ID:         c.ID,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
		DeletedAt:  c.DeletedAt,
		Name:       c.Name,
		Type:       c.Type,
		Generation: c.Generation,
		CreatorID:  c.CreatorID,
	}
}

// CharmCreator ...
type CharmCreator struct {
	ID        types.CharmCreatorID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Type      types.ArsenalType
}

// NewCharmCreator ...
func NewCharmCreator(charmCreatorDTO *dto.CharmCreator) (*CharmCreator, error) {
	if charmCreatorDTO.Name == "" {
		return nil, errors.BadRequestf("invalid creator name")
	}

	return &CharmCreator{
		Name: charmCreatorDTO.Name,
		Type: charmCreatorDTO.Type,
	}, nil
}

// NewCharmModel ...
func (c *CharmCreator) NewCharmModel(charmModelDTO *dto.CharmModel) (*CharmModel, error) {
	if charmModelDTO.Name == "" {
		return nil, errors.BadRequestf("invalid model name")
	}

	if charmModelDTO.CreatorID == 0 || c.ID != charmModelDTO.CreatorID {
		return nil, errors.BadRequestf("invalid creator id")
	}

	if charmModelDTO.Type == types.NoneTypeModel {
		return nil, errors.BadRequestf("invalid model type")
	}

	if charmModelDTO.Generation == types.NoneGeneration {
		return nil, errors.BadRequestf("invalid model generation")
	}

	return &CharmModel{
		Name:       charmModelDTO.Name,
		Type:       charmModelDTO.Type,
		Generation: charmModelDTO.Generation,
		CreatorID:  c.ID,
	}, nil
}

// DTO ...
func (c *CharmCreator) DTO() *dto.CharmCreator {
	return &dto.CharmCreator{
		ID:   c.ID,
		Name: c.Name,
		Type: c.Type,
	}
}
