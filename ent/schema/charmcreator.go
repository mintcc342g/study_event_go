package schema

import (
	"study-event-go/types"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CharmCreator holds the schema definition for the CharmCreator entity.
type CharmCreator struct {
	ent.Schema
}

// Fields of the CharmCreator.
func (CharmCreator) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(types.CharmCreatorID(0)),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
		field.String("name").Unique().NotEmpty(),
		field.Uint32("type").GoType(types.ArsenalType(0)),
	}
}

// Edges of the CharmCreator.
func (CharmCreator) Edges() []ent.Edge {
	return nil
}
