package schema

import (
	"study-event-go/types"
	"time"

	"entgo.io/ent"
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
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
		field.String("name"),
		field.String("location"),
		field.Uint64("mentorship_id").
			GoType(types.MentorshipID(0)),
	}
}

// Edges of the Garden.
func (Garden) Edges() []ent.Edge {
	return nil
}
