package vo

import "study_event_go/types"

// Name ...
type Name struct {
	first  string
	middle string
	last   string
}

// Huge ...
type Huge struct {
	class *types.HugeClass
}
