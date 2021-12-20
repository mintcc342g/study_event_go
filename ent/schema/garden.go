package schema

import (
	"study_event_go/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Garden holds the schema definition for the Garden entity.
type Garden struct {
	ent.Schema
}

// Fields of the Garden.
func (Garden) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(types.GardenID(0)),
		field.String("name"),
		field.String("location"),
		field.Uint64("mentorship_system_id").
			GoType(types.MentorshipSystemID(0)).
			Optional().
			Nillable(),
	}
}

// Edges of the Garden.
func (Garden) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mentorship_system", MentorshipSystem.Type).
			StorageKey(edge.Symbol("garden_mentorship_system")).
			Field("mentorship_system_id").
			Unique(),
	}
}
