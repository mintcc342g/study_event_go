package schema

import (
	"study-event-go/types"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Charm holds the schema definition for the Charm entity.
type Charm struct {
	ent.Schema
}

// Fields of the Charm.
func (Charm) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(types.CharmID(0)),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
		field.String("name").Unique().NotEmpty(),
		field.Uint64("model_id").GoType(types.CharmModelID(0)),
		field.Uint64("owner_id").GoType(types.LilyID(0)),
	}
}

// Edges of the Charm.
func (Charm) Edges() []ent.Edge {
	return nil
}
