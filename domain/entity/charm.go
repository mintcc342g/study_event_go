package entity

import (
	"study-event-go/types"
	"time"
)

// Charm ...
type Charm struct {
	ID        types.CharmID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	ModelID   types.CharmModelID
	Owner     types.LilyID
}

// CharmModel ...
type CharmModel struct {
	ID        types.CharmID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Creator   types.CharmCreatorID
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
