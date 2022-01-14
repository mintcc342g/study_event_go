package dto

import (
	"study-event-go/types"
	"time"
)

// Charm ...
type Charm struct {
	ID        types.CharmID      `json:"id,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
	DeletedAt *time.Time         `json:"deleted_at,omitempty"`
	Name      string             `json:"name,omitempty"`
	ModelID   types.CharmModelID `json:"model_id,omitempty"`
	OwnerID   types.LilyID       `json:"owner_id,omitempty"`
}

// CharmModel ...
type CharmModel struct {
	ID         types.CharmModelID         `json:"id,omitempty"`
	CreatedAt  time.Time                  `json:"created_at,omitempty"`
	UpdatedAt  time.Time                  `json:"updated_at,omitempty"`
	DeletedAt  *time.Time                 `json:"deleted_at,omitempty"`
	CreatorID  types.CharmCreatorID       `json:"creator_id,omitempty"`
	Name       string                     `json:"name,omitempty"`
	Type       types.CharmModelType       `json:"type,omitempty"`
	Generation types.CharmModelGeneration `json:"generation,omitempty"`
}

// CharmCreator ...
type CharmCreator struct {
	ID        types.CharmCreatorID `json:"id,omitempty"`
	CreatedAt time.Time            `json:"created_at,omitempty"`
	UpdatedAt time.Time            `json:"updated_at,omitempty"`
	DeletedAt *time.Time           `json:"deleted_at,omitempty"`
	Name      string               `json:"name,omitempty"`
	Type      types.ArsenalType    `json:"type,omitempty"`
}
