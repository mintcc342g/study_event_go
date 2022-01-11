// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"study-event-go/ent/charmmodel"

	"entgo.io/ent/dialect/sql"
)

// CharmModel is the model entity for the CharmModel schema.
type CharmModel struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CharmModel) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case charmmodel.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CharmModel", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CharmModel fields.
func (cm *CharmModel) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case charmmodel.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cm.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this CharmModel.
// Note that you need to call CharmModel.Unwrap() before calling this method if this CharmModel
// was returned from a transaction, and the transaction was committed or rolled back.
func (cm *CharmModel) Update() *CharmModelUpdateOne {
	return (&CharmModelClient{config: cm.config}).UpdateOne(cm)
}

// Unwrap unwraps the CharmModel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cm *CharmModel) Unwrap() *CharmModel {
	tx, ok := cm.config.driver.(*txDriver)
	if !ok {
		panic("ent: CharmModel is not a transactional entity")
	}
	cm.config.driver = tx.drv
	return cm
}

// String implements the fmt.Stringer.
func (cm *CharmModel) String() string {
	var builder strings.Builder
	builder.WriteString("CharmModel(")
	builder.WriteString(fmt.Sprintf("id=%v", cm.ID))
	builder.WriteByte(')')
	return builder.String()
}

// CharmModels is a parsable slice of CharmModel.
type CharmModels []*CharmModel

func (cm CharmModels) config(cfg config) {
	for _i := range cm {
		cm[_i].config = cfg
	}
}
