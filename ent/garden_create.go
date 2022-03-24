// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"study-event-go/ent/garden"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GardenCreate is the builder for creating a Garden entity.
type GardenCreate struct {
	config
	mutation *GardenMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (gc *GardenCreate) SetCreatedAt(t time.Time) *GardenCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GardenCreate) SetNillableCreatedAt(t *time.Time) *GardenCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GardenCreate) SetUpdatedAt(t time.Time) *GardenCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GardenCreate) SetNillableUpdatedAt(t *time.Time) *GardenCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetDeletedAt sets the "deleted_at" field.
func (gc *GardenCreate) SetDeletedAt(t time.Time) *GardenCreate {
	gc.mutation.SetDeletedAt(t)
	return gc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gc *GardenCreate) SetNillableDeletedAt(t *time.Time) *GardenCreate {
	if t != nil {
		gc.SetDeletedAt(*t)
	}
	return gc
}

// SetName sets the "name" field.
func (gc *GardenCreate) SetName(s string) *GardenCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetLocation sets the "location" field.
func (gc *GardenCreate) SetLocation(s string) *GardenCreate {
	gc.mutation.SetLocation(s)
	return gc
}

// SetMentorshipID sets the "mentorship_id" field.
func (gc *GardenCreate) SetMentorshipID(ti types.MentorshipID) *GardenCreate {
	gc.mutation.SetMentorshipID(ti)
	return gc
}

// SetLegionSystem sets the "legion_system" field.
func (gc *GardenCreate) SetLegionSystem(ts types.LegionSystem) *GardenCreate {
	gc.mutation.SetLegionSystem(ts)
	return gc
}

// SetID sets the "id" field.
func (gc *GardenCreate) SetID(ti types.GardenID) *GardenCreate {
	gc.mutation.SetID(ti)
	return gc
}

// Mutation returns the GardenMutation object of the builder.
func (gc *GardenCreate) Mutation() *GardenMutation {
	return gc.mutation
}

