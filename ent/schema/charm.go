package schema

import "entgo.io/ent"

// Charm holds the schema definition for the Charm entity.
type Charm struct {
	ent.Schema
}

// Fields of the Charm.
func (Charm) Fields() []ent.Field {
	return nil
}

// Edges of the Charm.
func (Charm) Edges() []ent.Edge {
	return nil
}
