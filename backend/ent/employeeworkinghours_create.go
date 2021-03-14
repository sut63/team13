// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team13/app/ent/day"
	"github.com/team13/app/ent/employee"
	"github.com/team13/app/ent/employeeworkinghours"
	"github.com/team13/app/ent/endwork"
	"github.com/team13/app/ent/role"
	"github.com/team13/app/ent/startwork"
)

// EmployeeWorkingHoursCreate is the builder for creating a EmployeeWorkingHours entity.
type EmployeeWorkingHoursCreate struct {
	config
	mutation *EmployeeWorkingHoursMutation
	hooks    []Hook
}

// SetCodeWork sets the CodeWork field.
func (ewhc *EmployeeWorkingHoursCreate) SetCodeWork(s string) *EmployeeWorkingHoursCreate {
	ewhc.mutation.SetCodeWork(s)
	return ewhc
}

// SetIDNumber sets the IDNumber field.
func (ewhc *EmployeeWorkingHoursCreate) SetIDNumber(s string) *EmployeeWorkingHoursCreate {
	ewhc.mutation.SetIDNumber(s)
	return ewhc
}

// SetWages sets the Wages field.
func (ewhc *EmployeeWorkingHoursCreate) SetWages(f float64) *EmployeeWorkingHoursCreate {
	ewhc.mutation.SetWages(f)
	return ewhc
}

// SetEmployeeID sets the employee edge to Employee by id.
func (ewhc *EmployeeWorkingHoursCreate) SetEmployeeID(id int) *EmployeeWorkingHoursCreate {
	ewhc.mutation.SetEmployeeID(id)
	return ewhc
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (ewhc *EmployeeWorkingHoursCreate) SetNillableEmployeeID(id *int) *EmployeeWorkingHoursCreate {
	if id != nil {
		ewhc = ewhc.SetEmployeeID(*id)
	}
	return ewhc
}

// SetEmployee sets the employee edge to Employee.
func (ewhc *EmployeeWorkingHoursCreate) SetEmployee(e *Employee) *EmployeeWorkingHoursCreate {
	return ewhc.SetEmployeeID(e.ID)
}

// SetDayID sets the day edge to Day by id.
func (ewhc *EmployeeWorkingHoursCreate) SetDayID(id int) *EmployeeWorkingHoursCreate {
	ewhc.mutation.SetDayID(id)
	return ewhc
}

// SetNillableDayID sets the day edge to Day by id if the given value is not nil.
func (ewhc *EmployeeWorkingHoursCreate) SetNillableDayID(id *int) *EmployeeWorkingHoursCreate {
	if id != nil {
		ewhc = ewhc.SetDayID(*id)
	}
	return ewhc
}

// SetDay sets the day edge to Day.
func (ewhc *EmployeeWorkingHoursCreate) SetDay(d *Day) *EmployeeWorkingHoursCreate {
	return ewhc.SetDayID(d.ID)
}

// SetStartworkID sets the startwork edge to StartWork by id.
func (ewhc *EmployeeWorkingHoursCreate) SetStartworkID(id int) *EmployeeWorkingHoursCreate {
	ewhc.mutation.SetStartworkID(id)
	return ewhc
}

// SetNillableStartworkID sets the startwork edge to StartWork by id if the given value is not nil.
func (ewhc *EmployeeWorkingHoursCreate) SetNillableStartworkID(id *int) *EmployeeWorkingHoursCreate {
	if id != nil {
		ewhc = ewhc.SetStartworkID(*id)
	}
	return ewhc
}

// SetStartwork sets the startwork edge to StartWork.
func (ewhc *EmployeeWorkingHoursCreate) SetStartwork(s *StartWork) *EmployeeWorkingHoursCreate {
	return ewhc.SetStartworkID(s.ID)
}

// SetEndworkID sets the endwork edge to EndWork by id.
func (ewhc *EmployeeWorkingHoursCreate) SetEndworkID(id int) *EmployeeWorkingHoursCreate {
	ewhc.mutation.SetEndworkID(id)
	return ewhc
}

// SetNillableEndworkID sets the endwork edge to EndWork by id if the given value is not nil.
func (ewhc *EmployeeWorkingHoursCreate) SetNillableEndworkID(id *int) *EmployeeWorkingHoursCreate {
	if id != nil {
		ewhc = ewhc.SetEndworkID(*id)
	}
	return ewhc
}

// SetEndwork sets the endwork edge to EndWork.
func (ewhc *EmployeeWorkingHoursCreate) SetEndwork(e *EndWork) *EmployeeWorkingHoursCreate {
	return ewhc.SetEndworkID(e.ID)
}

// SetRoleID sets the role edge to Role by id.
func (ewhc *EmployeeWorkingHoursCreate) SetRoleID(id int) *EmployeeWorkingHoursCreate {
	ewhc.mutation.SetRoleID(id)
	return ewhc
}

// SetNillableRoleID sets the role edge to Role by id if the given value is not nil.
func (ewhc *EmployeeWorkingHoursCreate) SetNillableRoleID(id *int) *EmployeeWorkingHoursCreate {
	if id != nil {
		ewhc = ewhc.SetRoleID(*id)
	}
	return ewhc
}

// SetRole sets the role edge to Role.
func (ewhc *EmployeeWorkingHoursCreate) SetRole(r *Role) *EmployeeWorkingHoursCreate {
	return ewhc.SetRoleID(r.ID)
}

// Mutation returns the EmployeeWorkingHoursMutation object of the builder.
func (ewhc *EmployeeWorkingHoursCreate) Mutation() *EmployeeWorkingHoursMutation {
	return ewhc.mutation
}

// Save creates the EmployeeWorkingHours in the database.
func (ewhc *EmployeeWorkingHoursCreate) Save(ctx context.Context) (*EmployeeWorkingHours, error) {
	if _, ok := ewhc.mutation.CodeWork(); !ok {
		return nil, &ValidationError{Name: "CodeWork", err: errors.New("ent: missing required field \"CodeWork\"")}
	}
	if v, ok := ewhc.mutation.CodeWork(); ok {
		if err := employeeworkinghours.CodeWorkValidator(v); err != nil {
			return nil, &ValidationError{Name: "CodeWork", err: fmt.Errorf("ent: validator failed for field \"CodeWork\": %w", err)}
		}
	}
	if _, ok := ewhc.mutation.IDNumber(); !ok {
		return nil, &ValidationError{Name: "IDNumber", err: errors.New("ent: missing required field \"IDNumber\"")}
	}
	if v, ok := ewhc.mutation.IDNumber(); ok {
		if err := employeeworkinghours.IDNumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "IDNumber", err: fmt.Errorf("ent: validator failed for field \"IDNumber\": %w", err)}
		}
	}
	if _, ok := ewhc.mutation.Wages(); !ok {
		return nil, &ValidationError{Name: "Wages", err: errors.New("ent: missing required field \"Wages\"")}
	}
	if v, ok := ewhc.mutation.Wages(); ok {
		if err := employeeworkinghours.WagesValidator(v); err != nil {
			return nil, &ValidationError{Name: "Wages", err: fmt.Errorf("ent: validator failed for field \"Wages\": %w", err)}
		}
	}
	var (
		err  error
		node *EmployeeWorkingHours
	)
	if len(ewhc.hooks) == 0 {
		node, err = ewhc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeWorkingHoursMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ewhc.mutation = mutation
			node, err = ewhc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ewhc.hooks) - 1; i >= 0; i-- {
			mut = ewhc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ewhc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ewhc *EmployeeWorkingHoursCreate) SaveX(ctx context.Context) *EmployeeWorkingHours {
	v, err := ewhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ewhc *EmployeeWorkingHoursCreate) sqlSave(ctx context.Context) (*EmployeeWorkingHours, error) {
	ewh, _spec := ewhc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ewhc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ewh.ID = int(id)
	return ewh, nil
}

