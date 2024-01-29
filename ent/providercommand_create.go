// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/provider"
	"github.com/cble-platform/cble-backend/ent/providercommand"
	"github.com/google/uuid"
)

// ProviderCommandCreate is the builder for creating a ProviderCommand entity.
type ProviderCommandCreate struct {
	config
	mutation *ProviderCommandMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pcc *ProviderCommandCreate) SetCreatedAt(t time.Time) *ProviderCommandCreate {
	pcc.mutation.SetCreatedAt(t)
	return pcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pcc *ProviderCommandCreate) SetNillableCreatedAt(t *time.Time) *ProviderCommandCreate {
	if t != nil {
		pcc.SetCreatedAt(*t)
	}
	return pcc
}

// SetUpdatedAt sets the "updated_at" field.
func (pcc *ProviderCommandCreate) SetUpdatedAt(t time.Time) *ProviderCommandCreate {
	pcc.mutation.SetUpdatedAt(t)
	return pcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pcc *ProviderCommandCreate) SetNillableUpdatedAt(t *time.Time) *ProviderCommandCreate {
	if t != nil {
		pcc.SetUpdatedAt(*t)
	}
	return pcc
}

// SetCommandType sets the "command_type" field.
func (pcc *ProviderCommandCreate) SetCommandType(pt providercommand.CommandType) *ProviderCommandCreate {
	pcc.mutation.SetCommandType(pt)
	return pcc
}

// SetStatus sets the "status" field.
func (pcc *ProviderCommandCreate) SetStatus(pr providercommand.Status) *ProviderCommandCreate {
	pcc.mutation.SetStatus(pr)
	return pcc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pcc *ProviderCommandCreate) SetNillableStatus(pr *providercommand.Status) *ProviderCommandCreate {
	if pr != nil {
		pcc.SetStatus(*pr)
	}
	return pcc
}

// SetArguments sets the "arguments" field.
func (pcc *ProviderCommandCreate) SetArguments(s []string) *ProviderCommandCreate {
	pcc.mutation.SetArguments(s)
	return pcc
}

// SetStartTime sets the "start_time" field.
func (pcc *ProviderCommandCreate) SetStartTime(t time.Time) *ProviderCommandCreate {
	pcc.mutation.SetStartTime(t)
	return pcc
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (pcc *ProviderCommandCreate) SetNillableStartTime(t *time.Time) *ProviderCommandCreate {
	if t != nil {
		pcc.SetStartTime(*t)
	}
	return pcc
}

// SetEndTime sets the "end_time" field.
func (pcc *ProviderCommandCreate) SetEndTime(t time.Time) *ProviderCommandCreate {
	pcc.mutation.SetEndTime(t)
	return pcc
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (pcc *ProviderCommandCreate) SetNillableEndTime(t *time.Time) *ProviderCommandCreate {
	if t != nil {
		pcc.SetEndTime(*t)
	}
	return pcc
}

// SetOutput sets the "output" field.
func (pcc *ProviderCommandCreate) SetOutput(s string) *ProviderCommandCreate {
	pcc.mutation.SetOutput(s)
	return pcc
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (pcc *ProviderCommandCreate) SetNillableOutput(s *string) *ProviderCommandCreate {
	if s != nil {
		pcc.SetOutput(*s)
	}
	return pcc
}

// SetError sets the "error" field.
func (pcc *ProviderCommandCreate) SetError(s string) *ProviderCommandCreate {
	pcc.mutation.SetError(s)
	return pcc
}

// SetNillableError sets the "error" field if the given value is not nil.
func (pcc *ProviderCommandCreate) SetNillableError(s *string) *ProviderCommandCreate {
	if s != nil {
		pcc.SetError(*s)
	}
	return pcc
}

// SetID sets the "id" field.
func (pcc *ProviderCommandCreate) SetID(u uuid.UUID) *ProviderCommandCreate {
	pcc.mutation.SetID(u)
	return pcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pcc *ProviderCommandCreate) SetNillableID(u *uuid.UUID) *ProviderCommandCreate {
	if u != nil {
		pcc.SetID(*u)
	}
	return pcc
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (pcc *ProviderCommandCreate) SetProviderID(id uuid.UUID) *ProviderCommandCreate {
	pcc.mutation.SetProviderID(id)
	return pcc
}

// SetProvider sets the "provider" edge to the Provider entity.
func (pcc *ProviderCommandCreate) SetProvider(p *Provider) *ProviderCommandCreate {
	return pcc.SetProviderID(p.ID)
}

// SetDeploymentID sets the "deployment" edge to the Deployment entity by ID.
func (pcc *ProviderCommandCreate) SetDeploymentID(id uuid.UUID) *ProviderCommandCreate {
	pcc.mutation.SetDeploymentID(id)
	return pcc
}

// SetNillableDeploymentID sets the "deployment" edge to the Deployment entity by ID if the given value is not nil.
func (pcc *ProviderCommandCreate) SetNillableDeploymentID(id *uuid.UUID) *ProviderCommandCreate {
	if id != nil {
		pcc = pcc.SetDeploymentID(*id)
	}
	return pcc
}

// SetDeployment sets the "deployment" edge to the Deployment entity.
func (pcc *ProviderCommandCreate) SetDeployment(d *Deployment) *ProviderCommandCreate {
	return pcc.SetDeploymentID(d.ID)
}

// Mutation returns the ProviderCommandMutation object of the builder.
func (pcc *ProviderCommandCreate) Mutation() *ProviderCommandMutation {
	return pcc.mutation
}

// Save creates the ProviderCommand in the database.
func (pcc *ProviderCommandCreate) Save(ctx context.Context) (*ProviderCommand, error) {
	pcc.defaults()
	return withHooks(ctx, pcc.sqlSave, pcc.mutation, pcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pcc *ProviderCommandCreate) SaveX(ctx context.Context) *ProviderCommand {
	v, err := pcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcc *ProviderCommandCreate) Exec(ctx context.Context) error {
	_, err := pcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcc *ProviderCommandCreate) ExecX(ctx context.Context) {
	if err := pcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcc *ProviderCommandCreate) defaults() {
	if _, ok := pcc.mutation.CreatedAt(); !ok {
		v := providercommand.DefaultCreatedAt()
		pcc.mutation.SetCreatedAt(v)
	}
	if _, ok := pcc.mutation.UpdatedAt(); !ok {
		v := providercommand.DefaultUpdatedAt()
		pcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pcc.mutation.Status(); !ok {
		v := providercommand.DefaultStatus
		pcc.mutation.SetStatus(v)
	}
	if _, ok := pcc.mutation.Output(); !ok {
		v := providercommand.DefaultOutput
		pcc.mutation.SetOutput(v)
	}
	if _, ok := pcc.mutation.Error(); !ok {
		v := providercommand.DefaultError
		pcc.mutation.SetError(v)
	}
	if _, ok := pcc.mutation.ID(); !ok {
		v := providercommand.DefaultID()
		pcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcc *ProviderCommandCreate) check() error {
	if _, ok := pcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProviderCommand.created_at"`)}
	}
	if _, ok := pcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProviderCommand.updated_at"`)}
	}
	if _, ok := pcc.mutation.CommandType(); !ok {
		return &ValidationError{Name: "command_type", err: errors.New(`ent: missing required field "ProviderCommand.command_type"`)}
	}
	if v, ok := pcc.mutation.CommandType(); ok {
		if err := providercommand.CommandTypeValidator(v); err != nil {
			return &ValidationError{Name: "command_type", err: fmt.Errorf(`ent: validator failed for field "ProviderCommand.command_type": %w`, err)}
		}
	}
	if _, ok := pcc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "ProviderCommand.status"`)}
	}
	if v, ok := pcc.mutation.Status(); ok {
		if err := providercommand.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ProviderCommand.status": %w`, err)}
		}
	}
	if _, ok := pcc.mutation.Output(); !ok {
		return &ValidationError{Name: "output", err: errors.New(`ent: missing required field "ProviderCommand.output"`)}
	}
	if _, ok := pcc.mutation.Error(); !ok {
		return &ValidationError{Name: "error", err: errors.New(`ent: missing required field "ProviderCommand.error"`)}
	}
	if _, ok := pcc.mutation.ProviderID(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`ent: missing required edge "ProviderCommand.provider"`)}
	}
	return nil
}

