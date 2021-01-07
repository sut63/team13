// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/day"
	"github.com/tanapon395/playlist-video/ent/employee"
	"github.com/tanapon395/playlist-video/ent/employeeworkinghours"
	"github.com/tanapon395/playlist-video/ent/role"
	"github.com/tanapon395/playlist-video/ent/shift"
)

// Employeeworkinghours is the model entity for the Employeeworkinghours schema.
type Employeeworkinghours struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EmployeeworkinghoursQuery when eager-loading is set.
	Edges          EmployeeworkinghoursEdges `json:"edges"`
	day_whatday    *int
	employee_whose *int
	role_todo      *int
	shift_when     *int
}

// EmployeeworkinghoursEdges holds the relations/edges for other nodes in the graph.
type EmployeeworkinghoursEdges struct {
	// Workinghour holds the value of the workinghour edge.
	Workinghour *Employee
	// Day holds the value of the day edge.
	Day *Day
	// Shift holds the value of the shift edge.
	Shift *Shift
	// Role holds the value of the role edge.
	Role *Role
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// WorkinghourOrErr returns the Workinghour value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EmployeeworkinghoursEdges) WorkinghourOrErr() (*Employee, error) {
	if e.loadedTypes[0] {
		if e.Workinghour == nil {
			// The edge workinghour was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: employee.Label}
		}
		return e.Workinghour, nil
	}
	return nil, &NotLoadedError{edge: "workinghour"}
}

// DayOrErr returns the Day value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EmployeeworkinghoursEdges) DayOrErr() (*Day, error) {
	if e.loadedTypes[1] {
		if e.Day == nil {
			// The edge day was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: day.Label}
		}
		return e.Day, nil
	}
	return nil, &NotLoadedError{edge: "day"}
}

// ShiftOrErr returns the Shift value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EmployeeworkinghoursEdges) ShiftOrErr() (*Shift, error) {
	if e.loadedTypes[2] {
		if e.Shift == nil {
			// The edge shift was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: shift.Label}
		}
		return e.Shift, nil
	}
	return nil, &NotLoadedError{edge: "shift"}
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EmployeeworkinghoursEdges) RoleOrErr() (*Role, error) {
	if e.loadedTypes[3] {
		if e.Role == nil {
			// The edge role was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: role.Label}
		}
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Employeeworkinghours) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Employeeworkinghours) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // day_whatday
		&sql.NullInt64{}, // employee_whose
		&sql.NullInt64{}, // role_todo
		&sql.NullInt64{}, // shift_when
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Employeeworkinghours fields.
func (e *Employeeworkinghours) assignValues(values ...interface{}) error {
	if m, n := len(values), len(employeeworkinghours.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	e.ID = int(value.Int64)
	values = values[1:]
	values = values[0:]
	if len(values) == len(employeeworkinghours.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field day_whatday", value)
		} else if value.Valid {
			e.day_whatday = new(int)
			*e.day_whatday = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field employee_whose", value)
		} else if value.Valid {
			e.employee_whose = new(int)
			*e.employee_whose = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field role_todo", value)
		} else if value.Valid {
			e.role_todo = new(int)
			*e.role_todo = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field shift_when", value)
		} else if value.Valid {
			e.shift_when = new(int)
			*e.shift_when = int(value.Int64)
		}
	}
	return nil
}

// QueryWorkinghour queries the workinghour edge of the Employeeworkinghours.
func (e *Employeeworkinghours) QueryWorkinghour() *EmployeeQuery {
	return (&EmployeeworkinghoursClient{config: e.config}).QueryWorkinghour(e)
}

// QueryDay queries the day edge of the Employeeworkinghours.
func (e *Employeeworkinghours) QueryDay() *DayQuery {
	return (&EmployeeworkinghoursClient{config: e.config}).QueryDay(e)
}

// QueryShift queries the shift edge of the Employeeworkinghours.
func (e *Employeeworkinghours) QueryShift() *ShiftQuery {
	return (&EmployeeworkinghoursClient{config: e.config}).QueryShift(e)
}

// QueryRole queries the role edge of the Employeeworkinghours.
func (e *Employeeworkinghours) QueryRole() *RoleQuery {
	return (&EmployeeworkinghoursClient{config: e.config}).QueryRole(e)
}

// Update returns a builder for updating this Employeeworkinghours.
// Note that, you need to call Employeeworkinghours.Unwrap() before calling this method, if this Employeeworkinghours
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Employeeworkinghours) Update() *EmployeeworkinghoursUpdateOne {
	return (&EmployeeworkinghoursClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (e *Employeeworkinghours) Unwrap() *Employeeworkinghours {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Employeeworkinghours is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Employeeworkinghours) String() string {
	var builder strings.Builder
	builder.WriteString("Employeeworkinghours(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteByte(')')
	return builder.String()
}

// EmployeeworkinghoursSlice is a parsable slice of Employeeworkinghours.
type EmployeeworkinghoursSlice []*Employeeworkinghours

func (e EmployeeworkinghoursSlice) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
