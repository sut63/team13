// Code generated by entc, DO NOT EDIT.

package promotion

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PromotionName applies equality check predicate on the "PromotionName" field. It's identical to PromotionNameEQ.
func PromotionName(v string) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPromotionName), v))
	})
}

// PromotionNameEQ applies the EQ predicate on the "PromotionName" field.
func PromotionNameEQ(v string) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPromotionName), v))
	})
}

// PromotionNameNEQ applies the NEQ predicate on the "PromotionName" field.
func PromotionNameNEQ(v string) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPromotionName), v))
	})
}

// PromotionNameIn applies the In predicate on the "PromotionName" field.
func PromotionNameIn(vs ...string) predicate.Promotion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Promotion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPromotionName), v...))
	})
}

// PromotionNameNotIn applies the NotIn predicate on the "PromotionName" field.
func PromotionNameNotIn(vs ...string) predicate.Promotion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Promotion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPromotionName), v...))
	})
}

// PromotionNameGT applies the GT predicate on the "PromotionName" field.
func PromotionNameGT(v string) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPromotionName), v))
	})
}

// PromotionNameGTE applies the GTE predicate on the "PromotionName" field.
func PromotionNameGTE(v string) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPromotionName), v))
	})
}

// PromotionNameLT applies the LT predicate on the "PromotionName" field.
func PromotionNameLT(v string) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPromotionName), v))
	})
}

// PromotionNameLTE applies the LTE predicate on the "PromotionName" field.
func PromotionNameLTE(v string) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPromotionName), v))
	})
}

// PromotionNameContains applies the Contains predicate on the "PromotionName" field.
func PromotionNameContains(v string) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPromotionName), v))
	})
}

// PromotionNameHasPrefix applies the HasPrefix predicate on the "PromotionName" field.
func PromotionNameHasPrefix(v string) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPromotionName), v))
	})
}

// PromotionNameHasSuffix applies the HasSuffix predicate on the "PromotionName" field.
func PromotionNameHasSuffix(v string) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPromotionName), v))
	})
}

// PromotionNameEqualFold applies the EqualFold predicate on the "PromotionName" field.
func PromotionNameEqualFold(v string) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPromotionName), v))
	})
}

// PromotionNameContainsFold applies the ContainsFold predicate on the "PromotionName" field.
func PromotionNameContainsFold(v string) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPromotionName), v))
	})
}

// HasSale applies the HasEdge predicate on the "sale" edge.
func HasSale() predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SaleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SaleTable, SaleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSaleWith applies the HasEdge predicate on the "sale" edge with a given conditions (other predicates).
func HasSaleWith(preds ...predicate.Discount) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SaleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SaleTable, SaleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGive applies the HasEdge predicate on the "give" edge.
func HasGive() predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GiveTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GiveTable, GiveColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGiveWith applies the HasEdge predicate on the "give" edge with a given conditions (other predicates).
func HasGiveWith(preds ...predicate.Giveaway) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GiveInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GiveTable, GiveColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Promotion) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Promotion) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
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
func Not(p predicate.Promotion) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		p(s.Not())
	})
}
