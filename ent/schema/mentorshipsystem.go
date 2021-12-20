package schema

import (
	"study_event_go/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MentorshipSystem holds the schema definition for the MentorshipSystem entity.
type MentorshipSystem struct {
	ent.Schema
}

// Fields of the MentorshipSystem.
func (MentorshipSystem) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(types.MentorshipSystemID(0)),
		field.String("name"),
	}
}

// Edges of the MentorshipSystem.
func (MentorshipSystem) Edges() []ent.Edge {
	return nil
}
