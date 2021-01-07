// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/day"
	"github.com/tanapon395/playlist-video/ent/employee"
	"github.com/tanapon395/playlist-video/ent/employeeworkinghours"
	"github.com/tanapon395/playlist-video/ent/role"
	"github.com/tanapon395/playlist-video/ent/shift"
)

// EmployeeworkinghoursCreate is the builder for creating a Employeeworkinghours entity.
type EmployeeworkinghoursCreate struct {
	config
	mutation *EmployeeworkinghoursMutation
	hooks    []Hook
}

// SetName sets the name field.
func (ec *EmployeeworkinghoursCreate) SetName(s string) *EmployeeworkinghoursCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetWorkinghourID sets the workinghour edge to Employee by id.
func (ec *EmployeeworkinghoursCreate) SetWorkinghourID(id int) *EmployeeworkinghoursCreate {
	ec.mutation.SetWorkinghourID(id)
	return ec
}

// SetNillableWorkinghourID sets the workinghour edge to Employee by id if the given value is not nil.
func (ec *EmployeeworkinghoursCreate) SetNillableWorkinghourID(id *int) *EmployeeworkinghoursCreate {
	if id != nil {
		ec = ec.SetWorkinghourID(*id)
	}
	return ec
}

// SetWorkinghour sets the workinghour edge to Employee.
func (ec *EmployeeworkinghoursCreate) SetWorkinghour(e *Employee) *EmployeeworkinghoursCreate {
	return ec.SetWorkinghourID(e.ID)
}

// SetDayID sets the day edge to Day by id.
func (ec *EmployeeworkinghoursCreate) SetDayID(id int) *EmployeeworkinghoursCreate {
	ec.mutation.SetDayID(id)
	return ec
}

// SetNillableDayID sets the day edge to Day by id if the given value is not nil.
func (ec *EmployeeworkinghoursCreate) SetNillableDayID(id *int) *EmployeeworkinghoursCreate {
	if id != nil {
		ec = ec.SetDayID(*id)
	}
	return ec
}

// SetDay sets the day edge to Day.
func (ec *EmployeeworkinghoursCreate) SetDay(d *Day) *EmployeeworkinghoursCreate {
	return ec.SetDayID(d.ID)
}

// SetShiftID sets the shift edge to Shift by id.
func (ec *EmployeeworkinghoursCreate) SetShiftID(id int) *EmployeeworkinghoursCreate {
	ec.mutation.SetShiftID(id)
	return ec
}

// SetNillableShiftID sets the shift edge to Shift by id if the given value is not nil.
func (ec *EmployeeworkinghoursCreate) SetNillableShiftID(id *int) *EmployeeworkinghoursCreate {
	if id != nil {
		ec = ec.SetShiftID(*id)
	}
	return ec
}

// SetShift sets the shift edge to Shift.
func (ec *EmployeeworkinghoursCreate) SetShift(s *Shift) *EmployeeworkinghoursCreate {
	return ec.SetShiftID(s.ID)
}

// SetRoleID sets the role edge to Role by id.
func (ec *EmployeeworkinghoursCreate) SetRoleID(id int) *EmployeeworkinghoursCreate {
	ec.mutation.SetRoleID(id)
	return ec
}

// SetNillableRoleID sets the role edge to Role by id if the given value is not nil.
func (ec *EmployeeworkinghoursCreate) SetNillableRoleID(id *int) *EmployeeworkinghoursCreate {
	if id != nil {
		ec = ec.SetRoleID(*id)
	}
	return ec
}

// SetRole sets the role edge to Role.
func (ec *EmployeeworkinghoursCreate) SetRole(r *Role) *EmployeeworkinghoursCreate {
	return ec.SetRoleID(r.ID)
}

// Mutation returns the EmployeeworkinghoursMutation object of the builder.
func (ec *EmployeeworkinghoursCreate) Mutation() *EmployeeworkinghoursMutation {
	return ec.mutation
}

// Save creates the Employeeworkinghours in the database.
func (ec *EmployeeworkinghoursCreate) Save(ctx context.Context) (*Employeeworkinghours, error) {
	if _, ok := ec.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	var (
		err  error
		node *Employeeworkinghours
	)
	if len(ec.hooks) == 0 {
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeworkinghoursMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ec.mutation = mutation
			node, err = ec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EmployeeworkinghoursCreate) SaveX(ctx context.Context) *Employeeworkinghours {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ec *EmployeeworkinghoursCreate) sqlSave(ctx context.Context) (*Employeeworkinghours, error) {
	e, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	e.ID = int(id)
	return e, nil
}

func (ec *EmployeeworkinghoursCreate) createSpec() (*Employeeworkinghours, *sqlgraph.CreateSpec) {
	var (
		e     = &Employeeworkinghours{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: employeeworkinghours.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employeeworkinghours.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employeeworkinghours.FieldName,
		})
		e.Name = value
	}
	if nodes := ec.mutation.WorkinghourIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employeeworkinghours.WorkinghourTable,
			Columns: []string{employeeworkinghours.WorkinghourColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.DayIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employeeworkinghours.DayTable,
			Columns: []string{employeeworkinghours.DayColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: day.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ShiftIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employeeworkinghours.ShiftTable,
			Columns: []string{employeeworkinghours.ShiftColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shift.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employeeworkinghours.RoleTable,
			Columns: []string{employeeworkinghours.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return e, _spec
}
