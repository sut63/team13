// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team13/app/ent/product"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NameProduct holds the value of the "NameProduct" field.
	NameProduct string `json:"NameProduct,omitempty"`
	// BarcodeProduct holds the value of the "BarcodeProduct" field.
	BarcodeProduct string `json:"BarcodeProduct,omitempty"`
	// MFG holds the value of the "MFG" field.
	MFG string `json:"MFG,omitempty"`
	// EXP holds the value of the "EXP" field.
	EXP string `json:"EXP,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges ProductEdges `json:"edges"`
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// Products holds the value of the products edge.
	Products []*Orderproduct
	// Stockproduct holds the value of the stockproduct edge.
	Stockproduct []*Stock
	// Forproduct holds the value of the forproduct edge.
	Forproduct []*Promotion
	// Formproductonline holds the value of the formproductonline edge.
	Formproductonline []*Orderonline
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ProductsOrErr() ([]*Orderproduct, error) {
	if e.loadedTypes[0] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// StockproductOrErr returns the Stockproduct value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) StockproductOrErr() ([]*Stock, error) {
	if e.loadedTypes[1] {
		return e.Stockproduct, nil
	}
	return nil, &NotLoadedError{edge: "stockproduct"}
}

// ForproductOrErr returns the Forproduct value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ForproductOrErr() ([]*Promotion, error) {
	if e.loadedTypes[2] {
		return e.Forproduct, nil
	}
	return nil, &NotLoadedError{edge: "forproduct"}
}

// FormproductonlineOrErr returns the Formproductonline value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) FormproductonlineOrErr() ([]*Orderonline, error) {
	if e.loadedTypes[3] {
		return e.Formproductonline, nil
	}
	return nil, &NotLoadedError{edge: "formproductonline"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // NameProduct
		&sql.NullString{}, // BarcodeProduct
		&sql.NullString{}, // MFG
		&sql.NullString{}, // EXP
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(values ...interface{}) error {
	if m, n := len(values), len(product.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field NameProduct", values[0])
	} else if value.Valid {
		pr.NameProduct = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field BarcodeProduct", values[1])
	} else if value.Valid {
		pr.BarcodeProduct = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field MFG", values[2])
	} else if value.Valid {
		pr.MFG = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field EXP", values[3])
	} else if value.Valid {
		pr.EXP = value.String
	}
	return nil
}

// QueryProducts queries the products edge of the Product.
func (pr *Product) QueryProducts() *OrderproductQuery {
	return (&ProductClient{config: pr.config}).QueryProducts(pr)
}

// QueryStockproduct queries the stockproduct edge of the Product.
func (pr *Product) QueryStockproduct() *StockQuery {
	return (&ProductClient{config: pr.config}).QueryStockproduct(pr)
}

// QueryForproduct queries the forproduct edge of the Product.
func (pr *Product) QueryForproduct() *PromotionQuery {
	return (&ProductClient{config: pr.config}).QueryForproduct(pr)
}

// QueryFormproductonline queries the formproductonline edge of the Product.
func (pr *Product) QueryFormproductonline() *OrderonlineQuery {
	return (&ProductClient{config: pr.config}).QueryFormproductonline(pr)
}

// Update returns a builder for updating this Product.
// Note that, you need to call Product.Unwrap() before calling this method, if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return (&ProductClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", NameProduct=")
	builder.WriteString(pr.NameProduct)
	builder.WriteString(", BarcodeProduct=")
	builder.WriteString(pr.BarcodeProduct)
	builder.WriteString(", MFG=")
	builder.WriteString(pr.MFG)
	builder.WriteString(", EXP=")
	builder.WriteString(pr.EXP)
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

func (pr Products) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
