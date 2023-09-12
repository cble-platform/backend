// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cble-platform/backend/ent/blueprint"
	"github.com/cble-platform/backend/ent/deployment"
	"github.com/cble-platform/backend/ent/user"
	"github.com/google/uuid"
)

// Deployment is the model entity for the Deployment schema.
type Deployment struct {
	config
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeploymentQuery when eager-loading is set.
	Edges                DeploymentEdges `json:"edges"`
	deployment_blueprint *uuid.UUID
	deployment_requester *uuid.UUID
	selectValues         sql.SelectValues
}

// DeploymentEdges holds the relations/edges for other nodes in the graph.
type DeploymentEdges struct {
	// Blueprint holds the value of the blueprint edge.
	Blueprint *Blueprint `json:"blueprint,omitempty"`
	// Requester holds the value of the requester edge.
	Requester *User `json:"requester,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BlueprintOrErr returns the Blueprint value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeploymentEdges) BlueprintOrErr() (*Blueprint, error) {
	if e.loadedTypes[0] {
		if e.Blueprint == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: blueprint.Label}
		}
		return e.Blueprint, nil
	}
	return nil, &NotLoadedError{edge: "blueprint"}
}

// RequesterOrErr returns the Requester value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeploymentEdges) RequesterOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Requester == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Requester, nil
	}
	return nil, &NotLoadedError{edge: "requester"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Deployment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case deployment.FieldID:
			values[i] = new(uuid.UUID)
		case deployment.ForeignKeys[0]: // deployment_blueprint
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case deployment.ForeignKeys[1]: // deployment_requester
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Deployment fields.
func (d *Deployment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deployment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case deployment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_blueprint", values[i])
			} else if value.Valid {
				d.deployment_blueprint = new(uuid.UUID)
				*d.deployment_blueprint = *value.S.(*uuid.UUID)
			}
		case deployment.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_requester", values[i])
			} else if value.Valid {
				d.deployment_requester = new(uuid.UUID)
				*d.deployment_requester = *value.S.(*uuid.UUID)
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Deployment.
// This includes values selected through modifiers, order, etc.
func (d *Deployment) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryBlueprint queries the "blueprint" edge of the Deployment entity.
func (d *Deployment) QueryBlueprint() *BlueprintQuery {
	return NewDeploymentClient(d.config).QueryBlueprint(d)
}

// QueryRequester queries the "requester" edge of the Deployment entity.
func (d *Deployment) QueryRequester() *UserQuery {
	return NewDeploymentClient(d.config).QueryRequester(d)
}

// Update returns a builder for updating this Deployment.
// Note that you need to call Deployment.Unwrap() before calling this method if this Deployment
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Deployment) Update() *DeploymentUpdateOne {
	return NewDeploymentClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Deployment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Deployment) Unwrap() *Deployment {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Deployment is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Deployment) String() string {
	var builder strings.Builder
	builder.WriteString("Deployment(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Deployments is a parsable slice of Deployment.
type Deployments []*Deployment
