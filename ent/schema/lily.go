package schema

import (
	"study-event-go/types"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Lily holds the schema definition for the Lily entity.
type Lily struct {
	ent.Schema
}

// Fields of the Lily.
func (Lily) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(types.LilyID(0)),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
		field.Uint32("cause_of_deletion").GoType(types.CauseOfDeletion(0)).Default(0),
		field.Time("birth").Nillable().Optional(),
		field.String("first_name").NotEmpty(),
		field.String("middle_name"),
		field.String("last_name").NotEmpty(),
		field.Uint32("rank").Default(0),
		field.Uint64("garden_id").GoType(types.GardenID(0)).Default(0),
		field.Uint64("legion_id").GoType(types.LegionID(0)).Default(0),
	}
}

// Edges of the Lily.
func (Lily) Edges() []ent.Edge {
	return nil
}
