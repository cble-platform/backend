// Code generated by ent, DO NOT EDIT.

package blueprint

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cble-platform/backend/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldName, v))
}

// BlueprintTemplate applies equality check predicate on the "blueprint_template" field. It's identical to BlueprintTemplateEQ.
func BlueprintTemplate(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldBlueprintTemplate, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldContainsFold(FieldName, v))
}

// BlueprintTemplateEQ applies the EQ predicate on the "blueprint_template" field.
func BlueprintTemplateEQ(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldBlueprintTemplate, v))
}

// BlueprintTemplateNEQ applies the NEQ predicate on the "blueprint_template" field.
func BlueprintTemplateNEQ(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNEQ(FieldBlueprintTemplate, v))
}

// BlueprintTemplateIn applies the In predicate on the "blueprint_template" field.
func BlueprintTemplateIn(vs ...[]byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldIn(FieldBlueprintTemplate, vs...))
}

// BlueprintTemplateNotIn applies the NotIn predicate on the "blueprint_template" field.
func BlueprintTemplateNotIn(vs ...[]byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNotIn(FieldBlueprintTemplate, vs...))
}

// BlueprintTemplateGT applies the GT predicate on the "blueprint_template" field.
func BlueprintTemplateGT(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGT(FieldBlueprintTemplate, v))
}

// BlueprintTemplateGTE applies the GTE predicate on the "blueprint_template" field.
func BlueprintTemplateGTE(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGTE(FieldBlueprintTemplate, v))
}

// BlueprintTemplateLT applies the LT predicate on the "blueprint_template" field.
func BlueprintTemplateLT(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLT(FieldBlueprintTemplate, v))
}

// BlueprintTemplateLTE applies the LTE predicate on the "blueprint_template" field.
func BlueprintTemplateLTE(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLTE(FieldBlueprintTemplate, v))
}

// HasParentGroup applies the HasEdge predicate on the "parent_group" edge.
func HasParentGroup() predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ParentGroupTable, ParentGroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentGroupWith applies the HasEdge predicate on the "parent_group" edge with a given conditions (other predicates).
func HasParentGroupWith(preds ...predicate.Group) predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := newParentGroupStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVirtualizationProvider applies the HasEdge predicate on the "virtualization_provider" edge.
func HasVirtualizationProvider() predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, VirtualizationProviderTable, VirtualizationProviderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVirtualizationProviderWith applies the HasEdge predicate on the "virtualization_provider" edge with a given conditions (other predicates).
func HasVirtualizationProviderWith(preds ...predicate.VirtualizationProvider) predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := newVirtualizationProviderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDeployments applies the HasEdge predicate on the "deployments" edge.
func HasDeployments() predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, DeploymentsTable, DeploymentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeploymentsWith applies the HasEdge predicate on the "deployments" edge with a given conditions (other predicates).
func HasDeploymentsWith(preds ...predicate.Deployment) predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := newDeploymentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Blueprint) predicate.Blueprint {
	return predicate.Blueprint(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Blueprint) predicate.Blueprint {
	return predicate.Blueprint(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Blueprint) predicate.Blueprint {
	return predicate.Blueprint(sql.NotPredicates(p))
}
