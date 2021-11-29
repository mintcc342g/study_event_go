package dto

import "study_event_go/types"

// Alarm ...
type Alarm struct {
	CaveLocation string           `json:"caveLocation"`
	Huges        []*Huge          `json:"huges"`
	TotalCount   uint32           `json:"totalCount"`
	AlertLevel   types.AlertLevel `json:"alertLevel"`
}

// Huge ...
type Huge struct {
	Class types.HugeClass `json:"class"`
	Type  types.HugeType  `json:"type"`
}
