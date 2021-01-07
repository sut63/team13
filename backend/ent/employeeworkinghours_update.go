// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/day"
	"github.com/tanapon395/playlist-video/ent/employee"
	"github.com/tanapon395/playlist-video/ent/employeeworkinghours"
	"github.com/tanapon395/playlist-video/ent/predicate"
	"github.com/tanapon395/playlist-video/ent/role"
	"github.com/tanapon395/playlist-video/ent/shift"
)

// EmployeeworkinghoursUpdate is the builder for updating Employeeworkinghours entities.
type EmployeeworkinghoursUpdate struct {
	config
	hooks      []Hook
	mutation   *EmployeeworkinghoursMutation
	predicates []predicate.Employeeworkinghours
}

// Where adds a new predicate for the builder.
func (eu *EmployeeworkinghoursUpdate) Where(ps ...predicate.Employeeworkinghours) *EmployeeworkinghoursUpdate {
	eu.predicates = append(eu.predicates, ps...)
	return eu
}

// SetName sets the name field.
func (eu *EmployeeworkinghoursUpdate) SetName(s string) *EmployeeworkinghoursUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetWorkinghourID sets the workinghour edge to Employee by id.
func (eu *EmployeeworkinghoursUpdate) SetWorkinghourID(id int) *EmployeeworkinghoursUpdate {
	eu.mutation.SetWorkinghourID(id)
	return eu
}

// SetNillableWorkinghourID sets the workinghour edge to Employee by id if the given value is not nil.
func (eu *EmployeeworkinghoursUpdate) SetNillableWorkinghourID(id *int) *EmployeeworkinghoursUpdate {
	if id != nil {
		eu = eu.SetWorkinghourID(*id)
	}
	return eu
}

// SetWorkinghour sets the workinghour edge to Employee.
func (eu *EmployeeworkinghoursUpdate) SetWorkinghour(e *Employee) *EmployeeworkinghoursUpdate {
	return eu.SetWorkinghourID(e.ID)
}

// SetDayID sets the day edge to Day by id.
func (eu *EmployeeworkinghoursUpdate) SetDayID(id int) *EmployeeworkinghoursUpdate {
	eu.mutation.SetDayID(id)
	return eu
}

// SetNillableDayID sets the day edge to Day by id if the given value is not nil.
func (eu *EmployeeworkinghoursUpdate) SetNillableDayID(id *int) *EmployeeworkinghoursUpdate {
	if id != nil {
		eu = eu.SetDayID(*id)
	}
	return eu
}

// SetDay sets the day edge to Day.
func (eu *EmployeeworkinghoursUpdate) SetDay(d *Day) *EmployeeworkinghoursUpdate {
	return eu.SetDayID(d.ID)
}

// SetShiftID sets the shift edge to Shift by id.
func (eu *EmployeeworkinghoursUpdate) SetShiftID(id int) *EmployeeworkinghoursUpdate {
	eu.mutation.SetShiftID(id)
	return eu
}

// SetNillableShiftID sets the shift edge to Shift by id if the given value is not nil.
func (eu *EmployeeworkinghoursUpdate) SetNillableShiftID(id *int) *EmployeeworkinghoursUpdate {
	if id != nil {
		eu = eu.SetShiftID(*id)
	}
	return eu
}

// SetShift sets the shift edge to Shift.
func (eu *EmployeeworkinghoursUpdate) SetShift(s *Shift) *EmployeeworkinghoursUpdate {
	return eu.SetShiftID(s.ID)
}

// SetRoleID sets the role edge to Role by id.
func (eu *EmployeeworkinghoursUpdate) SetRoleID(id int) *EmployeeworkinghoursUpdate {
	eu.mutation.SetRoleID(id)
	return eu
}

// SetNillableRoleID sets the role edge to Role by id if the given value is not nil.
func (eu *EmployeeworkinghoursUpdate) SetNillableRoleID(id *int) *EmployeeworkinghoursUpdate {
	if id != nil {
		eu = eu.SetRoleID(*id)
	}
	return eu
}

