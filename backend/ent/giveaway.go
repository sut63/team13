// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team13/app/ent/giveaway"
)

// Giveaway is the model entity for the Giveaway schema.
type Giveaway struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// GiveawayName holds the value of the "giveawayName" field.
	GiveawayName string `json:"giveawayName,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GiveawayQuery when eager-loading is set.
	Edges GiveawayEdges `json:"edges"`
}

// GiveawayEdges holds the relations/edges for other nodes in the graph.
type GiveawayEdges struct {
	// Forgiveaway holds the value of the forgiveaway edge.
	Forgiveaway []*Promotion
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ForgiveawayOrErr returns the Forgiveaway value or an error if the edge
// was not loaded in eager-loading.
func (e GiveawayEdges) ForgiveawayOrErr() ([]*Promotion, error) {
	if e.loadedTypes[0] {
		return e.Forgiveaway, nil
	}
	return nil, &NotLoadedError{edge: "forgiveaway"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Giveaway) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // giveawayName
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Giveaway fields.
func (gi *Giveaway) assignValues(values ...interface{}) error {
	if m, n := len(values), len(giveaway.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	gi.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field giveawayName", values[0])
	} else if value.Valid {
		gi.GiveawayName = value.String
	}
	return nil
}

// QueryForgiveaway queries the forgiveaway edge of the Giveaway.
func (gi *Giveaway) QueryForgiveaway() *PromotionQuery {
	return (&GiveawayClient{config: gi.config}).QueryForgiveaway(gi)
}

// Update returns a builder for updating this Giveaway.
// Note that, you need to call Giveaway.Unwrap() before calling this method, if this Giveaway
// was returned from a transaction, and the transaction was committed or rolled back.
func (gi *Giveaway) Update() *GiveawayUpdateOne {
	return (&GiveawayClient{config: gi.config}).UpdateOne(gi)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gi *Giveaway) Unwrap() *Giveaway {
	tx, ok := gi.config.driver.(*txDriver)
	if !ok {
		panic("ent: Giveaway is not a transactional entity")
	}
	gi.config.driver = tx.drv
	return gi
}

// String implements the fmt.Stringer.
func (gi *Giveaway) String() string {
	var builder strings.Builder
	builder.WriteString("Giveaway(")
	builder.WriteString(fmt.Sprintf("id=%v", gi.ID))
	builder.WriteString(", giveawayName=")
	builder.WriteString(gi.GiveawayName)
	builder.WriteByte(')')
	return builder.String()
}

// Giveaways is a parsable slice of Giveaway.
type Giveaways []*Giveaway

func (gi Giveaways) config(cfg config) {
	for _i := range gi {
		gi[_i].config = cfg
	}
}
