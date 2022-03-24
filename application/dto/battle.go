package dto

import "study-event-go/types"

// Alarm ...
type Alarm struct {
	GardenID          types.GardenID
	CaveLocation      string
	Huges             []*Huge
	TotalHugeCount    uint32
	AlertLevel        types.AlertLevel
	LegionMemberCount uint32
}

// Huge ...
type Huge struct {
	Class types.HugeClass
	Type  types.HugeType
}