// Save creates the Garden in the database.
func (gc *GardenCreate) Save(ctx context.Context) (*Garden, error) {
	var (
		err  error
		node *Garden
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GardenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			if node, err = gc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			if gc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GardenCreate) SaveX(ctx context.Context) *Garden {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GardenCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GardenCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GardenCreate) defaults() {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := garden.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		v := garden.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GardenCreate) check() error {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := gc.mutation.Name(); ok {
		if err := garden.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := gc.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "location"`)}
	}
	if _, ok := gc.mutation.MentorshipID(); !ok {
		return &ValidationError{Name: "mentorship_id", err: errors.New(`ent: missing required field "mentorship_id"`)}
	}
	if _, ok := gc.mutation.LegionSystem(); !ok {
		return &ValidationError{Name: "legion_system", err: errors.New(`ent: missing required field "legion_system"`)}
	}
	return nil
}

func (gc *GardenCreate) sqlSave(ctx context.Context) (*Garden, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = types.GardenID(id)
	}
	return _node, nil
}

func (gc *GardenCreate) createSpec() (*Garden, *sqlgraph.CreateSpec) {
	var (
		_node = &Garden{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: garden.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: garden.FieldID,
			},
		}
	)
	_spec.OnConflict = gc.conflict
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: garden.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: garden.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := gc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: garden.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: garden.FieldName,
		})
		_node.Name = value
	}
	if value, ok := gc.mutation.Location(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: garden.FieldLocation,
		})
		_node.Location = value
	}
	if value, ok := gc.mutation.MentorshipID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: garden.FieldMentorshipID,
		})
		_node.MentorshipID = value
	}
	if value, ok := gc.mutation.LegionSystem(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: garden.FieldLegionSystem,
		})
		_node.LegionSystem = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Garden.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GardenUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (gc *GardenCreate) OnConflict(opts ...sql.ConflictOption) *GardenUpsertOne {
	gc.conflict = opts
	return &GardenUpsertOne{
		create: gc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Garden.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (gc *GardenCreate) OnConflictColumns(columns ...string) *GardenUpsertOne {
	gc.conflict = append(gc.conflict, sql.ConflictColumns(columns...))
	return &GardenUpsertOne{
		create: gc,
	}
}

type (
	// GardenUpsertOne is the builder for "upsert"-ing
	//  one Garden node.
	GardenUpsertOne struct {
		create *GardenCreate
	}

	// GardenUpsert is the "OnConflict" setter.
	GardenUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *GardenUpsert) SetCreatedAt(v time.Time) *GardenUpsert {
	u.Set(garden.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GardenUpsert) UpdateCreatedAt() *GardenUpsert {
	u.SetExcluded(garden.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GardenUpsert) SetUpdatedAt(v time.Time) *GardenUpsert {
	u.Set(garden.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GardenUpsert) UpdateUpdatedAt() *GardenUpsert {
	u.SetExcluded(garden.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GardenUpsert) SetDeletedAt(v time.Time) *GardenUpsert {
	u.Set(garden.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GardenUpsert) UpdateDeletedAt() *GardenUpsert {
	u.SetExcluded(garden.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *GardenUpsert) ClearDeletedAt() *GardenUpsert {
	u.SetNull(garden.FieldDeletedAt)
	return u
}

// SetName sets the "name" field.
func (u *GardenUpsert) SetName(v string) *GardenUpsert {
	u.Set(garden.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GardenUpsert) UpdateName() *GardenUpsert {
	u.SetExcluded(garden.FieldName)
	return u
}

// SetLocation sets the "location" field.
func (u *GardenUpsert) SetLocation(v string) *GardenUpsert {
	u.Set(garden.FieldLocation, v)
	return u
}

// UpdateLocation sets the "location" field to the value that was provided on create.
func (u *GardenUpsert) UpdateLocation() *GardenUpsert {
	u.SetExcluded(garden.FieldLocation)
	return u
}

// SetMentorshipID sets the "mentorship_id" field.
func (u *GardenUpsert) SetMentorshipID(v types.MentorshipID) *GardenUpsert {
	u.Set(garden.FieldMentorshipID, v)
	return u
}

// UpdateMentorshipID sets the "mentorship_id" field to the value that was provided on create.
func (u *GardenUpsert) UpdateMentorshipID() *GardenUpsert {
	u.SetExcluded(garden.FieldMentorshipID)
	return u
}

// SetLegionSystem sets the "legion_system" field.
func (u *GardenUpsert) SetLegionSystem(v types.LegionSystem) *GardenUpsert {
	u.Set(garden.FieldLegionSystem, v)
	return u
}

// UpdateLegionSystem sets the "legion_system" field to the value that was provided on create.
func (u *GardenUpsert) UpdateLegionSystem() *GardenUpsert {
	u.SetExcluded(garden.FieldLegionSystem)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Garden.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(garden.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *GardenUpsertOne) UpdateNewValues() *GardenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(garden.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Garden.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *GardenUpsertOne) Ignore() *GardenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GardenUpsertOne) DoNothing() *GardenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GardenCreate.OnConflict
// documentation for more info.
func (u *GardenUpsertOne) Update(set func(*GardenUpsert)) *GardenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GardenUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GardenUpsertOne) SetCreatedAt(v time.Time) *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GardenUpsertOne) UpdateCreatedAt() *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GardenUpsertOne) SetUpdatedAt(v time.Time) *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GardenUpsertOne) UpdateUpdatedAt() *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GardenUpsertOne) SetDeletedAt(v time.Time) *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GardenUpsertOne) UpdateDeletedAt() *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *GardenUpsertOne) ClearDeletedAt() *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *GardenUpsertOne) SetName(v string) *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GardenUpsertOne) UpdateName() *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateName()
	})
}

// SetLocation sets the "location" field.
func (u *GardenUpsertOne) SetLocation(v string) *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.SetLocation(v)
	})
}

// UpdateLocation sets the "location" field to the value that was provided on create.
func (u *GardenUpsertOne) UpdateLocation() *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateLocation()
	})
}

// SetMentorshipID sets the "mentorship_id" field.
func (u *GardenUpsertOne) SetMentorshipID(v types.MentorshipID) *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.SetMentorshipID(v)
	})
}

// UpdateMentorshipID sets the "mentorship_id" field to the value that was provided on create.
func (u *GardenUpsertOne) UpdateMentorshipID() *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateMentorshipID()
	})
}

// SetLegionSystem sets the "legion_system" field.
func (u *GardenUpsertOne) SetLegionSystem(v types.LegionSystem) *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.SetLegionSystem(v)
	})
}

// UpdateLegionSystem sets the "legion_system" field to the value that was provided on create.
func (u *GardenUpsertOne) UpdateLegionSystem() *GardenUpsertOne {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateLegionSystem()
	})
}

// Exec executes the query.
func (u *GardenUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GardenCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GardenUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GardenUpsertOne) ID(ctx context.Context) (id types.GardenID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GardenUpsertOne) IDX(ctx context.Context) types.GardenID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GardenCreateBulk is the builder for creating many Garden entities in bulk.
type GardenCreateBulk struct {
	config
	builders []*GardenCreate
	conflict []sql.ConflictOption
}

// Save creates the Garden entities in the database.
func (gcb *GardenCreateBulk) Save(ctx context.Context) ([]*Garden, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Garden, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GardenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = types.GardenID(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GardenCreateBulk) SaveX(ctx context.Context) []*Garden {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GardenCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GardenCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Garden.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GardenUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (gcb *GardenCreateBulk) OnConflict(opts ...sql.ConflictOption) *GardenUpsertBulk {
	gcb.conflict = opts
	return &GardenUpsertBulk{
		create: gcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Garden.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (gcb *GardenCreateBulk) OnConflictColumns(columns ...string) *GardenUpsertBulk {
	gcb.conflict = append(gcb.conflict, sql.ConflictColumns(columns...))
	return &GardenUpsertBulk{
		create: gcb,
	}
}

// GardenUpsertBulk is the builder for "upsert"-ing
// a bulk of Garden nodes.
type GardenUpsertBulk struct {
	create *GardenCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Garden.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(garden.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *GardenUpsertBulk) UpdateNewValues() *GardenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(garden.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Garden.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *GardenUpsertBulk) Ignore() *GardenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GardenUpsertBulk) DoNothing() *GardenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GardenCreateBulk.OnConflict
// documentation for more info.
func (u *GardenUpsertBulk) Update(set func(*GardenUpsert)) *GardenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GardenUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GardenUpsertBulk) SetCreatedAt(v time.Time) *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GardenUpsertBulk) UpdateCreatedAt() *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GardenUpsertBulk) SetUpdatedAt(v time.Time) *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GardenUpsertBulk) UpdateUpdatedAt() *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GardenUpsertBulk) SetDeletedAt(v time.Time) *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GardenUpsertBulk) UpdateDeletedAt() *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *GardenUpsertBulk) ClearDeletedAt() *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *GardenUpsertBulk) SetName(v string) *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GardenUpsertBulk) UpdateName() *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateName()
	})
}

// SetLocation sets the "location" field.
func (u *GardenUpsertBulk) SetLocation(v string) *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.SetLocation(v)
	})
}

// UpdateLocation sets the "location" field to the value that was provided on create.
func (u *GardenUpsertBulk) UpdateLocation() *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateLocation()
	})
}

// SetMentorshipID sets the "mentorship_id" field.
func (u *GardenUpsertBulk) SetMentorshipID(v types.MentorshipID) *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.SetMentorshipID(v)
	})
}

// UpdateMentorshipID sets the "mentorship_id" field to the value that was provided on create.
func (u *GardenUpsertBulk) UpdateMentorshipID() *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateMentorshipID()
	})
}

// SetLegionSystem sets the "legion_system" field.
func (u *GardenUpsertBulk) SetLegionSystem(v types.LegionSystem) *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.SetLegionSystem(v)
	})
}

// UpdateLegionSystem sets the "legion_system" field to the value that was provided on create.
func (u *GardenUpsertBulk) UpdateLegionSystem() *GardenUpsertBulk {
	return u.Update(func(s *GardenUpsert) {
		s.UpdateLegionSystem()
	})
}

// Exec executes the query.
func (u *GardenUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GardenCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GardenCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GardenUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
