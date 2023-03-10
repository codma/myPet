// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/predicate"
	"main/ent/store"
	"main/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StoreUpdate is the builder for updating Store entities.
type StoreUpdate struct {
	config
	hooks    []Hook
	mutation *StoreMutation
}

// Where appends a list predicates to the StoreUpdate builder.
func (su *StoreUpdate) Where(ps ...predicate.Store) *StoreUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetIsUsed sets the "is_used" field.
func (su *StoreUpdate) SetIsUsed(b bool) *StoreUpdate {
	su.mutation.SetIsUsed(b)
	return su
}

// SetNillableIsUsed sets the "is_used" field if the given value is not nil.
func (su *StoreUpdate) SetNillableIsUsed(b *bool) *StoreUpdate {
	if b != nil {
		su.SetIsUsed(*b)
	}
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *StoreUpdate) SetCreatedAt(t time.Time) *StoreUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *StoreUpdate) SetNillableCreatedAt(t *time.Time) *StoreUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *StoreUpdate) SetUpdatedAt(t time.Time) *StoreUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetStoreID sets the "store_id" field.
func (su *StoreUpdate) SetStoreID(s string) *StoreUpdate {
	su.mutation.SetStoreID(s)
	return su
}

// SetAppAccessKey sets the "app_access_key" field.
func (su *StoreUpdate) SetAppAccessKey(s string) *StoreUpdate {
	su.mutation.SetAppAccessKey(s)
	return su
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (su *StoreUpdate) AddUserIDs(ids ...uint64) *StoreUpdate {
	su.mutation.AddUserIDs(ids...)
	return su
}

// AddUsers adds the "users" edges to the User entity.
func (su *StoreUpdate) AddUsers(u ...*User) *StoreUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.AddUserIDs(ids...)
}

// Mutation returns the StoreMutation object of the builder.
func (su *StoreUpdate) Mutation() *StoreMutation {
	return su.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (su *StoreUpdate) ClearUsers() *StoreUpdate {
	su.mutation.ClearUsers()
	return su
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (su *StoreUpdate) RemoveUserIDs(ids ...uint64) *StoreUpdate {
	su.mutation.RemoveUserIDs(ids...)
	return su
}

// RemoveUsers removes "users" edges to User entities.
func (su *StoreUpdate) RemoveUsers(u ...*User) *StoreUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StoreUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks[int, StoreMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StoreUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StoreUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StoreUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *StoreUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := store.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *StoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   store.Table,
			Columns: store.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: store.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.IsUsed(); ok {
		_spec.SetField(store.FieldIsUsed, field.TypeBool, value)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(store.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(store.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.StoreID(); ok {
		_spec.SetField(store.FieldStoreID, field.TypeString, value)
	}
	if value, ok := su.mutation.AppAccessKey(); ok {
		_spec.SetField(store.FieldAppAccessKey, field.TypeString, value)
	}
	if su.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   store.UsersTable,
			Columns: store.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedUsersIDs(); len(nodes) > 0 && !su.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   store.UsersTable,
			Columns: store.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   store.UsersTable,
			Columns: store.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{store.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StoreUpdateOne is the builder for updating a single Store entity.
type StoreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StoreMutation
}

// SetIsUsed sets the "is_used" field.
func (suo *StoreUpdateOne) SetIsUsed(b bool) *StoreUpdateOne {
	suo.mutation.SetIsUsed(b)
	return suo
}

// SetNillableIsUsed sets the "is_used" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableIsUsed(b *bool) *StoreUpdateOne {
	if b != nil {
		suo.SetIsUsed(*b)
	}
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *StoreUpdateOne) SetCreatedAt(t time.Time) *StoreUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableCreatedAt(t *time.Time) *StoreUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *StoreUpdateOne) SetUpdatedAt(t time.Time) *StoreUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetStoreID sets the "store_id" field.
func (suo *StoreUpdateOne) SetStoreID(s string) *StoreUpdateOne {
	suo.mutation.SetStoreID(s)
	return suo
}

// SetAppAccessKey sets the "app_access_key" field.
func (suo *StoreUpdateOne) SetAppAccessKey(s string) *StoreUpdateOne {
	suo.mutation.SetAppAccessKey(s)
	return suo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (suo *StoreUpdateOne) AddUserIDs(ids ...uint64) *StoreUpdateOne {
	suo.mutation.AddUserIDs(ids...)
	return suo
}

// AddUsers adds the "users" edges to the User entity.
func (suo *StoreUpdateOne) AddUsers(u ...*User) *StoreUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.AddUserIDs(ids...)
}

// Mutation returns the StoreMutation object of the builder.
func (suo *StoreUpdateOne) Mutation() *StoreMutation {
	return suo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (suo *StoreUpdateOne) ClearUsers() *StoreUpdateOne {
	suo.mutation.ClearUsers()
	return suo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (suo *StoreUpdateOne) RemoveUserIDs(ids ...uint64) *StoreUpdateOne {
	suo.mutation.RemoveUserIDs(ids...)
	return suo
}

// RemoveUsers removes "users" edges to User entities.
func (suo *StoreUpdateOne) RemoveUsers(u ...*User) *StoreUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StoreUpdateOne) Select(field string, fields ...string) *StoreUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Store entity.
func (suo *StoreUpdateOne) Save(ctx context.Context) (*Store, error) {
	suo.defaults()
	return withHooks[*Store, StoreMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StoreUpdateOne) SaveX(ctx context.Context) *Store {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StoreUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StoreUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *StoreUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := store.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *StoreUpdateOne) sqlSave(ctx context.Context) (_node *Store, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   store.Table,
			Columns: store.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: store.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Store.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, store.FieldID)
		for _, f := range fields {
			if !store.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != store.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.IsUsed(); ok {
		_spec.SetField(store.FieldIsUsed, field.TypeBool, value)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(store.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(store.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.StoreID(); ok {
		_spec.SetField(store.FieldStoreID, field.TypeString, value)
	}
	if value, ok := suo.mutation.AppAccessKey(); ok {
		_spec.SetField(store.FieldAppAccessKey, field.TypeString, value)
	}
	if suo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   store.UsersTable,
			Columns: store.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !suo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   store.UsersTable,
			Columns: store.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   store.UsersTable,
			Columns: store.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Store{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{store.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
