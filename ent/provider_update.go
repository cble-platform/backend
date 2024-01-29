// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble-backend/ent/blueprint"
	"github.com/cble-platform/cble-backend/ent/predicate"
	"github.com/cble-platform/cble-backend/ent/provider"
	"github.com/google/uuid"
)

// ProviderUpdate is the builder for updating Provider entities.
type ProviderUpdate struct {
	config
	hooks    []Hook
	mutation *ProviderMutation
}

// Where appends a list predicates to the ProviderUpdate builder.
func (pu *ProviderUpdate) Where(ps ...predicate.Provider) *ProviderUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProviderUpdate) SetUpdatedAt(t time.Time) *ProviderUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetDisplayName sets the "display_name" field.
func (pu *ProviderUpdate) SetDisplayName(s string) *ProviderUpdate {
	pu.mutation.SetDisplayName(s)
	return pu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (pu *ProviderUpdate) SetNillableDisplayName(s *string) *ProviderUpdate {
	if s != nil {
		pu.SetDisplayName(*s)
	}
	return pu
}

// SetProviderGitURL sets the "provider_git_url" field.
func (pu *ProviderUpdate) SetProviderGitURL(s string) *ProviderUpdate {
	pu.mutation.SetProviderGitURL(s)
	return pu
}

// SetNillableProviderGitURL sets the "provider_git_url" field if the given value is not nil.
func (pu *ProviderUpdate) SetNillableProviderGitURL(s *string) *ProviderUpdate {
	if s != nil {
		pu.SetProviderGitURL(*s)
	}
	return pu
}

// SetProviderVersion sets the "provider_version" field.
func (pu *ProviderUpdate) SetProviderVersion(s string) *ProviderUpdate {
	pu.mutation.SetProviderVersion(s)
	return pu
}

// SetNillableProviderVersion sets the "provider_version" field if the given value is not nil.
func (pu *ProviderUpdate) SetNillableProviderVersion(s *string) *ProviderUpdate {
	if s != nil {
		pu.SetProviderVersion(*s)
	}
	return pu
}

// SetConfigBytes sets the "config_bytes" field.
func (pu *ProviderUpdate) SetConfigBytes(b []byte) *ProviderUpdate {
	pu.mutation.SetConfigBytes(b)
	return pu
}

// SetIsLoaded sets the "is_loaded" field.
func (pu *ProviderUpdate) SetIsLoaded(b bool) *ProviderUpdate {
	pu.mutation.SetIsLoaded(b)
	return pu
}

// SetNillableIsLoaded sets the "is_loaded" field if the given value is not nil.
func (pu *ProviderUpdate) SetNillableIsLoaded(b *bool) *ProviderUpdate {
	if b != nil {
		pu.SetIsLoaded(*b)
	}
	return pu
}

// AddBlueprintIDs adds the "blueprints" edge to the Blueprint entity by IDs.
func (pu *ProviderUpdate) AddBlueprintIDs(ids ...uuid.UUID) *ProviderUpdate {
	pu.mutation.AddBlueprintIDs(ids...)
	return pu
}

// AddBlueprints adds the "blueprints" edges to the Blueprint entity.
func (pu *ProviderUpdate) AddBlueprints(b ...*Blueprint) *ProviderUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.AddBlueprintIDs(ids...)
}

// Mutation returns the ProviderMutation object of the builder.
func (pu *ProviderUpdate) Mutation() *ProviderMutation {
	return pu.mutation
}

// ClearBlueprints clears all "blueprints" edges to the Blueprint entity.
func (pu *ProviderUpdate) ClearBlueprints() *ProviderUpdate {
	pu.mutation.ClearBlueprints()
	return pu
}

// RemoveBlueprintIDs removes the "blueprints" edge to Blueprint entities by IDs.
func (pu *ProviderUpdate) RemoveBlueprintIDs(ids ...uuid.UUID) *ProviderUpdate {
	pu.mutation.RemoveBlueprintIDs(ids...)
	return pu
}

