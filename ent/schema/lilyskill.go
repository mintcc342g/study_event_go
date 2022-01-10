package schema

import (
	"study-event-go/types"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// LilySkill holds the schema definition for the LilySkill entity.
type LilySkill struct {
	ent.Schema
}

// Fields of the LilySkill.
func (LilySkill) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("lily_id").GoType(types.LilyID(0)),
		field.Uint64("skill_id").GoType(types.SkillID(0)),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the LilySkill.
func (LilySkill) Edges() []ent.Edge {
	return nil
}
