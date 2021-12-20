// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"study_event_go/ent/garden"
	"study_event_go/ent/mentorshipsystem"
	"study_event_go/types"

	"entgo.io/ent/dialect/sql"
)

// Garden is the model entity for the Garden schema.
type Garden struct {
	config `json:"-"`
	// ID of the ent.
	ID types.GardenID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// MentorshipSystemID holds the value of the "mentorship_system_id" field.
	MentorshipSystemID *types.MentorshipSystemID `json:"mentorship_system_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GardenQuery when eager-loading is set.
	Edges GardenEdges `json:"edges"`
}

// GardenEdges holds the relations/edges for other nodes in the graph.
type GardenEdges struct {
	// MentorshipSystem holds the value of the mentorship_system edge.
	MentorshipSystem *MentorshipSystem `json:"mentorship_system,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MentorshipSystemOrErr returns the MentorshipSystem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GardenEdges) MentorshipSystemOrErr() (*MentorshipSystem, error) {
	if e.loadedTypes[0] {
		if e.MentorshipSystem == nil {
			// The edge mentorship_system was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: mentorshipsystem.Label}
		}
		return e.MentorshipSystem, nil
	}
	return nil, &NotLoadedError{edge: "mentorship_system"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Garden) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case garden.FieldID, garden.FieldMentorshipSystemID:
			values[i] = new(sql.NullInt64)
		case garden.FieldName, garden.FieldLocation:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Garden", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Garden fields.
func (ga *Garden) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case garden.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ga.ID = types.GardenID(value.Int64)
		case garden.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ga.Name = value.String
			}
		case garden.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				ga.Location = value.String
			}
		case garden.FieldMentorshipSystemID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mentorship_system_id", values[i])
			} else if value.Valid {
				ga.MentorshipSystemID = new(types.MentorshipSystemID)
				*ga.MentorshipSystemID = types.MentorshipSystemID(value.Int64)
			}
		}
	}
	return nil
}

// QueryMentorshipSystem queries the "mentorship_system" edge of the Garden entity.
func (ga *Garden) QueryMentorshipSystem() *MentorshipSystemQuery {
	return (&GardenClient{config: ga.config}).QueryMentorshipSystem(ga)
}

// Update returns a builder for updating this Garden.
// Note that you need to call Garden.Unwrap() before calling this method if this Garden
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *Garden) Update() *GardenUpdateOne {
	return (&GardenClient{config: ga.config}).UpdateOne(ga)
}

// Unwrap unwraps the Garden entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *Garden) Unwrap() *Garden {
	tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: Garden is not a transactional entity")
	}
	ga.config.driver = tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *Garden) String() string {
	var builder strings.Builder
	builder.WriteString("Garden(")
	builder.WriteString(fmt.Sprintf("id=%v", ga.ID))
	builder.WriteString(", name=")
	builder.WriteString(ga.Name)
	builder.WriteString(", location=")
	builder.WriteString(ga.Location)
	if v := ga.MentorshipSystemID; v != nil {
		builder.WriteString(", mentorship_system_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Gardens is a parsable slice of Garden.
type Gardens []*Garden

func (ga Gardens) config(cfg config) {
	for _i := range ga {
		ga[_i].config = cfg
	}
}
