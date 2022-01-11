package schema

import (
	"study-event-go/types"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CharmModel holds the schema definition for the CharmModel entity.
type CharmModel struct {
	ent.Schema
}

// Fields of the CharmModel.
func (CharmModel) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(types.CharmModelID(0)),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
		field.String("name").Unique().NotEmpty(),
		field.Uint64("creator_id").GoType(types.CharmCreatorID(0)),
	}
}

// Edges of the CharmModel.
func (CharmModel) Edges() []ent.Edge {
	return nil
}
