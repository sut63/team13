// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/assessment"
	"github.com/tanapon395/playlist-video/ent/employee"
	"github.com/tanapon395/playlist-video/ent/position"
	"github.com/tanapon395/playlist-video/ent/salary"
)

// Salary is the model entity for the Salary schema.
type Salary struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Position holds the value of the "position" field.
	Position string `json:"position,omitempty"`
	// Salary holds the value of the "Salary" field.
	Salary int `json:"Salary,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SalaryQuery when eager-loading is set.
	Edges                     SalaryEdges `json:"edges"`
	assessment_formassessment *int
	employee_formemployee     *int
	position_formposition     *int
}

// SalaryEdges holds the relations/edges for other nodes in the graph.
type SalaryEdges struct {
	// Assessment holds the value of the assessment edge.
	Assessment *Assessment
	// Position holds the value of the position edge.
	Position *Position
	// Employee holds the value of the employee edge.
	Employee *Employee
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// AssessmentOrErr returns the Assessment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SalaryEdges) AssessmentOrErr() (*Assessment, error) {
	if e.loadedTypes[0] {
		if e.Assessment == nil {
			// The edge assessment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: assessment.Label}
		}
		return e.Assessment, nil
	}
	return nil, &NotLoadedError{edge: "assessment"}
}

// PositionOrErr returns the Position value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SalaryEdges) PositionOrErr() (*Position, error) {
	if e.loadedTypes[1] {
		if e.Position == nil {
			// The edge position was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: position.Label}
		}
		return e.Position, nil
	}
	return nil, &NotLoadedError{edge: "position"}
}

// EmployeeOrErr returns the Employee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SalaryEdges) EmployeeOrErr() (*Employee, error) {
	if e.loadedTypes[2] {
		if e.Employee == nil {
			// The edge employee was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: employee.Label}
		}
		return e.Employee, nil
	}
	return nil, &NotLoadedError{edge: "employee"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Salary) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // position
		&sql.NullInt64{},  // Salary
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Salary) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // assessment_formassessment
		&sql.NullInt64{}, // employee_formemployee
		&sql.NullInt64{}, // position_formposition
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Salary fields.
func (s *Salary) assignValues(values ...interface{}) error {
	if m, n := len(values), len(salary.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		s.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field position", values[1])
	} else if value.Valid {
		s.Position = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Salary", values[2])
	} else if value.Valid {
		s.Salary = int(value.Int64)
	}
	values = values[3:]
	if len(values) == len(salary.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field assessment_formassessment", value)
		} else if value.Valid {
			s.assessment_formassessment = new(int)
			*s.assessment_formassessment = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field employee_formemployee", value)
		} else if value.Valid {
			s.employee_formemployee = new(int)
			*s.employee_formemployee = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field position_formposition", value)
		} else if value.Valid {
			s.position_formposition = new(int)
			*s.position_formposition = int(value.Int64)
		}
	}
	return nil
}

// QueryAssessment queries the assessment edge of the Salary.
func (s *Salary) QueryAssessment() *AssessmentQuery {
	return (&SalaryClient{config: s.config}).QueryAssessment(s)
}

// QueryPosition queries the position edge of the Salary.
func (s *Salary) QueryPosition() *PositionQuery {
	return (&SalaryClient{config: s.config}).QueryPosition(s)
}

// QueryEmployee queries the employee edge of the Salary.
func (s *Salary) QueryEmployee() *EmployeeQuery {
	return (&SalaryClient{config: s.config}).QueryEmployee(s)
}

// Update returns a builder for updating this Salary.
// Note that, you need to call Salary.Unwrap() before calling this method, if this Salary
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Salary) Update() *SalaryUpdateOne {
	return (&SalaryClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Salary) Unwrap() *Salary {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Salary is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Salary) String() string {
	var builder strings.Builder
	builder.WriteString("Salary(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteString(", position=")
	builder.WriteString(s.Position)
	builder.WriteString(", Salary=")
	builder.WriteString(fmt.Sprintf("%v", s.Salary))
	builder.WriteByte(')')
	return builder.String()
}

// Salaries is a parsable slice of Salary.
type Salaries []*Salary

func (s Salaries) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
