package schema

import "entgo.io/ent"

// CharmModel holds the schema definition for the CharmModel entity.
type CharmModel struct {
	ent.Schema
}

// Fields of the CharmModel.
func (CharmModel) Fields() []ent.Field {
	return nil
}

// Edges of the CharmModel.
func (CharmModel) Edges() []ent.Edge {
	return nil
}
