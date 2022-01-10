// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"study-event-go/ent/lilyskill"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
)

// LilySkill is the model entity for the LilySkill schema.
type LilySkill struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// LilyID holds the value of the "lily_id" field.
	LilyID types.LilyID `json:"lily_id,omitempty"`
	// SkillID holds the value of the "skill_id" field.
	SkillID types.SkillID `json:"skill_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LilySkill) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case lilyskill.FieldID, lilyskill.FieldLilyID, lilyskill.FieldSkillID:
			values[i] = new(sql.NullInt64)
		case lilyskill.FieldCreatedAt, lilyskill.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type LilySkill", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LilySkill fields.
func (ls *LilySkill) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lilyskill.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ls.ID = int(value.Int64)
		case lilyskill.FieldLilyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lily_id", values[i])
			} else if value.Valid {
				ls.LilyID = types.LilyID(value.Int64)
			}
		case lilyskill.FieldSkillID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field skill_id", values[i])
			} else if value.Valid {
				ls.SkillID = types.SkillID(value.Int64)
			}
		case lilyskill.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ls.CreatedAt = value.Time
			}
		case lilyskill.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ls.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this LilySkill.
// Note that you need to call LilySkill.Unwrap() before calling this method if this LilySkill
// was returned from a transaction, and the transaction was committed or rolled back.
func (ls *LilySkill) Update() *LilySkillUpdateOne {
	return (&LilySkillClient{config: ls.config}).UpdateOne(ls)
}

// Unwrap unwraps the LilySkill entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ls *LilySkill) Unwrap() *LilySkill {
	tx, ok := ls.config.driver.(*txDriver)
	if !ok {
		panic("ent: LilySkill is not a transactional entity")
	}
	ls.config.driver = tx.drv
	return ls
}

// String implements the fmt.Stringer.
func (ls *LilySkill) String() string {
	var builder strings.Builder
	builder.WriteString("LilySkill(")
	builder.WriteString(fmt.Sprintf("id=%v", ls.ID))
	builder.WriteString(", lily_id=")
	builder.WriteString(fmt.Sprintf("%v", ls.LilyID))
	builder.WriteString(", skill_id=")
	builder.WriteString(fmt.Sprintf("%v", ls.SkillID))
	builder.WriteString(", created_at=")
	builder.WriteString(ls.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ls.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// LilySkills is a parsable slice of LilySkill.
type LilySkills []*LilySkill

func (ls LilySkills) config(cfg config) {
	for _i := range ls {
		ls[_i].config = cfg
	}
}
