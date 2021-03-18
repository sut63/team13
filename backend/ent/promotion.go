// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team13/app/ent/discount"
	"github.com/team13/app/ent/giveaway"
	"github.com/team13/app/ent/product"
	"github.com/team13/app/ent/promotion"
)

// Promotion is the model entity for the Promotion schema.
type Promotion struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PromotionName holds the value of the "PromotionName" field.
	PromotionName string `json:"PromotionName,omitempty"`
	// StartPromotion holds the value of the "StartPromotion" field.
	StartPromotion time.Time `json:"StartPromotion,omitempty"`
	// EndPromotion holds the value of the "EndPromotion" field.
	EndPromotion time.Time `json:"EndPromotion,omitempty"`
	// Price holds the value of the "Price" field.
	Price float64 `json:"Price,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PromotionQuery when eager-loading is set.
	Edges                PromotionEdges `json:"edges"`
	discount_fordiscount *int
	giveaway_forgiveaway *int
	product_forproduct   *int
}

// PromotionEdges holds the relations/edges for other nodes in the graph.
type PromotionEdges struct {
	// Sale holds the value of the sale edge.
	Sale *Discount
	// Give holds the value of the give edge.
	Give *Giveaway
	// Product holds the value of the product edge.
	Product *Product
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SaleOrErr returns the Sale value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PromotionEdges) SaleOrErr() (*Discount, error) {
	if e.loadedTypes[0] {
		if e.Sale == nil {
			// The edge sale was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: discount.Label}
		}
		return e.Sale, nil
	}
	return nil, &NotLoadedError{edge: "sale"}
}

// GiveOrErr returns the Give value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PromotionEdges) GiveOrErr() (*Giveaway, error) {
	if e.loadedTypes[1] {
		if e.Give == nil {
			// The edge give was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: giveaway.Label}
		}
		return e.Give, nil
	}
	return nil, &NotLoadedError{edge: "give"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PromotionEdges) ProductOrErr() (*Product, error) {
	if e.loadedTypes[2] {
		if e.Product == nil {
			// The edge product was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Promotion) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},   // id
		&sql.NullString{},  // PromotionName
		&sql.NullTime{},    // StartPromotion
		&sql.NullTime{},    // EndPromotion
		&sql.NullFloat64{}, // Price
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Promotion) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // discount_fordiscount
		&sql.NullInt64{}, // giveaway_forgiveaway
		&sql.NullInt64{}, // product_forproduct
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Promotion fields.
func (pr *Promotion) assignValues(values ...interface{}) error {
	if m, n := len(values), len(promotion.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PromotionName", values[0])
	} else if value.Valid {
		pr.PromotionName = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field StartPromotion", values[1])
	} else if value.Valid {
		pr.StartPromotion = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field EndPromotion", values[2])
	} else if value.Valid {
		pr.EndPromotion = value.Time
	}
	if value, ok := values[3].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field Price", values[3])
	} else if value.Valid {
		pr.Price = value.Float64
	}
	values = values[4:]
	if len(values) == len(promotion.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field discount_fordiscount", value)
		} else if value.Valid {
			pr.discount_fordiscount = new(int)
			*pr.discount_fordiscount = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field giveaway_forgiveaway", value)
		} else if value.Valid {
			pr.giveaway_forgiveaway = new(int)
			*pr.giveaway_forgiveaway = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field product_forproduct", value)
		} else if value.Valid {
			pr.product_forproduct = new(int)
			*pr.product_forproduct = int(value.Int64)
		}
	}
	return nil
}

// QuerySale queries the sale edge of the Promotion.
func (pr *Promotion) QuerySale() *DiscountQuery {
	return (&PromotionClient{config: pr.config}).QuerySale(pr)
}

// QueryGive queries the give edge of the Promotion.
func (pr *Promotion) QueryGive() *GiveawayQuery {
	return (&PromotionClient{config: pr.config}).QueryGive(pr)
}

// QueryProduct queries the product edge of the Promotion.
func (pr *Promotion) QueryProduct() *ProductQuery {
	return (&PromotionClient{config: pr.config}).QueryProduct(pr)
}

// Update returns a builder for updating this Promotion.
// Note that, you need to call Promotion.Unwrap() before calling this method, if this Promotion
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Promotion) Update() *PromotionUpdateOne {
	return (&PromotionClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pr *Promotion) Unwrap() *Promotion {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Promotion is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Promotion) String() string {
	var builder strings.Builder
	builder.WriteString("Promotion(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", PromotionName=")
	builder.WriteString(pr.PromotionName)
	builder.WriteString(", StartPromotion=")
	builder.WriteString(pr.StartPromotion.Format(time.ANSIC))
	builder.WriteString(", EndPromotion=")
	builder.WriteString(pr.EndPromotion.Format(time.ANSIC))
	builder.WriteString(", Price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteByte(')')
	return builder.String()
}

// Promotions is a parsable slice of Promotion.
type Promotions []*Promotion

func (pr Promotions) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
