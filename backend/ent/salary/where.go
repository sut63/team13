// Code generated by entc, DO NOT EDIT.

package salary

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team13/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Salary applies equality check predicate on the "Salary" field. It's identical to SalaryEQ.
func Salary(v int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalary), v))
	})
}

// SalaryEQ applies the EQ predicate on the "Salary" field.
func SalaryEQ(v int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalary), v))
	})
}

// SalaryNEQ applies the NEQ predicate on the "Salary" field.
func SalaryNEQ(v int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSalary), v))
	})
}

// SalaryIn applies the In predicate on the "Salary" field.
func SalaryIn(vs ...int) predicate.Salary {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Salary(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSalary), v...))
	})
}

// SalaryNotIn applies the NotIn predicate on the "Salary" field.
func SalaryNotIn(vs ...int) predicate.Salary {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Salary(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSalary), v...))
	})
}

// SalaryGT applies the GT predicate on the "Salary" field.
func SalaryGT(v int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSalary), v))
	})
}

// SalaryGTE applies the GTE predicate on the "Salary" field.
func SalaryGTE(v int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSalary), v))
	})
}

// SalaryLT applies the LT predicate on the "Salary" field.
func SalaryLT(v int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSalary), v))
	})
}

// SalaryLTE applies the LTE predicate on the "Salary" field.
func SalaryLTE(v int) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSalary), v))
	})
}

// HasAssessment applies the HasEdge predicate on the "assessment" edge.
func HasAssessment() predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssessmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AssessmentTable, AssessmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssessmentWith applies the HasEdge predicate on the "assessment" edge with a given conditions (other predicates).
func HasAssessmentWith(preds ...predicate.Assessment) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssessmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AssessmentTable, AssessmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosition applies the HasEdge predicate on the "position" edge.
func HasPosition() predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PositionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PositionTable, PositionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPositionWith applies the HasEdge predicate on the "position" edge with a given conditions (other predicates).
func HasPositionWith(preds ...predicate.Position) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PositionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PositionTable, PositionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEmployee applies the HasEdge predicate on the "employee" edge.
func HasEmployee() predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Salary) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Salary) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
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
func Not(p predicate.Salary) predicate.Salary {
	return predicate.Salary(func(s *sql.Selector) {
		p(s.Not())
	})
}
