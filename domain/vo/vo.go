package vo

import "study_event_go/types"

// Name ...
type Name struct {
	First  string
	Middle string
	Last   string
}

// Huge ...
type Huge struct {
	Class types.HugeClass
	Type  types.HugeType
}
