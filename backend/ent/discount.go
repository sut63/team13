// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team13/app/ent/discount"
)

// Discount is the model entity for the Discount schema.
type Discount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Sale holds the value of the "sale" field.
	Sale int `json:"sale,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DiscountQuery when eager-loading is set.
	Edges DiscountEdges `json:"edges"`
}

// DiscountEdges holds the relations/edges for other nodes in the graph.
type DiscountEdges struct {
	// Fordiscount holds the value of the fordiscount edge.
	Fordiscount []*Promotion
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FordiscountOrErr returns the Fordiscount value or an error if the edge
// was not loaded in eager-loading.
func (e DiscountEdges) FordiscountOrErr() ([]*Promotion, error) {
	if e.loadedTypes[0] {
		return e.Fordiscount, nil
	}
	return nil, &NotLoadedError{edge: "fordiscount"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Discount) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // sale
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Discount fields.
func (d *Discount) assignValues(values ...interface{}) error {
	if m, n := len(values), len(discount.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field sale", values[0])
	} else if value.Valid {
		d.Sale = int(value.Int64)
	}
	return nil
}

// QueryFordiscount queries the fordiscount edge of the Discount.
func (d *Discount) QueryFordiscount() *PromotionQuery {
	return (&DiscountClient{config: d.config}).QueryFordiscount(d)
}

// Update returns a builder for updating this Discount.
// Note that, you need to call Discount.Unwrap() before calling this method, if this Discount
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Discount) Update() *DiscountUpdateOne {
	return (&DiscountClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Discount) Unwrap() *Discount {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Discount is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Discount) String() string {
	var builder strings.Builder
	builder.WriteString("Discount(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", sale=")
	builder.WriteString(fmt.Sprintf("%v", d.Sale))
	builder.WriteByte(')')
	return builder.String()
}

// Discounts is a parsable slice of Discount.
type Discounts []*Discount

func (d Discounts) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
