package vo

import "study-event-go/types"

// Name ...
type Name struct {
	First  string
	Middle string
	Last   string
}

// NewName ...
func NewName(first, middle, last string) *Name {
	return &Name{
		First:  first,
		Middle: middle,
		Last:   last,
	}
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
