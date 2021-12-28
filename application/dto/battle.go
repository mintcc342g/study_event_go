package dto

import "study-event-go/types"

// Alarm ...
type Alarm struct {
	GardenID     types.GardenID
	CaveLocation string
	Huges        []*Huge
	TotalCount   uint32
	AlertLevel   types.AlertLevel
}

// Huge ...
type Huge struct {
	Class types.HugeClass
	Type  types.HugeType
}