// SetRole sets the role edge to Role.
func (eu *EmployeeworkinghoursUpdate) SetRole(r *Role) *EmployeeworkinghoursUpdate {
	return eu.SetRoleID(r.ID)
}

// Mutation returns the EmployeeworkinghoursMutation object of the builder.
func (eu *EmployeeworkinghoursUpdate) Mutation() *EmployeeworkinghoursMutation {
	return eu.mutation
}

// ClearWorkinghour clears the workinghour edge to Employee.
func (eu *EmployeeworkinghoursUpdate) ClearWorkinghour() *EmployeeworkinghoursUpdate {
	eu.mutation.ClearWorkinghour()
	return eu
}

// ClearDay clears the day edge to Day.
func (eu *EmployeeworkinghoursUpdate) ClearDay() *EmployeeworkinghoursUpdate {
	eu.mutation.ClearDay()
	return eu
}

// ClearShift clears the shift edge to Shift.
func (eu *EmployeeworkinghoursUpdate) ClearShift() *EmployeeworkinghoursUpdate {
	eu.mutation.ClearShift()
	return eu
}

// ClearRole clears the role edge to Role.
func (eu *EmployeeworkinghoursUpdate) ClearRole() *EmployeeworkinghoursUpdate {
	eu.mutation.ClearRole()
	return eu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (eu *EmployeeworkinghoursUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeworkinghoursMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmployeeworkinghoursUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmployeeworkinghoursUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmployeeworkinghoursUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EmployeeworkinghoursUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   employeeworkinghours.Table,
			Columns: employeeworkinghours.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employeeworkinghours.FieldID,
			},
		},
	}
	if ps := eu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employeeworkinghours.FieldName,
		})
	}
	if eu.mutation.WorkinghourCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.WorkinghourIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.DayCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.DayIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.ShiftCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ShiftIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.RoleCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RoleIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employeeworkinghours.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// EmployeeworkinghoursUpdateOne is the builder for updating a single Employeeworkinghours entity.
type EmployeeworkinghoursUpdateOne struct {
	config
	hooks    []Hook
	mutation *EmployeeworkinghoursMutation
}

// SetName sets the name field.
func (euo *EmployeeworkinghoursUpdateOne) SetName(s string) *EmployeeworkinghoursUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetWorkinghourID sets the workinghour edge to Employee by id.
func (euo *EmployeeworkinghoursUpdateOne) SetWorkinghourID(id int) *EmployeeworkinghoursUpdateOne {
	euo.mutation.SetWorkinghourID(id)
	return euo
}

// SetNillableWorkinghourID sets the workinghour edge to Employee by id if the given value is not nil.
func (euo *EmployeeworkinghoursUpdateOne) SetNillableWorkinghourID(id *int) *EmployeeworkinghoursUpdateOne {
	if id != nil {
		euo = euo.SetWorkinghourID(*id)
	}
	return euo
}

// SetWorkinghour sets the workinghour edge to Employee.
func (euo *EmployeeworkinghoursUpdateOne) SetWorkinghour(e *Employee) *EmployeeworkinghoursUpdateOne {
	return euo.SetWorkinghourID(e.ID)
}

// SetDayID sets the day edge to Day by id.
func (euo *EmployeeworkinghoursUpdateOne) SetDayID(id int) *EmployeeworkinghoursUpdateOne {
	euo.mutation.SetDayID(id)
	return euo
}

// SetNillableDayID sets the day edge to Day by id if the given value is not nil.
func (euo *EmployeeworkinghoursUpdateOne) SetNillableDayID(id *int) *EmployeeworkinghoursUpdateOne {
	if id != nil {
		euo = euo.SetDayID(*id)
	}
	return euo
}

// SetDay sets the day edge to Day.
func (euo *EmployeeworkinghoursUpdateOne) SetDay(d *Day) *EmployeeworkinghoursUpdateOne {
	return euo.SetDayID(d.ID)
}

