// Code generated by ent, DO NOT EDIT.

package deployment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cble-platform/backend/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Deployment {
	return predicate.Deployment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Deployment {
	return predicate.Deployment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(sql.FieldLTE(FieldID, id))
}

// HasBlueprint applies the HasEdge predicate on the "blueprint" edge.
func HasBlueprint() predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BlueprintTable, BlueprintColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlueprintWith applies the HasEdge predicate on the "blueprint" edge with a given conditions (other predicates).
func HasBlueprintWith(preds ...predicate.Blueprint) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		step := newBlueprintStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRequester applies the HasEdge predicate on the "requester" edge.
func HasRequester() predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RequesterTable, RequesterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRequesterWith applies the HasEdge predicate on the "requester" edge with a given conditions (other predicates).
func HasRequesterWith(preds ...predicate.User) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		step := newRequesterStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Deployment) predicate.Deployment {
	return predicate.Deployment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Deployment) predicate.Deployment {
	return predicate.Deployment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Deployment) predicate.Deployment {
	return predicate.Deployment(sql.NotPredicates(p))
}
