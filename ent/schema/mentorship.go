package schema

import (
	"study-event-go/types"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Mentorship holds the schema definition for the Mentorship entity.
type Mentorship struct {
	ent.Schema
}

// Fields of the Mentorship.
func (Mentorship) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(types.MentorshipID(0)),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
		field.String("name").Unique().NotEmpty(),
	}
}

// Edges of the Mentorship.
func (Mentorship) Edges() []ent.Edge {
	return nil
}