// RemoveBlueprints removes "blueprints" edges to Blueprint entities.
func (pu *ProviderUpdate) RemoveBlueprints(b ...*Blueprint) *ProviderUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.RemoveBlueprintIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProviderUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProviderUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProviderUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProviderUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := provider.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *ProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(provider.Table, provider.Columns, sqlgraph.NewFieldSpec(provider.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(provider.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.DisplayName(); ok {
		_spec.SetField(provider.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := pu.mutation.ProviderGitURL(); ok {
		_spec.SetField(provider.FieldProviderGitURL, field.TypeString, value)
	}
	if value, ok := pu.mutation.ProviderVersion(); ok {
		_spec.SetField(provider.FieldProviderVersion, field.TypeString, value)
	}
	if value, ok := pu.mutation.ConfigBytes(); ok {
		_spec.SetField(provider.FieldConfigBytes, field.TypeBytes, value)
	}
	if value, ok := pu.mutation.IsLoaded(); ok {
		_spec.SetField(provider.FieldIsLoaded, field.TypeBool, value)
	}
	if pu.mutation.BlueprintsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provider.BlueprintsTable,
			Columns: []string{provider.BlueprintsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedBlueprintsIDs(); len(nodes) > 0 && !pu.mutation.BlueprintsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provider.BlueprintsTable,
			Columns: []string{provider.BlueprintsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.BlueprintsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provider.BlueprintsTable,
			Columns: []string{provider.BlueprintsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProviderUpdateOne is the builder for updating a single Provider entity.
type ProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProviderMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProviderUpdateOne) SetUpdatedAt(t time.Time) *ProviderUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetDisplayName sets the "display_name" field.
func (puo *ProviderUpdateOne) SetDisplayName(s string) *ProviderUpdateOne {
	puo.mutation.SetDisplayName(s)
	return puo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableDisplayName(s *string) *ProviderUpdateOne {
	if s != nil {
		puo.SetDisplayName(*s)
	}
	return puo
}

// SetProviderGitURL sets the "provider_git_url" field.
func (puo *ProviderUpdateOne) SetProviderGitURL(s string) *ProviderUpdateOne {
	puo.mutation.SetProviderGitURL(s)
	return puo
}

// SetNillableProviderGitURL sets the "provider_git_url" field if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableProviderGitURL(s *string) *ProviderUpdateOne {
	if s != nil {
		puo.SetProviderGitURL(*s)
	}
	return puo
}

// SetProviderVersion sets the "provider_version" field.
func (puo *ProviderUpdateOne) SetProviderVersion(s string) *ProviderUpdateOne {
	puo.mutation.SetProviderVersion(s)
	return puo
}

// SetNillableProviderVersion sets the "provider_version" field if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableProviderVersion(s *string) *ProviderUpdateOne {
	if s != nil {
		puo.SetProviderVersion(*s)
	}
	return puo
}

// SetConfigBytes sets the "config_bytes" field.
func (puo *ProviderUpdateOne) SetConfigBytes(b []byte) *ProviderUpdateOne {
	puo.mutation.SetConfigBytes(b)
	return puo
}

// SetIsLoaded sets the "is_loaded" field.
func (puo *ProviderUpdateOne) SetIsLoaded(b bool) *ProviderUpdateOne {
	puo.mutation.SetIsLoaded(b)
	return puo
}

// SetNillableIsLoaded sets the "is_loaded" field if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableIsLoaded(b *bool) *ProviderUpdateOne {
	if b != nil {
		puo.SetIsLoaded(*b)
	}
	return puo
}

// AddBlueprintIDs adds the "blueprints" edge to the Blueprint entity by IDs.
func (puo *ProviderUpdateOne) AddBlueprintIDs(ids ...uuid.UUID) *ProviderUpdateOne {
	puo.mutation.AddBlueprintIDs(ids...)
	return puo
}

// AddBlueprints adds the "blueprints" edges to the Blueprint entity.
func (puo *ProviderUpdateOne) AddBlueprints(b ...*Blueprint) *ProviderUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.AddBlueprintIDs(ids...)
}

// Mutation returns the ProviderMutation object of the builder.
func (puo *ProviderUpdateOne) Mutation() *ProviderMutation {
	return puo.mutation
}

// ClearBlueprints clears all "blueprints" edges to the Blueprint entity.
func (puo *ProviderUpdateOne) ClearBlueprints() *ProviderUpdateOne {
	puo.mutation.ClearBlueprints()
	return puo
}

// RemoveBlueprintIDs removes the "blueprints" edge to Blueprint entities by IDs.
func (puo *ProviderUpdateOne) RemoveBlueprintIDs(ids ...uuid.UUID) *ProviderUpdateOne {
	puo.mutation.RemoveBlueprintIDs(ids...)
	return puo
}

// RemoveBlueprints removes "blueprints" edges to Blueprint entities.
func (puo *ProviderUpdateOne) RemoveBlueprints(b ...*Blueprint) *ProviderUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.RemoveBlueprintIDs(ids...)
}

// Where appends a list predicates to the ProviderUpdate builder.
func (puo *ProviderUpdateOne) Where(ps ...predicate.Provider) *ProviderUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProviderUpdateOne) Select(field string, fields ...string) *ProviderUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Provider entity.
func (puo *ProviderUpdateOne) Save(ctx context.Context) (*Provider, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProviderUpdateOne) SaveX(ctx context.Context) *Provider {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProviderUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProviderUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := provider.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *ProviderUpdateOne) sqlSave(ctx context.Context) (_node *Provider, err error) {
	_spec := sqlgraph.NewUpdateSpec(provider.Table, provider.Columns, sqlgraph.NewFieldSpec(provider.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Provider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, provider.FieldID)
		for _, f := range fields {
			if !provider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != provider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(provider.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.DisplayName(); ok {
		_spec.SetField(provider.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := puo.mutation.ProviderGitURL(); ok {
		_spec.SetField(provider.FieldProviderGitURL, field.TypeString, value)
	}
	if value, ok := puo.mutation.ProviderVersion(); ok {
		_spec.SetField(provider.FieldProviderVersion, field.TypeString, value)
	}
	if value, ok := puo.mutation.ConfigBytes(); ok {
		_spec.SetField(provider.FieldConfigBytes, field.TypeBytes, value)
	}
	if value, ok := puo.mutation.IsLoaded(); ok {
		_spec.SetField(provider.FieldIsLoaded, field.TypeBool, value)
	}
	if puo.mutation.BlueprintsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provider.BlueprintsTable,
			Columns: []string{provider.BlueprintsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedBlueprintsIDs(); len(nodes) > 0 && !puo.mutation.BlueprintsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provider.BlueprintsTable,
			Columns: []string{provider.BlueprintsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.BlueprintsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provider.BlueprintsTable,
			Columns: []string{provider.BlueprintsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Provider{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
