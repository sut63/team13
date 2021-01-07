// Code generated by entc, DO NOT EDIT.

package assessment

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Assessment applies equality check predicate on the "assessment" field. It's identical to AssessmentEQ.
func Assessment(v string) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssessment), v))
	})
}

// AssessmentEQ applies the EQ predicate on the "assessment" field.
func AssessmentEQ(v string) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssessment), v))
	})
}

// AssessmentNEQ applies the NEQ predicate on the "assessment" field.
func AssessmentNEQ(v string) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAssessment), v))
	})
}

// AssessmentIn applies the In predicate on the "assessment" field.
func AssessmentIn(vs ...string) predicate.Assessment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assessment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAssessment), v...))
	})
}

// AssessmentNotIn applies the NotIn predicate on the "assessment" field.
func AssessmentNotIn(vs ...string) predicate.Assessment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assessment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAssessment), v...))
	})
}

// AssessmentGT applies the GT predicate on the "assessment" field.
func AssessmentGT(v string) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAssessment), v))
	})
}

// AssessmentGTE applies the GTE predicate on the "assessment" field.
func AssessmentGTE(v string) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAssessment), v))
	})
}

// AssessmentLT applies the LT predicate on the "assessment" field.
func AssessmentLT(v string) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAssessment), v))
	})
}

// AssessmentLTE applies the LTE predicate on the "assessment" field.
func AssessmentLTE(v string) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAssessment), v))
	})
}

// AssessmentContains applies the Contains predicate on the "assessment" field.
func AssessmentContains(v string) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAssessment), v))
	})
}

// AssessmentHasPrefix applies the HasPrefix predicate on the "assessment" field.
func AssessmentHasPrefix(v string) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAssessment), v))
	})
}

// AssessmentHasSuffix applies the HasSuffix predicate on the "assessment" field.
func AssessmentHasSuffix(v string) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAssessment), v))
	})
}

// AssessmentEqualFold applies the EqualFold predicate on the "assessment" field.
func AssessmentEqualFold(v string) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAssessment), v))
	})
}

// AssessmentContainsFold applies the ContainsFold predicate on the "assessment" field.
func AssessmentContainsFold(v string) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAssessment), v))
	})
}

// HasFormassessment applies the HasEdge predicate on the "formassessment" edge.
func HasFormassessment() predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FormassessmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FormassessmentTable, FormassessmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFormassessmentWith applies the HasEdge predicate on the "formassessment" edge with a given conditions (other predicates).
func HasFormassessmentWith(preds ...predicate.Salary) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FormassessmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FormassessmentTable, FormassessmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Assessment) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Assessment) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
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
func Not(p predicate.Assessment) predicate.Assessment {
	return predicate.Assessment(func(s *sql.Selector) {
		p(s.Not())
	})
}