// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"study-event-go/ent/garden"
	"study-event-go/ent/predicate"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GardenUpdate is the builder for updating Garden entities.
type GardenUpdate struct {
	config
	hooks    []Hook
	mutation *GardenMutation
}

// Where appends a list predicates to the GardenUpdate builder.
func (gu *GardenUpdate) Where(ps ...predicate.Garden) *GardenUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetCreatedAt sets the "created_at" field.
func (gu *GardenUpdate) SetCreatedAt(t time.Time) *GardenUpdate {
	gu.mutation.SetCreatedAt(t)
	return gu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gu *GardenUpdate) SetNillableCreatedAt(t *time.Time) *GardenUpdate {
	if t != nil {
		gu.SetCreatedAt(*t)
	}
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GardenUpdate) SetUpdatedAt(t time.Time) *GardenUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetDeletedAt sets the "deleted_at" field.
func (gu *GardenUpdate) SetDeletedAt(t time.Time) *GardenUpdate {
	gu.mutation.SetDeletedAt(t)
	return gu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gu *GardenUpdate) SetNillableDeletedAt(t *time.Time) *GardenUpdate {
	if t != nil {
		gu.SetDeletedAt(*t)
	}
	return gu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gu *GardenUpdate) ClearDeletedAt() *GardenUpdate {
	gu.mutation.ClearDeletedAt()
	return gu
}

// SetName sets the "name" field.
func (gu *GardenUpdate) SetName(s string) *GardenUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetLocation sets the "location" field.
func (gu *GardenUpdate) SetLocation(s string) *GardenUpdate {
	gu.mutation.SetLocation(s)
	return gu
}

// SetMentorshipID sets the "mentorship_id" field.
func (gu *GardenUpdate) SetMentorshipID(ti types.MentorshipID) *GardenUpdate {
	gu.mutation.ResetMentorshipID()
	gu.mutation.SetMentorshipID(ti)
	return gu
}

// AddMentorshipID adds ti to the "mentorship_id" field.
func (gu *GardenUpdate) AddMentorshipID(ti types.MentorshipID) *GardenUpdate {
	gu.mutation.AddMentorshipID(ti)
	return gu
}

// Mutation returns the GardenMutation object of the builder.
func (gu *GardenUpdate) Mutation() *GardenMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GardenUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gu.defaults()
	if len(gu.hooks) == 0 {
		if err = gu.check(); err != nil {
			return 0, err
		}
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GardenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gu.check(); err != nil {
				return 0, err
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GardenUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GardenUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GardenUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GardenUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := garden.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GardenUpdate) check() error {
	if v, ok := gu.mutation.Name(); ok {
		if err := garden.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (gu *GardenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   garden.Table,
			Columns: garden.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: garden.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: garden.FieldCreatedAt,
		})
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: garden.FieldUpdatedAt,
		})
	}
	if value, ok := gu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: garden.FieldDeletedAt,
		})
	}
	if gu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: garden.FieldDeletedAt,
		})
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: garden.FieldName,
		})
	}
	if value, ok := gu.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: garden.FieldLocation,
		})
	}
	if value, ok := gu.mutation.MentorshipID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: garden.FieldMentorshipID,
		})
	}
	if value, ok := gu.mutation.AddedMentorshipID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: garden.FieldMentorshipID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{garden.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// GardenUpdateOne is the builder for updating a single Garden entity.
type GardenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GardenMutation
}

// SetCreatedAt sets the "created_at" field.
func (guo *GardenUpdateOne) SetCreatedAt(t time.Time) *GardenUpdateOne {
	guo.mutation.SetCreatedAt(t)
	return guo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guo *GardenUpdateOne) SetNillableCreatedAt(t *time.Time) *GardenUpdateOne {
	if t != nil {
		guo.SetCreatedAt(*t)
	}
	return guo
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GardenUpdateOne) SetUpdatedAt(t time.Time) *GardenUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetDeletedAt sets the "deleted_at" field.
func (guo *GardenUpdateOne) SetDeletedAt(t time.Time) *GardenUpdateOne {
	guo.mutation.SetDeletedAt(t)
	return guo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guo *GardenUpdateOne) SetNillableDeletedAt(t *time.Time) *GardenUpdateOne {
	if t != nil {
		guo.SetDeletedAt(*t)
	}
	return guo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (guo *GardenUpdateOne) ClearDeletedAt() *GardenUpdateOne {
	guo.mutation.ClearDeletedAt()
	return guo
}

// SetName sets the "name" field.
func (guo *GardenUpdateOne) SetName(s string) *GardenUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetLocation sets the "location" field.
func (guo *GardenUpdateOne) SetLocation(s string) *GardenUpdateOne {
	guo.mutation.SetLocation(s)
	return guo
}

// SetMentorshipID sets the "mentorship_id" field.
func (guo *GardenUpdateOne) SetMentorshipID(ti types.MentorshipID) *GardenUpdateOne {
	guo.mutation.ResetMentorshipID()
	guo.mutation.SetMentorshipID(ti)
	return guo
}

// AddMentorshipID adds ti to the "mentorship_id" field.
func (guo *GardenUpdateOne) AddMentorshipID(ti types.MentorshipID) *GardenUpdateOne {
	guo.mutation.AddMentorshipID(ti)
	return guo
}

// Mutation returns the GardenMutation object of the builder.
func (guo *GardenUpdateOne) Mutation() *GardenMutation {
	return guo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GardenUpdateOne) Select(field string, fields ...string) *GardenUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Garden entity.
func (guo *GardenUpdateOne) Save(ctx context.Context) (*Garden, error) {
	var (
		err  error
		node *Garden
	)
	guo.defaults()
	if len(guo.hooks) == 0 {
		if err = guo.check(); err != nil {
			return nil, err
		}
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GardenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guo.check(); err != nil {
				return nil, err
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GardenUpdateOne) SaveX(ctx context.Context) *Garden {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GardenUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GardenUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GardenUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := garden.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GardenUpdateOne) check() error {
	if v, ok := guo.mutation.Name(); ok {
		if err := garden.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (guo *GardenUpdateOne) sqlSave(ctx context.Context) (_node *Garden, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   garden.Table,
			Columns: garden.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: garden.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Garden.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, garden.FieldID)
		for _, f := range fields {
			if !garden.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != garden.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: garden.FieldCreatedAt,
		})
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: garden.FieldUpdatedAt,
		})
	}
	if value, ok := guo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: garden.FieldDeletedAt,
		})
	}
	if guo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: garden.FieldDeletedAt,
		})
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: garden.FieldName,
		})
	}
	if value, ok := guo.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: garden.FieldLocation,
		})
	}
	if value, ok := guo.mutation.MentorshipID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: garden.FieldMentorshipID,
		})
	}
	if value, ok := guo.mutation.AddedMentorshipID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: garden.FieldMentorshipID,
		})
	}
	_node = &Garden{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{garden.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
