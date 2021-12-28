package schema

import (
	"study-event-go/types"
	"time"

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
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
		field.String("name").Unique().NotEmpty(),
	}
}

// Edges of the MentorshipSystem.
func (MentorshipSystem) Edges() []ent.Edge {
	return nil
}
