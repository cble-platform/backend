// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/predicate"
	"github.com/cble-platform/cble-backend/ent/provider"
	"github.com/cble-platform/cble-backend/ent/providercommand"
	"github.com/google/uuid"
)

// ProviderCommandUpdate is the builder for updating ProviderCommand entities.
type ProviderCommandUpdate struct {
	config
	hooks    []Hook
	mutation *ProviderCommandMutation
}

// Where appends a list predicates to the ProviderCommandUpdate builder.
func (pcu *ProviderCommandUpdate) Where(ps ...predicate.ProviderCommand) *ProviderCommandUpdate {
	pcu.mutation.Where(ps...)
	return pcu
}

// SetUpdatedAt sets the "updated_at" field.
func (pcu *ProviderCommandUpdate) SetUpdatedAt(t time.Time) *ProviderCommandUpdate {
	pcu.mutation.SetUpdatedAt(t)
	return pcu
}

// SetCommandType sets the "command_type" field.
func (pcu *ProviderCommandUpdate) SetCommandType(pt providercommand.CommandType) *ProviderCommandUpdate {
	pcu.mutation.SetCommandType(pt)
	return pcu
}

// SetStatus sets the "status" field.
func (pcu *ProviderCommandUpdate) SetStatus(pr providercommand.Status) *ProviderCommandUpdate {
	pcu.mutation.SetStatus(pr)
	return pcu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pcu *ProviderCommandUpdate) SetNillableStatus(pr *providercommand.Status) *ProviderCommandUpdate {
	if pr != nil {
		pcu.SetStatus(*pr)
	}
	return pcu
}

// SetArguments sets the "arguments" field.
func (pcu *ProviderCommandUpdate) SetArguments(s []string) *ProviderCommandUpdate {
	pcu.mutation.SetArguments(s)
	return pcu
}

// AppendArguments appends s to the "arguments" field.
func (pcu *ProviderCommandUpdate) AppendArguments(s []string) *ProviderCommandUpdate {
	pcu.mutation.AppendArguments(s)
	return pcu
}

// ClearArguments clears the value of the "arguments" field.
func (pcu *ProviderCommandUpdate) ClearArguments() *ProviderCommandUpdate {
	pcu.mutation.ClearArguments()
	return pcu
}