// SetShiftID sets the shift edge to Shift by id.
func (euo *EmployeeworkinghoursUpdateOne) SetShiftID(id int) *EmployeeworkinghoursUpdateOne {
	euo.mutation.SetShiftID(id)
	return euo
}

// SetNillableShiftID sets the shift edge to Shift by id if the given value is not nil.
func (euo *EmployeeworkinghoursUpdateOne) SetNillableShiftID(id *int) *EmployeeworkinghoursUpdateOne {
	if id != nil {
		euo = euo.SetShiftID(*id)
	}
	return euo
}

// SetShift sets the shift edge to Shift.
func (euo *EmployeeworkinghoursUpdateOne) SetShift(s *Shift) *EmployeeworkinghoursUpdateOne {
	return euo.SetShiftID(s.ID)
}

// SetRoleID sets the role edge to Role by id.
func (euo *EmployeeworkinghoursUpdateOne) SetRoleID(id int) *EmployeeworkinghoursUpdateOne {
	euo.mutation.SetRoleID(id)
	return euo
}

// SetNillableRoleID sets the role edge to Role by id if the given value is not nil.
func (euo *EmployeeworkinghoursUpdateOne) SetNillableRoleID(id *int) *EmployeeworkinghoursUpdateOne {
	if id != nil {
		euo = euo.SetRoleID(*id)
	}
	return euo
}

// SetRole sets the role edge to Role.
func (euo *EmployeeworkinghoursUpdateOne) SetRole(r *Role) *EmployeeworkinghoursUpdateOne {
	return euo.SetRoleID(r.ID)
}

// Mutation returns the EmployeeworkinghoursMutation object of the builder.
func (euo *EmployeeworkinghoursUpdateOne) Mutation() *EmployeeworkinghoursMutation {
	return euo.mutation
}

// ClearWorkinghour clears the workinghour edge to Employee.
func (euo *EmployeeworkinghoursUpdateOne) ClearWorkinghour() *EmployeeworkinghoursUpdateOne {
	euo.mutation.ClearWorkinghour()
	return euo
}

// ClearDay clears the day edge to Day.
func (euo *EmployeeworkinghoursUpdateOne) ClearDay() *EmployeeworkinghoursUpdateOne {
	euo.mutation.ClearDay()
	return euo
}

// ClearShift clears the shift edge to Shift.
func (euo *EmployeeworkinghoursUpdateOne) ClearShift() *EmployeeworkinghoursUpdateOne {
	euo.mutation.ClearShift()
	return euo
}

// ClearRole clears the role edge to Role.
func (euo *EmployeeworkinghoursUpdateOne) ClearRole() *EmployeeworkinghoursUpdateOne {
	euo.mutation.ClearRole()
	return euo
}

// Save executes the query and returns the updated entity.
func (euo *EmployeeworkinghoursUpdateOne) Save(ctx context.Context) (*Employeeworkinghours, error) {

	var (
		err  error
		node *Employeeworkinghours
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeworkinghoursMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmployeeworkinghoursUpdateOne) SaveX(ctx context.Context) *Employeeworkinghours {
	e, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return e
}

// Exec executes the query on the entity.
func (euo *EmployeeworkinghoursUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmployeeworkinghoursUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EmployeeworkinghoursUpdateOne) sqlSave(ctx context.Context) (e *Employeeworkinghours, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   employeeworkinghours.Table,
			Columns: employeeworkinghours.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employeeworkinghours.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Employeeworkinghours.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := euo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employeeworkinghours.FieldName,
		})
	}
	if euo.mutation.WorkinghourCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.WorkinghourIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.DayCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.DayIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.ShiftCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ShiftIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.RoleCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RoleIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	e = &Employeeworkinghours{config: euo.config}
	_spec.Assign = e.assignValues
	_spec.ScanValues = e.scanValues()
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employeeworkinghours.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return e, nil
}