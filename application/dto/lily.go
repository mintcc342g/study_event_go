package dto

import (
	"study-event-go/types"
	"time"
)

// Lily ...
type Lily struct {
	ID              types.LilyID          `json:"id,omitempty"`
	CreatedAt       time.Time             `json:"created_at,omitempty"`
	UpdatedAt       time.Time             `json:"updated_at,omitempty"`
	DeletedAt       *time.Time            `json:"deleted_at,omitempty"`
	CauseOfDeletion types.CauseOfDeletion `json:"cause_of_deletion,omitempty"`
	FirstName       string                `json:"first_name,omitempty"`
	MiddleName      string                `json:"middle_name,omitempty"`
	LastName        string                `json:"last_name,omitempty"`
	Birth           time.Time             `json:"birth,omitempty"`
	Rank            uint32                `json:"rank,omitempty"`
	GardenID        types.GardenID        `json:"garden_id,omitempty"`
	LegionID        types.LegionID        `json:"legion_id,omitempty"`
	// Charms          []*Charms             `json:"charms,omitempty"`
	// Skills          []*Skill              `json:"skills,omitempty"`
	// Relations       []*LilysRelation      `json:"relations,omitempty"`
}
