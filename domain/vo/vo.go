package vo

import "study-event-go/types"

// Name ...
type Name struct {
	First  string
	Middle string
	Last   string
}

// FullName ...
func (n *Name) FullName() string {
	return n.First + " " + n.Middle + " " + n.Last
}

// Huge ...
type Huge struct {
	Class types.HugeClass
	Type  types.HugeType
}
