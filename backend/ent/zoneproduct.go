// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team13/app/ent/zoneproduct"
)

// Zoneproduct is the model entity for the Zoneproduct schema.
type Zoneproduct struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Zone holds the value of the "Zone" field.
	Zone string `json:"Zone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ZoneproductQuery when eager-loading is set.
	Edges ZoneproductEdges `json:"edges"`
}

// ZoneproductEdges holds the relations/edges for other nodes in the graph.
type ZoneproductEdges struct {
	// Zonestock holds the value of the zonestock edge.
	Zonestock []*Stock
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ZonestockOrErr returns the Zonestock value or an error if the edge
// was not loaded in eager-loading.
func (e ZoneproductEdges) ZonestockOrErr() ([]*Stock, error) {
	if e.loadedTypes[0] {
		return e.Zonestock, nil
	}
	return nil, &NotLoadedError{edge: "zonestock"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Zoneproduct) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Zone
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Zoneproduct fields.
func (z *Zoneproduct) assignValues(values ...interface{}) error {
	if m, n := len(values), len(zoneproduct.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	z.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Zone", values[0])
	} else if value.Valid {
		z.Zone = value.String
	}
	return nil
}

// QueryZonestock queries the zonestock edge of the Zoneproduct.
func (z *Zoneproduct) QueryZonestock() *StockQuery {
	return (&ZoneproductClient{config: z.config}).QueryZonestock(z)
}

// Update returns a builder for updating this Zoneproduct.
// Note that, you need to call Zoneproduct.Unwrap() before calling this method, if this Zoneproduct
// was returned from a transaction, and the transaction was committed or rolled back.
func (z *Zoneproduct) Update() *ZoneproductUpdateOne {
	return (&ZoneproductClient{config: z.config}).UpdateOne(z)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (z *Zoneproduct) Unwrap() *Zoneproduct {
	tx, ok := z.config.driver.(*txDriver)
	if !ok {
		panic("ent: Zoneproduct is not a transactional entity")
	}
	z.config.driver = tx.drv
	return z
}

// String implements the fmt.Stringer.
func (z *Zoneproduct) String() string {
	var builder strings.Builder
	builder.WriteString("Zoneproduct(")
	builder.WriteString(fmt.Sprintf("id=%v", z.ID))
	builder.WriteString(", Zone=")
	builder.WriteString(z.Zone)
	builder.WriteByte(')')
	return builder.String()
}

// Zoneproducts is a parsable slice of Zoneproduct.
type Zoneproducts []*Zoneproduct

func (z Zoneproducts) config(cfg config) {
	for _i := range z {
		z[_i].config = cfg
	}
}
