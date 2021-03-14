// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team13/app/ent/startwork"
)

// StartWork is the model entity for the StartWork schema.
type StartWork struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StartWork holds the value of the "StartWork" field.
	StartWork time.Time `json:"StartWork,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StartWorkQuery when eager-loading is set.
	Edges StartWorkEdges `json:"edges"`
}

// StartWorkEdges holds the relations/edges for other nodes in the graph.
type StartWorkEdges struct {
	// Whenwork holds the value of the whenwork edge.
	Whenwork []*EmployeeWorkingHours
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WhenworkOrErr returns the Whenwork value or an error if the edge
// was not loaded in eager-loading.
func (e StartWorkEdges) WhenworkOrErr() ([]*EmployeeWorkingHours, error) {
	if e.loadedTypes[0] {
		return e.Whenwork, nil
	}
	return nil, &NotLoadedError{edge: "whenwork"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StartWork) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // StartWork
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StartWork fields.
func (sw *StartWork) assignValues(values ...interface{}) error {
	if m, n := len(values), len(startwork.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	sw.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field StartWork", values[0])
	} else if value.Valid {
		sw.StartWork = value.Time
	}
	return nil
}

// QueryWhenwork queries the whenwork edge of the StartWork.
func (sw *StartWork) QueryWhenwork() *EmployeeWorkingHoursQuery {
	return (&StartWorkClient{config: sw.config}).QueryWhenwork(sw)
}

// Update returns a builder for updating this StartWork.
// Note that, you need to call StartWork.Unwrap() before calling this method, if this StartWork
// was returned from a transaction, and the transaction was committed or rolled back.
func (sw *StartWork) Update() *StartWorkUpdateOne {
	return (&StartWorkClient{config: sw.config}).UpdateOne(sw)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (sw *StartWork) Unwrap() *StartWork {
	tx, ok := sw.config.driver.(*txDriver)
	if !ok {
		panic("ent: StartWork is not a transactional entity")
	}
	sw.config.driver = tx.drv
	return sw
}

// String implements the fmt.Stringer.
func (sw *StartWork) String() string {
	var builder strings.Builder
	builder.WriteString("StartWork(")
	builder.WriteString(fmt.Sprintf("id=%v", sw.ID))
	builder.WriteString(", StartWork=")
	builder.WriteString(sw.StartWork.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// StartWorks is a parsable slice of StartWork.
type StartWorks []*StartWork

func (sw StartWorks) config(cfg config) {
	for _i := range sw {
		sw[_i].config = cfg
	}
}