// SetStartTime sets the "start_time" field.
func (pcu *ProviderCommandUpdate) SetStartTime(t time.Time) *ProviderCommandUpdate {
	pcu.mutation.SetStartTime(t)
	return pcu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (pcu *ProviderCommandUpdate) SetNillableStartTime(t *time.Time) *ProviderCommandUpdate {
	if t != nil {
		pcu.SetStartTime(*t)
	}
	return pcu
}

// ClearStartTime clears the value of the "start_time" field.
func (pcu *ProviderCommandUpdate) ClearStartTime() *ProviderCommandUpdate {
	pcu.mutation.ClearStartTime()
	return pcu
}

// SetEndTime sets the "end_time" field.
func (pcu *ProviderCommandUpdate) SetEndTime(t time.Time) *ProviderCommandUpdate {
	pcu.mutation.SetEndTime(t)
	return pcu
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (pcu *ProviderCommandUpdate) SetNillableEndTime(t *time.Time) *ProviderCommandUpdate {
	if t != nil {
		pcu.SetEndTime(*t)
	}
	return pcu
}

// ClearEndTime clears the value of the "end_time" field.
func (pcu *ProviderCommandUpdate) ClearEndTime() *ProviderCommandUpdate {
	pcu.mutation.ClearEndTime()
	return pcu
}

// SetOutput sets the "output" field.
func (pcu *ProviderCommandUpdate) SetOutput(s string) *ProviderCommandUpdate {
	pcu.mutation.SetOutput(s)
	return pcu
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (pcu *ProviderCommandUpdate) SetNillableOutput(s *string) *ProviderCommandUpdate {
	if s != nil {
		pcu.SetOutput(*s)
	}
	return pcu
}

// SetError sets the "error" field.
func (pcu *ProviderCommandUpdate) SetError(s string) *ProviderCommandUpdate {
	pcu.mutation.SetError(s)
	return pcu
}

// SetNillableError sets the "error" field if the given value is not nil.
func (pcu *ProviderCommandUpdate) SetNillableError(s *string) *ProviderCommandUpdate {
	if s != nil {
		pcu.SetError(*s)
	}
	return pcu
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (pcu *ProviderCommandUpdate) SetProviderID(id uuid.UUID) *ProviderCommandUpdate {
	pcu.mutation.SetProviderID(id)
	return pcu
}

// SetProvider sets the "provider" edge to the Provider entity.
func (pcu *ProviderCommandUpdate) SetProvider(p *Provider) *ProviderCommandUpdate {
	return pcu.SetProviderID(p.ID)
}

// SetDeploymentID sets the "deployment" edge to the Deployment entity by ID.
func (pcu *ProviderCommandUpdate) SetDeploymentID(id uuid.UUID) *ProviderCommandUpdate {
	pcu.mutation.SetDeploymentID(id)
	return pcu
}

// SetNillableDeploymentID sets the "deployment" edge to the Deployment entity by ID if the given value is not nil.
func (pcu *ProviderCommandUpdate) SetNillableDeploymentID(id *uuid.UUID) *ProviderCommandUpdate {
	if id != nil {
		pcu = pcu.SetDeploymentID(*id)
	}
	return pcu
}

// SetDeployment sets the "deployment" edge to the Deployment entity.
func (pcu *ProviderCommandUpdate) SetDeployment(d *Deployment) *ProviderCommandUpdate {
	return pcu.SetDeploymentID(d.ID)
}

// Mutation returns the ProviderCommandMutation object of the builder.
func (pcu *ProviderCommandUpdate) Mutation() *ProviderCommandMutation {
	return pcu.mutation
}

// ClearProvider clears the "provider" edge to the Provider entity.
func (pcu *ProviderCommandUpdate) ClearProvider() *ProviderCommandUpdate {
	pcu.mutation.ClearProvider()
	return pcu
}

// ClearDeployment clears the "deployment" edge to the Deployment entity.
func (pcu *ProviderCommandUpdate) ClearDeployment() *ProviderCommandUpdate {
	pcu.mutation.ClearDeployment()
	return pcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *ProviderCommandUpdate) Save(ctx context.Context) (int, error) {
	pcu.defaults()
	return withHooks(ctx, pcu.sqlSave, pcu.mutation, pcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcu *ProviderCommandUpdate) SaveX(ctx context.Context) int {
	affected, err := pcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcu *ProviderCommandUpdate) Exec(ctx context.Context) error {
	_, err := pcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcu *ProviderCommandUpdate) ExecX(ctx context.Context) {
	if err := pcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcu *ProviderCommandUpdate) defaults() {
	if _, ok := pcu.mutation.UpdatedAt(); !ok {
		v := providercommand.UpdateDefaultUpdatedAt()
		pcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcu *ProviderCommandUpdate) check() error {
	if v, ok := pcu.mutation.CommandType(); ok {
		if err := providercommand.CommandTypeValidator(v); err != nil {
			return &ValidationError{Name: "command_type", err: fmt.Errorf(`ent: validator failed for field "ProviderCommand.command_type": %w`, err)}
		}
	}
	if v, ok := pcu.mutation.Status(); ok {
		if err := providercommand.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ProviderCommand.status": %w`, err)}
		}
	}
	if _, ok := pcu.mutation.ProviderID(); pcu.mutation.ProviderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProviderCommand.provider"`)
	}
	return nil
}

func (pcu *ProviderCommandUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(providercommand.Table, providercommand.Columns, sqlgraph.NewFieldSpec(providercommand.FieldID, field.TypeUUID))
	if ps := pcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcu.mutation.UpdatedAt(); ok {
		_spec.SetField(providercommand.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pcu.mutation.CommandType(); ok {
		_spec.SetField(providercommand.FieldCommandType, field.TypeEnum, value)
	}
	if value, ok := pcu.mutation.Status(); ok {
		_spec.SetField(providercommand.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := pcu.mutation.Arguments(); ok {
		_spec.SetField(providercommand.FieldArguments, field.TypeJSON, value)
	}
	if value, ok := pcu.mutation.AppendedArguments(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, providercommand.FieldArguments, value)
		})
	}
	if pcu.mutation.ArgumentsCleared() {
		_spec.ClearField(providercommand.FieldArguments, field.TypeJSON)
	}
	if value, ok := pcu.mutation.StartTime(); ok {
		_spec.SetField(providercommand.FieldStartTime, field.TypeTime, value)
	}
	if pcu.mutation.StartTimeCleared() {
		_spec.ClearField(providercommand.FieldStartTime, field.TypeTime)
	}
	if value, ok := pcu.mutation.EndTime(); ok {
		_spec.SetField(providercommand.FieldEndTime, field.TypeTime, value)
	}
	if pcu.mutation.EndTimeCleared() {
		_spec.ClearField(providercommand.FieldEndTime, field.TypeTime)
	}
	if value, ok := pcu.mutation.Output(); ok {
		_spec.SetField(providercommand.FieldOutput, field.TypeString, value)
	}
	if value, ok := pcu.mutation.Error(); ok {
		_spec.SetField(providercommand.FieldError, field.TypeString, value)
	}
	if pcu.mutation.ProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   providercommand.ProviderTable,
			Columns: []string{providercommand.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provider.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   providercommand.ProviderTable,
			Columns: []string{providercommand.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provider.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcu.mutation.DeploymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   providercommand.DeploymentTable,
			Columns: []string{providercommand.DeploymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.DeploymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   providercommand.DeploymentTable,
			Columns: []string{providercommand.DeploymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{providercommand.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pcu.mutation.done = true
	return n, nil
}

// ProviderCommandUpdateOne is the builder for updating a single ProviderCommand entity.
type ProviderCommandUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProviderCommandMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pcuo *ProviderCommandUpdateOne) SetUpdatedAt(t time.Time) *ProviderCommandUpdateOne {
	pcuo.mutation.SetUpdatedAt(t)
	return pcuo
}

// SetCommandType sets the "command_type" field.
func (pcuo *ProviderCommandUpdateOne) SetCommandType(pt providercommand.CommandType) *ProviderCommandUpdateOne {
	pcuo.mutation.SetCommandType(pt)
	return pcuo
}

// SetStatus sets the "status" field.
func (pcuo *ProviderCommandUpdateOne) SetStatus(pr providercommand.Status) *ProviderCommandUpdateOne {
	pcuo.mutation.SetStatus(pr)
	return pcuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pcuo *ProviderCommandUpdateOne) SetNillableStatus(pr *providercommand.Status) *ProviderCommandUpdateOne {
	if pr != nil {
		pcuo.SetStatus(*pr)
	}
	return pcuo
}

// SetArguments sets the "arguments" field.
func (pcuo *ProviderCommandUpdateOne) SetArguments(s []string) *ProviderCommandUpdateOne {
	pcuo.mutation.SetArguments(s)
	return pcuo
}

// AppendArguments appends s to the "arguments" field.
func (pcuo *ProviderCommandUpdateOne) AppendArguments(s []string) *ProviderCommandUpdateOne {
	pcuo.mutation.AppendArguments(s)
	return pcuo
}

// ClearArguments clears the value of the "arguments" field.
func (pcuo *ProviderCommandUpdateOne) ClearArguments() *ProviderCommandUpdateOne {
	pcuo.mutation.ClearArguments()
	return pcuo
}

// SetStartTime sets the "start_time" field.
func (pcuo *ProviderCommandUpdateOne) SetStartTime(t time.Time) *ProviderCommandUpdateOne {
	pcuo.mutation.SetStartTime(t)
	return pcuo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (pcuo *ProviderCommandUpdateOne) SetNillableStartTime(t *time.Time) *ProviderCommandUpdateOne {
	if t != nil {
		pcuo.SetStartTime(*t)
	}
	return pcuo
}

// ClearStartTime clears the value of the "start_time" field.
func (pcuo *ProviderCommandUpdateOne) ClearStartTime() *ProviderCommandUpdateOne {
	pcuo.mutation.ClearStartTime()
	return pcuo
}

// SetEndTime sets the "end_time" field.
func (pcuo *ProviderCommandUpdateOne) SetEndTime(t time.Time) *ProviderCommandUpdateOne {
	pcuo.mutation.SetEndTime(t)
	return pcuo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (pcuo *ProviderCommandUpdateOne) SetNillableEndTime(t *time.Time) *ProviderCommandUpdateOne {
	if t != nil {
		pcuo.SetEndTime(*t)
	}
	return pcuo
}

// ClearEndTime clears the value of the "end_time" field.
func (pcuo *ProviderCommandUpdateOne) ClearEndTime() *ProviderCommandUpdateOne {
	pcuo.mutation.ClearEndTime()
	return pcuo
}

// SetOutput sets the "output" field.
func (pcuo *ProviderCommandUpdateOne) SetOutput(s string) *ProviderCommandUpdateOne {
	pcuo.mutation.SetOutput(s)
	return pcuo
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (pcuo *ProviderCommandUpdateOne) SetNillableOutput(s *string) *ProviderCommandUpdateOne {
	if s != nil {
		pcuo.SetOutput(*s)
	}
	return pcuo
}

// SetError sets the "error" field.
func (pcuo *ProviderCommandUpdateOne) SetError(s string) *ProviderCommandUpdateOne {
	pcuo.mutation.SetError(s)
	return pcuo
}

// SetNillableError sets the "error" field if the given value is not nil.
func (pcuo *ProviderCommandUpdateOne) SetNillableError(s *string) *ProviderCommandUpdateOne {
	if s != nil {
		pcuo.SetError(*s)
	}
	return pcuo
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (pcuo *ProviderCommandUpdateOne) SetProviderID(id uuid.UUID) *ProviderCommandUpdateOne {
	pcuo.mutation.SetProviderID(id)
	return pcuo
}

// SetProvider sets the "provider" edge to the Provider entity.
func (pcuo *ProviderCommandUpdateOne) SetProvider(p *Provider) *ProviderCommandUpdateOne {
	return pcuo.SetProviderID(p.ID)
}

// SetDeploymentID sets the "deployment" edge to the Deployment entity by ID.
func (pcuo *ProviderCommandUpdateOne) SetDeploymentID(id uuid.UUID) *ProviderCommandUpdateOne {
	pcuo.mutation.SetDeploymentID(id)
	return pcuo
}

// SetNillableDeploymentID sets the "deployment" edge to the Deployment entity by ID if the given value is not nil.
func (pcuo *ProviderCommandUpdateOne) SetNillableDeploymentID(id *uuid.UUID) *ProviderCommandUpdateOne {
	if id != nil {
		pcuo = pcuo.SetDeploymentID(*id)
	}
	return pcuo
}

// SetDeployment sets the "deployment" edge to the Deployment entity.
func (pcuo *ProviderCommandUpdateOne) SetDeployment(d *Deployment) *ProviderCommandUpdateOne {
	return pcuo.SetDeploymentID(d.ID)
}

// Mutation returns the ProviderCommandMutation object of the builder.
func (pcuo *ProviderCommandUpdateOne) Mutation() *ProviderCommandMutation {
	return pcuo.mutation
}

// ClearProvider clears the "provider" edge to the Provider entity.
func (pcuo *ProviderCommandUpdateOne) ClearProvider() *ProviderCommandUpdateOne {
	pcuo.mutation.ClearProvider()
	return pcuo
}

// ClearDeployment clears the "deployment" edge to the Deployment entity.
func (pcuo *ProviderCommandUpdateOne) ClearDeployment() *ProviderCommandUpdateOne {
	pcuo.mutation.ClearDeployment()
	return pcuo
}

// Where appends a list predicates to the ProviderCommandUpdate builder.
func (pcuo *ProviderCommandUpdateOne) Where(ps ...predicate.ProviderCommand) *ProviderCommandUpdateOne {
	pcuo.mutation.Where(ps...)
	return pcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcuo *ProviderCommandUpdateOne) Select(field string, fields ...string) *ProviderCommandUpdateOne {
	pcuo.fields = append([]string{field}, fields...)
	return pcuo
}

// Save executes the query and returns the updated ProviderCommand entity.
func (pcuo *ProviderCommandUpdateOne) Save(ctx context.Context) (*ProviderCommand, error) {
	pcuo.defaults()
	return withHooks(ctx, pcuo.sqlSave, pcuo.mutation, pcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcuo *ProviderCommandUpdateOne) SaveX(ctx context.Context) *ProviderCommand {
	node, err := pcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcuo *ProviderCommandUpdateOne) Exec(ctx context.Context) error {
	_, err := pcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcuo *ProviderCommandUpdateOne) ExecX(ctx context.Context) {
	if err := pcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcuo *ProviderCommandUpdateOne) defaults() {
	if _, ok := pcuo.mutation.UpdatedAt(); !ok {
		v := providercommand.UpdateDefaultUpdatedAt()
		pcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcuo *ProviderCommandUpdateOne) check() error {
	if v, ok := pcuo.mutation.CommandType(); ok {
		if err := providercommand.CommandTypeValidator(v); err != nil {
			return &ValidationError{Name: "command_type", err: fmt.Errorf(`ent: validator failed for field "ProviderCommand.command_type": %w`, err)}
		}
	}
	if v, ok := pcuo.mutation.Status(); ok {
		if err := providercommand.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ProviderCommand.status": %w`, err)}
		}
	}
	if _, ok := pcuo.mutation.ProviderID(); pcuo.mutation.ProviderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProviderCommand.provider"`)
	}
	return nil
}

func (pcuo *ProviderCommandUpdateOne) sqlSave(ctx context.Context) (_node *ProviderCommand, err error) {
	if err := pcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(providercommand.Table, providercommand.Columns, sqlgraph.NewFieldSpec(providercommand.FieldID, field.TypeUUID))
	id, ok := pcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProviderCommand.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, providercommand.FieldID)
		for _, f := range fields {
			if !providercommand.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != providercommand.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(providercommand.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pcuo.mutation.CommandType(); ok {
		_spec.SetField(providercommand.FieldCommandType, field.TypeEnum, value)
	}
	if value, ok := pcuo.mutation.Status(); ok {
		_spec.SetField(providercommand.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := pcuo.mutation.Arguments(); ok {
		_spec.SetField(providercommand.FieldArguments, field.TypeJSON, value)
	}
	if value, ok := pcuo.mutation.AppendedArguments(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, providercommand.FieldArguments, value)
		})
	}
	if pcuo.mutation.ArgumentsCleared() {
		_spec.ClearField(providercommand.FieldArguments, field.TypeJSON)
	}
	if value, ok := pcuo.mutation.StartTime(); ok {
		_spec.SetField(providercommand.FieldStartTime, field.TypeTime, value)
	}
	if pcuo.mutation.StartTimeCleared() {
		_spec.ClearField(providercommand.FieldStartTime, field.TypeTime)
	}
	if value, ok := pcuo.mutation.EndTime(); ok {
		_spec.SetField(providercommand.FieldEndTime, field.TypeTime, value)
	}
	if pcuo.mutation.EndTimeCleared() {
		_spec.ClearField(providercommand.FieldEndTime, field.TypeTime)
	}
	if value, ok := pcuo.mutation.Output(); ok {
		_spec.SetField(providercommand.FieldOutput, field.TypeString, value)
	}
	if value, ok := pcuo.mutation.Error(); ok {
		_spec.SetField(providercommand.FieldError, field.TypeString, value)
	}
	if pcuo.mutation.ProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   providercommand.ProviderTable,
			Columns: []string{providercommand.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provider.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   providercommand.ProviderTable,
			Columns: []string{providercommand.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provider.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcuo.mutation.DeploymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   providercommand.DeploymentTable,
			Columns: []string{providercommand.DeploymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.DeploymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   providercommand.DeploymentTable,
			Columns: []string{providercommand.DeploymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProviderCommand{config: pcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{providercommand.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pcuo.mutation.done = true
	return _node, nil
}