func (ewhc *EmployeeWorkingHoursCreate) createSpec() (*EmployeeWorkingHours, *sqlgraph.CreateSpec) {
	var (
		ewh   = &EmployeeWorkingHours{config: ewhc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: employeeworkinghours.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employeeworkinghours.FieldID,
			},
		}
	)
	if value, ok := ewhc.mutation.CodeWork(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employeeworkinghours.FieldCodeWork,
		})
		ewh.CodeWork = value
	}
	if value, ok := ewhc.mutation.IDNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employeeworkinghours.FieldIDNumber,
		})
		ewh.IDNumber = value
	}
	if value, ok := ewhc.mutation.Wages(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: employeeworkinghours.FieldWages,
		})
		ewh.Wages = value
	}
	if nodes := ewhc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employeeworkinghours.EmployeeTable,
			Columns: []string{employeeworkinghours.EmployeeColumn},
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
	if nodes := ewhc.mutation.DayIDs(); len(nodes) > 0 {
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
	if nodes := ewhc.mutation.StartworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employeeworkinghours.StartworkTable,
			Columns: []string{employeeworkinghours.StartworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: startwork.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ewhc.mutation.EndworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employeeworkinghours.EndworkTable,
			Columns: []string{employeeworkinghours.EndworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: endwork.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ewhc.mutation.RoleIDs(); len(nodes) > 0 {
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
	return ewh, _spec
}
