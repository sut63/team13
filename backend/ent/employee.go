// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team13/app/ent/employee"
)

// Employee is the model entity for the Employee schema.
type Employee struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EmployeeQuery when eager-loading is set.
	Edges EmployeeEdges `json:"edges"`
}

// EmployeeEdges holds the relations/edges for other nodes in the graph.
type EmployeeEdges struct {
	// Whose holds the value of the whose edge.
	Whose []*EmployeeWorkingHours
	// Employeestock holds the value of the employeestock edge.
	Employeestock []*Stock
	// Formemployee holds the value of the formemployee edge.
	Formemployee []*Salary
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// WhoseOrErr returns the Whose value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) WhoseOrErr() ([]*EmployeeWorkingHours, error) {
	if e.loadedTypes[0] {
		return e.Whose, nil
	}
	return nil, &NotLoadedError{edge: "whose"}
}

// EmployeestockOrErr returns the Employeestock value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) EmployeestockOrErr() ([]*Stock, error) {
	if e.loadedTypes[1] {
		return e.Employeestock, nil
	}
	return nil, &NotLoadedError{edge: "employeestock"}
}

// FormemployeeOrErr returns the Formemployee value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) FormemployeeOrErr() ([]*Salary, error) {
	if e.loadedTypes[2] {
		return e.Formemployee, nil
	}
	return nil, &NotLoadedError{edge: "formemployee"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Employee) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // email
		&sql.NullString{}, // password
		&sql.NullInt64{},  // age
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Employee fields.
func (e *Employee) assignValues(values ...interface{}) error {
	if m, n := len(values), len(employee.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	e.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		e.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[1])
	} else if value.Valid {
		e.Email = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[2])
	} else if value.Valid {
		e.Password = value.String
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[3])
	} else if value.Valid {
		e.Age = int(value.Int64)
	}
	return nil
}

// QueryWhose queries the whose edge of the Employee.
func (e *Employee) QueryWhose() *EmployeeWorkingHoursQuery {
	return (&EmployeeClient{config: e.config}).QueryWhose(e)
}

// QueryEmployeestock queries the employeestock edge of the Employee.
func (e *Employee) QueryEmployeestock() *StockQuery {
	return (&EmployeeClient{config: e.config}).QueryEmployeestock(e)
}

// QueryFormemployee queries the formemployee edge of the Employee.
func (e *Employee) QueryFormemployee() *SalaryQuery {
	return (&EmployeeClient{config: e.config}).QueryFormemployee(e)
}

// Update returns a builder for updating this Employee.
// Note that, you need to call Employee.Unwrap() before calling this method, if this Employee
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Employee) Update() *EmployeeUpdateOne {
	return (&EmployeeClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (e *Employee) Unwrap() *Employee {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Employee is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Employee) String() string {
	var builder strings.Builder
	builder.WriteString("Employee(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", name=")
	builder.WriteString(e.Name)
	builder.WriteString(", email=")
	builder.WriteString(e.Email)
	builder.WriteString(", password=")
	builder.WriteString(e.Password)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", e.Age))
	builder.WriteByte(')')
	return builder.String()
}

// Employees is a parsable slice of Employee.
type Employees []*Employee

func (e Employees) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
