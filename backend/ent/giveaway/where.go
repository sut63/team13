// Code generated by entc, DO NOT EDIT.

package giveaway

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GiveawayName applies equality check predicate on the "giveawayName" field. It's identical to GiveawayNameEQ.
func GiveawayName(v int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGiveawayName), v))
	})
}

// GiveawayNameEQ applies the EQ predicate on the "giveawayName" field.
func GiveawayNameEQ(v int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGiveawayName), v))
	})
}

// GiveawayNameNEQ applies the NEQ predicate on the "giveawayName" field.
func GiveawayNameNEQ(v int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGiveawayName), v))
	})
}

// GiveawayNameIn applies the In predicate on the "giveawayName" field.
func GiveawayNameIn(vs ...int) predicate.Giveaway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Giveaway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGiveawayName), v...))
	})
}

// GiveawayNameNotIn applies the NotIn predicate on the "giveawayName" field.
func GiveawayNameNotIn(vs ...int) predicate.Giveaway {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Giveaway(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGiveawayName), v...))
	})
}

// GiveawayNameGT applies the GT predicate on the "giveawayName" field.
func GiveawayNameGT(v int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGiveawayName), v))
	})
}

// GiveawayNameGTE applies the GTE predicate on the "giveawayName" field.
func GiveawayNameGTE(v int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGiveawayName), v))
	})
}

// GiveawayNameLT applies the LT predicate on the "giveawayName" field.
func GiveawayNameLT(v int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGiveawayName), v))
	})
}

// GiveawayNameLTE applies the LTE predicate on the "giveawayName" field.
func GiveawayNameLTE(v int) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGiveawayName), v))
	})
}

// HasForgiveaway applies the HasEdge predicate on the "forgiveaway" edge.
func HasForgiveaway() predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ForgiveawayTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ForgiveawayTable, ForgiveawayColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasForgiveawayWith applies the HasEdge predicate on the "forgiveaway" edge with a given conditions (other predicates).
func HasForgiveawayWith(preds ...predicate.Promotion) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ForgiveawayInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ForgiveawayTable, ForgiveawayColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Giveaway) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Giveaway) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Giveaway) predicate.Giveaway {
	return predicate.Giveaway(func(s *sql.Selector) {
		p(s.Not())
	})
}
