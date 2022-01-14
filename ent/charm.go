// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"study-event-go/ent/charm"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Charm is the model entity for the Charm schema.
type Charm struct {
	config `json:"-"`
	// ID of the ent.
	ID types.CharmID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type types.CharmType `json:"type,omitempty"`
	// ModelID holds the value of the "model_id" field.
	ModelID types.CharmModelID `json:"model_id,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID types.LilyID `json:"owner_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Charm) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case charm.FieldID, charm.FieldType, charm.FieldModelID, charm.FieldOwnerID:
			values[i] = new(sql.NullInt64)
		case charm.FieldName:
			values[i] = new(sql.NullString)
		case charm.FieldCreatedAt, charm.FieldUpdatedAt, charm.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Charm", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Charm fields.
func (c *Charm) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case charm.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = types.CharmID(value.Int64)
		case charm.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case charm.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case charm.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = new(time.Time)
				*c.DeletedAt = value.Time
			}
		case charm.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case charm.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = types.CharmType(value.Int64)
			}
		case charm.FieldModelID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field model_id", values[i])
			} else if value.Valid {
				c.ModelID = types.CharmModelID(value.Int64)
			}
		case charm.FieldOwnerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				c.OwnerID = types.LilyID(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Charm.
// Note that you need to call Charm.Unwrap() before calling this method if this Charm
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Charm) Update() *CharmUpdateOne {
	return (&CharmClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Charm entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Charm) Unwrap() *Charm {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Charm is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Charm) String() string {
	var builder strings.Builder
	builder.WriteString("Charm(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	if v := c.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", c.Type))
	builder.WriteString(", model_id=")
	builder.WriteString(fmt.Sprintf("%v", c.ModelID))
	builder.WriteString(", owner_id=")
	builder.WriteString(fmt.Sprintf("%v", c.OwnerID))
	builder.WriteByte(')')
	return builder.String()
}

// Charms is a parsable slice of Charm.
type Charms []*Charm

func (c Charms) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
