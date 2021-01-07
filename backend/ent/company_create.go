// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/company"
	"github.com/tanapon395/playlist-video/ent/orderproduct"
)

// CompanyCreate is the builder for creating a Company entity.
type CompanyCreate struct {
	config
	mutation *CompanyMutation
	hooks    []Hook
}

// SetName sets the Name field.
func (cc *CompanyCreate) SetName(s string) *CompanyCreate {
	cc.mutation.SetName(s)
	return cc
}

// AddCompanyIDs adds the companys edge to Orderproduct by ids.
func (cc *CompanyCreate) AddCompanyIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddCompanyIDs(ids...)
	return cc
}

// AddCompanys adds the companys edges to Orderproduct.
func (cc *CompanyCreate) AddCompanys(o ...*Orderproduct) *CompanyCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cc.AddCompanyIDs(ids...)
}

// Mutation returns the CompanyMutation object of the builder.
func (cc *CompanyCreate) Mutation() *CompanyMutation {
	return cc.mutation
}

// Save creates the Company in the database.
func (cc *CompanyCreate) Save(ctx context.Context) (*Company, error) {
	if _, ok := cc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "Name", err: errors.New("ent: missing required field \"Name\"")}
	}
	var (
		err  error
		node *Company
	)
	if len(cc.hooks) == 0 {
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompanyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CompanyCreate) SaveX(ctx context.Context) *Company {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CompanyCreate) sqlSave(ctx context.Context) (*Company, error) {
	c, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}

func (cc *CompanyCreate) createSpec() (*Company, *sqlgraph.CreateSpec) {
	var (
		c     = &Company{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: company.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: company.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldName,
		})
		c.Name = value
	}
	if nodes := cc.mutation.CompanysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.CompanysTable,
			Columns: []string{company.CompanysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderproduct.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return c, _spec
}