func (pcc *ProviderCommandCreate) sqlSave(ctx context.Context) (*ProviderCommand, error) {
	if err := pcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pcc.mutation.id = &_node.ID
	pcc.mutation.done = true
	return _node, nil
}

func (pcc *ProviderCommandCreate) createSpec() (*ProviderCommand, *sqlgraph.CreateSpec) {
	var (
		_node = &ProviderCommand{config: pcc.config}
		_spec = sqlgraph.NewCreateSpec(providercommand.Table, sqlgraph.NewFieldSpec(providercommand.FieldID, field.TypeUUID))
	)
	if id, ok := pcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pcc.mutation.CreatedAt(); ok {
		_spec.SetField(providercommand.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pcc.mutation.UpdatedAt(); ok {
		_spec.SetField(providercommand.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pcc.mutation.CommandType(); ok {
		_spec.SetField(providercommand.FieldCommandType, field.TypeEnum, value)
		_node.CommandType = value
	}
	if value, ok := pcc.mutation.Status(); ok {
		_spec.SetField(providercommand.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := pcc.mutation.Arguments(); ok {
		_spec.SetField(providercommand.FieldArguments, field.TypeJSON, value)
		_node.Arguments = value
	}
	if value, ok := pcc.mutation.StartTime(); ok {
		_spec.SetField(providercommand.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := pcc.mutation.EndTime(); ok {
		_spec.SetField(providercommand.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := pcc.mutation.Output(); ok {
		_spec.SetField(providercommand.FieldOutput, field.TypeString, value)
		_node.Output = value
	}
	if value, ok := pcc.mutation.Error(); ok {
		_spec.SetField(providercommand.FieldError, field.TypeString, value)
		_node.Error = value
	}
	if nodes := pcc.mutation.ProviderIDs(); len(nodes) > 0 {
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
		_node.provider_command_provider = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcc.mutation.DeploymentIDs(); len(nodes) > 0 {
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
		_node.provider_command_deployment = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProviderCommandCreateBulk is the builder for creating many ProviderCommand entities in bulk.
type ProviderCommandCreateBulk struct {
	config
	err      error
	builders []*ProviderCommandCreate
}

// Save creates the ProviderCommand entities in the database.
func (pccb *ProviderCommandCreateBulk) Save(ctx context.Context) ([]*ProviderCommand, error) {
	if pccb.err != nil {
		return nil, pccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pccb.builders))
	nodes := make([]*ProviderCommand, len(pccb.builders))
	mutators := make([]Mutator, len(pccb.builders))
	for i := range pccb.builders {
		func(i int, root context.Context) {
			builder := pccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProviderCommandMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pccb *ProviderCommandCreateBulk) SaveX(ctx context.Context) []*ProviderCommand {
	v, err := pccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pccb *ProviderCommandCreateBulk) Exec(ctx context.Context) error {
	_, err := pccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pccb *ProviderCommandCreateBulk) ExecX(ctx context.Context) {
	if err := pccb.Exec(ctx); err != nil {
		panic(err)
	}
}
