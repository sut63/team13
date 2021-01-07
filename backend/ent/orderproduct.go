// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/company"
	"github.com/tanapon395/playlist-video/ent/manager"
	"github.com/tanapon395/playlist-video/ent/orderproduct"
	"github.com/tanapon395/playlist-video/ent/product"
	"github.com/tanapon395/playlist-video/ent/typeproduct"
)

// Orderproduct is the model entity for the Orderproduct schema.
type Orderproduct struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Addedtime holds the value of the "addedtime" field.
	Addedtime time.Time `json:"addedtime,omitempty"`
	// Stock holds the value of the "stock" field.
	Stock int `json:"stock,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderproductQuery when eager-loading is set.
	Edges                    OrderproductEdges `json:"edges"`
	company_companys         *int
	manager_managers         *int
	product_products         *int
	typeproduct_typeproducts *int
}

// OrderproductEdges holds the relations/edges for other nodes in the graph.
type OrderproductEdges struct {
	// Product holds the value of the product edge.
	Product *Product
	// Company holds the value of the company edge.
	Company *Company
	// Typeproduct holds the value of the Typeproduct edge.
	Typeproduct *Typeproduct
	// Managers holds the value of the managers edge.
	Managers *Manager
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderproductEdges) ProductOrErr() (*Product, error) {
	if e.loadedTypes[0] {
		if e.Product == nil {
			// The edge product was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderproductEdges) CompanyOrErr() (*Company, error) {
	if e.loadedTypes[1] {
		if e.Company == nil {
			// The edge company was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: company.Label}
		}
		return e.Company, nil
	}
	return nil, &NotLoadedError{edge: "company"}
}

// TypeproductOrErr returns the Typeproduct value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderproductEdges) TypeproductOrErr() (*Typeproduct, error) {
	if e.loadedTypes[2] {
		if e.Typeproduct == nil {
			// The edge Typeproduct was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: typeproduct.Label}
		}
		return e.Typeproduct, nil
	}
	return nil, &NotLoadedError{edge: "Typeproduct"}
}

// ManagersOrErr returns the Managers value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderproductEdges) ManagersOrErr() (*Manager, error) {
	if e.loadedTypes[3] {
		if e.Managers == nil {
			// The edge managers was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: manager.Label}
		}
		return e.Managers, nil
	}
	return nil, &NotLoadedError{edge: "managers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Orderproduct) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // addedtime
		&sql.NullInt64{}, // stock
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Orderproduct) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // company_companys
		&sql.NullInt64{}, // manager_managers
		&sql.NullInt64{}, // product_products
		&sql.NullInt64{}, // typeproduct_typeproducts
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Orderproduct fields.
func (o *Orderproduct) assignValues(values ...interface{}) error {
	if m, n := len(values), len(orderproduct.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	o.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field addedtime", values[0])
	} else if value.Valid {
		o.Addedtime = value.Time
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field stock", values[1])
	} else if value.Valid {
		o.Stock = int(value.Int64)
	}
	values = values[2:]
	if len(values) == len(orderproduct.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field company_companys", value)
		} else if value.Valid {
			o.company_companys = new(int)
			*o.company_companys = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field manager_managers", value)
		} else if value.Valid {
			o.manager_managers = new(int)
			*o.manager_managers = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field product_products", value)
		} else if value.Valid {
			o.product_products = new(int)
			*o.product_products = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field typeproduct_typeproducts", value)
		} else if value.Valid {
			o.typeproduct_typeproducts = new(int)
			*o.typeproduct_typeproducts = int(value.Int64)
		}
	}
	return nil
}

// QueryProduct queries the product edge of the Orderproduct.
func (o *Orderproduct) QueryProduct() *ProductQuery {
	return (&OrderproductClient{config: o.config}).QueryProduct(o)
}

// QueryCompany queries the company edge of the Orderproduct.
func (o *Orderproduct) QueryCompany() *CompanyQuery {
	return (&OrderproductClient{config: o.config}).QueryCompany(o)
}

// QueryTypeproduct queries the Typeproduct edge of the Orderproduct.
func (o *Orderproduct) QueryTypeproduct() *TypeproductQuery {
	return (&OrderproductClient{config: o.config}).QueryTypeproduct(o)
}

// QueryManagers queries the managers edge of the Orderproduct.
func (o *Orderproduct) QueryManagers() *ManagerQuery {
	return (&OrderproductClient{config: o.config}).QueryManagers(o)
}

// Update returns a builder for updating this Orderproduct.
// Note that, you need to call Orderproduct.Unwrap() before calling this method, if this Orderproduct
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Orderproduct) Update() *OrderproductUpdateOne {
	return (&OrderproductClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (o *Orderproduct) Unwrap() *Orderproduct {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Orderproduct is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Orderproduct) String() string {
	var builder strings.Builder
	builder.WriteString("Orderproduct(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", addedtime=")
	builder.WriteString(o.Addedtime.Format(time.ANSIC))
	builder.WriteString(", stock=")
	builder.WriteString(fmt.Sprintf("%v", o.Stock))
	builder.WriteByte(')')
	return builder.String()
}

// Orderproducts is a parsable slice of Orderproduct.
type Orderproducts []*Orderproduct

func (o Orderproducts) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}