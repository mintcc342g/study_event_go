package dto

import "study_event_go/types"

// Warning ...
type Warning struct {
	CaveLocation string           `json:"caveLocation"`
	Huges        []*Huge          `json:"huges"`
	TotalCount   uint32           `json:"totalCount"`
	AlertLevel   types.AlertLevel `json:"alertLevel"`
}

// Huge ...
type Huge struct {
	HugeClass types.HugeClass `json:"hugeClass"`
	HugeType  types.HugeType  `json:"hugeType"`
}
