// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team13/app/ent/employee"
	"github.com/team13/app/ent/product"
	"github.com/team13/app/ent/stock"
	"github.com/team13/app/ent/typeproduct"
	"github.com/team13/app/ent/zoneproduct"
)

// Stock is the model entity for the Stock schema.
type Stock struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Priceproduct holds the value of the "Priceproduct" field.
	Priceproduct string `json:"Priceproduct,omitempty"`
	// Amount holds the value of the "Amount" field.
	Amount string `json:"Amount,omitempty"`
	// Time holds the value of the "Time" field.
	Time time.Time `json:"Time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StockQuery when eager-loading is set.
	Edges                  StockEdges `json:"edges"`
	employee_employeestock *int
	product_stockproduct   *int
	typeproduct_typestock  *int
	zoneproduct_zonestock  *int
}

// StockEdges holds the relations/edges for other nodes in the graph.
type StockEdges struct {
	// Product holds the value of the product edge.
	Product *Product
	// Zoneproduct holds the value of the zoneproduct edge.
	Zoneproduct *Zoneproduct
	// Employee holds the value of the employee edge.
	Employee *Employee
	// Typeproduct holds the value of the Typeproduct edge.
	Typeproduct *Typeproduct
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StockEdges) ProductOrErr() (*Product, error) {
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

// ZoneproductOrErr returns the Zoneproduct value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StockEdges) ZoneproductOrErr() (*Zoneproduct, error) {
	if e.loadedTypes[1] {
		if e.Zoneproduct == nil {
			// The edge zoneproduct was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: zoneproduct.Label}
		}
		return e.Zoneproduct, nil
	}
	return nil, &NotLoadedError{edge: "zoneproduct"}
}

// EmployeeOrErr returns the Employee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StockEdges) EmployeeOrErr() (*Employee, error) {
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

// TypeproductOrErr returns the Typeproduct value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StockEdges) TypeproductOrErr() (*Typeproduct, error) {
	if e.loadedTypes[3] {
		if e.Typeproduct == nil {
			// The edge Typeproduct was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: typeproduct.Label}
		}
		return e.Typeproduct, nil
	}
	return nil, &NotLoadedError{edge: "Typeproduct"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Stock) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Priceproduct
		&sql.NullString{}, // Amount
		&sql.NullTime{},   // Time
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Stock) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // employee_employeestock
		&sql.NullInt64{}, // product_stockproduct
		&sql.NullInt64{}, // typeproduct_typestock
		&sql.NullInt64{}, // zoneproduct_zonestock
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Stock fields.
func (s *Stock) assignValues(values ...interface{}) error {
	if m, n := len(values), len(stock.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Priceproduct", values[0])
	} else if value.Valid {
		s.Priceproduct = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Amount", values[1])
	} else if value.Valid {
		s.Amount = value.String
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field Time", values[2])
	} else if value.Valid {
		s.Time = value.Time
	}
	values = values[3:]
	if len(values) == len(stock.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field employee_employeestock", value)
		} else if value.Valid {
			s.employee_employeestock = new(int)
			*s.employee_employeestock = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field product_stockproduct", value)
		} else if value.Valid {
			s.product_stockproduct = new(int)
			*s.product_stockproduct = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field typeproduct_typestock", value)
		} else if value.Valid {
			s.typeproduct_typestock = new(int)
			*s.typeproduct_typestock = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field zoneproduct_zonestock", value)
		} else if value.Valid {
			s.zoneproduct_zonestock = new(int)
			*s.zoneproduct_zonestock = int(value.Int64)
		}
	}
	return nil
}

// QueryProduct queries the product edge of the Stock.
func (s *Stock) QueryProduct() *ProductQuery {
	return (&StockClient{config: s.config}).QueryProduct(s)
}

// QueryZoneproduct queries the zoneproduct edge of the Stock.
func (s *Stock) QueryZoneproduct() *ZoneproductQuery {
	return (&StockClient{config: s.config}).QueryZoneproduct(s)
}

// QueryEmployee queries the employee edge of the Stock.
func (s *Stock) QueryEmployee() *EmployeeQuery {
	return (&StockClient{config: s.config}).QueryEmployee(s)
}

// QueryTypeproduct queries the Typeproduct edge of the Stock.
func (s *Stock) QueryTypeproduct() *TypeproductQuery {
	return (&StockClient{config: s.config}).QueryTypeproduct(s)
}

// Update returns a builder for updating this Stock.
// Note that, you need to call Stock.Unwrap() before calling this method, if this Stock
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Stock) Update() *StockUpdateOne {
	return (&StockClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Stock) Unwrap() *Stock {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Stock is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Stock) String() string {
	var builder strings.Builder
	builder.WriteString("Stock(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", Priceproduct=")
	builder.WriteString(s.Priceproduct)
	builder.WriteString(", Amount=")
	builder.WriteString(s.Amount)
	builder.WriteString(", Time=")
	builder.WriteString(s.Time.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Stocks is a parsable slice of Stock.
type Stocks []*Stock

func (s Stocks) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
