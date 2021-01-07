// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team13/app/ent/day"
	"github.com/team13/app/ent/employee"
	"github.com/team13/app/ent/employeeworkinghours"
	"github.com/team13/app/ent/predicate"
	"github.com/team13/app/ent/role"
	"github.com/team13/app/ent/shift"
)

// EmployeeWorkingHoursUpdate is the builder for updating EmployeeWorkingHours entities.
type EmployeeWorkingHoursUpdate struct {
	config
	hooks      []Hook
	mutation   *EmployeeWorkingHoursMutation
	predicates []predicate.EmployeeWorkingHours
}

// Where adds a new predicate for the builder.
func (ewhu *EmployeeWorkingHoursUpdate) Where(ps ...predicate.EmployeeWorkingHours) *EmployeeWorkingHoursUpdate {
	ewhu.predicates = append(ewhu.predicates, ps...)
	return ewhu
}

// SetEmployeeWorkingHoursID sets the EmployeeWorkingHours edge to Employee by id.
func (ewhu *EmployeeWorkingHoursUpdate) SetEmployeeWorkingHoursID(id int) *EmployeeWorkingHoursUpdate {
	ewhu.mutation.SetEmployeeWorkingHoursID(id)
	return ewhu
}

// SetNillableEmployeeWorkingHoursID sets the EmployeeWorkingHours edge to Employee by id if the given value is not nil.
func (ewhu *EmployeeWorkingHoursUpdate) SetNillableEmployeeWorkingHoursID(id *int) *EmployeeWorkingHoursUpdate {
	if id != nil {
		ewhu = ewhu.SetEmployeeWorkingHoursID(*id)
	}
	return ewhu
}

// SetEmployeeWorkingHours sets the EmployeeWorkingHours edge to Employee.
func (ewhu *EmployeeWorkingHoursUpdate) SetEmployeeWorkingHours(e *Employee) *EmployeeWorkingHoursUpdate {
	return ewhu.SetEmployeeWorkingHoursID(e.ID)
}

// SetDayID sets the day edge to Day by id.
func (ewhu *EmployeeWorkingHoursUpdate) SetDayID(id int) *EmployeeWorkingHoursUpdate {
	ewhu.mutation.SetDayID(id)
	return ewhu
}

// SetNillableDayID sets the day edge to Day by id if the given value is not nil.
func (ewhu *EmployeeWorkingHoursUpdate) SetNillableDayID(id *int) *EmployeeWorkingHoursUpdate {
	if id != nil {
		ewhu = ewhu.SetDayID(*id)
	}
	return ewhu
}

// SetDay sets the day edge to Day.
func (ewhu *EmployeeWorkingHoursUpdate) SetDay(d *Day) *EmployeeWorkingHoursUpdate {
	return ewhu.SetDayID(d.ID)
}

// SetShiftID sets the shift edge to Shift by id.
func (ewhu *EmployeeWorkingHoursUpdate) SetShiftID(id int) *EmployeeWorkingHoursUpdate {
	ewhu.mutation.SetShiftID(id)
	return ewhu
}

// SetNillableShiftID sets the shift edge to Shift by id if the given value is not nil.
func (ewhu *EmployeeWorkingHoursUpdate) SetNillableShiftID(id *int) *EmployeeWorkingHoursUpdate {
	if id != nil {
		ewhu = ewhu.SetShiftID(*id)
	}
	return ewhu
}

// SetShift sets the shift edge to Shift.
func (ewhu *EmployeeWorkingHoursUpdate) SetShift(s *Shift) *EmployeeWorkingHoursUpdate {
	return ewhu.SetShiftID(s.ID)
}

// SetRoleID sets the role edge to Role by id.
func (ewhu *EmployeeWorkingHoursUpdate) SetRoleID(id int) *EmployeeWorkingHoursUpdate {
	ewhu.mutation.SetRoleID(id)
	return ewhu
}

// SetNillableRoleID sets the role edge to Role by id if the given value is not nil.
func (ewhu *EmployeeWorkingHoursUpdate) SetNillableRoleID(id *int) *EmployeeWorkingHoursUpdate {
	if id != nil {
		ewhu = ewhu.SetRoleID(*id)
	}
	return ewhu
}

// SetRole sets the role edge to Role.
func (ewhu *EmployeeWorkingHoursUpdate) SetRole(r *Role) *EmployeeWorkingHoursUpdate {
	return ewhu.SetRoleID(r.ID)
}

// Mutation returns the EmployeeWorkingHoursMutation object of the builder.
func (ewhu *EmployeeWorkingHoursUpdate) Mutation() *EmployeeWorkingHoursMutation {
	return ewhu.mutation
}

// ClearEmployeeWorkingHours clears the EmployeeWorkingHours edge to Employee.
func (ewhu *EmployeeWorkingHoursUpdate) ClearEmployeeWorkingHours() *EmployeeWorkingHoursUpdate {
	ewhu.mutation.ClearEmployeeWorkingHours()
	return ewhu
}

// ClearDay clears the day edge to Day.
func (ewhu *EmployeeWorkingHoursUpdate) ClearDay() *EmployeeWorkingHoursUpdate {
	ewhu.mutation.ClearDay()
	return ewhu
}

// ClearShift clears the shift edge to Shift.
func (ewhu *EmployeeWorkingHoursUpdate) ClearShift() *EmployeeWorkingHoursUpdate {
	ewhu.mutation.ClearShift()
	return ewhu
}

// ClearRole clears the role edge to Role.
func (ewhu *EmployeeWorkingHoursUpdate) ClearRole() *EmployeeWorkingHoursUpdate {
	ewhu.mutation.ClearRole()
	return ewhu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ewhu *EmployeeWorkingHoursUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(ewhu.hooks) == 0 {
		affected, err = ewhu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeWorkingHoursMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ewhu.mutation = mutation
			affected, err = ewhu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ewhu.hooks) - 1; i >= 0; i-- {
			mut = ewhu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ewhu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ewhu *EmployeeWorkingHoursUpdate) SaveX(ctx context.Context) int {
	affected, err := ewhu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ewhu *EmployeeWorkingHoursUpdate) Exec(ctx context.Context) error {
	_, err := ewhu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ewhu *EmployeeWorkingHoursUpdate) ExecX(ctx context.Context) {
	if err := ewhu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ewhu *EmployeeWorkingHoursUpdate) sqlSave(ctx context.Context) (n int, err error) {
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
	if ps := ewhu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ewhu.mutation.EmployeeWorkingHoursCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employeeworkinghours.EmployeeWorkingHoursTable,
			Columns: []string{employeeworkinghours.EmployeeWorkingHoursColumn},
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
	if nodes := ewhu.mutation.EmployeeWorkingHoursIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employeeworkinghours.EmployeeWorkingHoursTable,
			Columns: []string{employeeworkinghours.EmployeeWorkingHoursColumn},
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
	if ewhu.mutation.DayCleared() {
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
	if nodes := ewhu.mutation.DayIDs(); len(nodes) > 0 {
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
	if ewhu.mutation.ShiftCleared() {
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
	if nodes := ewhu.mutation.ShiftIDs(); len(nodes) > 0 {
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
	if ewhu.mutation.RoleCleared() {
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
	if nodes := ewhu.mutation.RoleIDs(); len(nodes) > 0 {
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
	if n, err = sqlgraph.UpdateNodes(ctx, ewhu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employeeworkinghours.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// EmployeeWorkingHoursUpdateOne is the builder for updating a single EmployeeWorkingHours entity.
type EmployeeWorkingHoursUpdateOne struct {
	config
	hooks    []Hook
	mutation *EmployeeWorkingHoursMutation
}

// SetEmployeeWorkingHoursID sets the EmployeeWorkingHours edge to Employee by id.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SetEmployeeWorkingHoursID(id int) *EmployeeWorkingHoursUpdateOne {
	ewhuo.mutation.SetEmployeeWorkingHoursID(id)
	return ewhuo
}

// SetNillableEmployeeWorkingHoursID sets the EmployeeWorkingHours edge to Employee by id if the given value is not nil.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SetNillableEmployeeWorkingHoursID(id *int) *EmployeeWorkingHoursUpdateOne {
	if id != nil {
		ewhuo = ewhuo.SetEmployeeWorkingHoursID(*id)
	}
	return ewhuo
}

// SetEmployeeWorkingHours sets the EmployeeWorkingHours edge to Employee.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SetEmployeeWorkingHours(e *Employee) *EmployeeWorkingHoursUpdateOne {
	return ewhuo.SetEmployeeWorkingHoursID(e.ID)
}

// SetDayID sets the day edge to Day by id.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SetDayID(id int) *EmployeeWorkingHoursUpdateOne {
	ewhuo.mutation.SetDayID(id)
	return ewhuo
}

// SetNillableDayID sets the day edge to Day by id if the given value is not nil.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SetNillableDayID(id *int) *EmployeeWorkingHoursUpdateOne {
	if id != nil {
		ewhuo = ewhuo.SetDayID(*id)
	}
	return ewhuo
}

// SetDay sets the day edge to Day.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SetDay(d *Day) *EmployeeWorkingHoursUpdateOne {
	return ewhuo.SetDayID(d.ID)
}

// SetShiftID sets the shift edge to Shift by id.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SetShiftID(id int) *EmployeeWorkingHoursUpdateOne {
	ewhuo.mutation.SetShiftID(id)
	return ewhuo
}

// SetNillableShiftID sets the shift edge to Shift by id if the given value is not nil.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SetNillableShiftID(id *int) *EmployeeWorkingHoursUpdateOne {
	if id != nil {
		ewhuo = ewhuo.SetShiftID(*id)
	}
	return ewhuo
}

// SetShift sets the shift edge to Shift.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SetShift(s *Shift) *EmployeeWorkingHoursUpdateOne {
	return ewhuo.SetShiftID(s.ID)
}

// SetRoleID sets the role edge to Role by id.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SetRoleID(id int) *EmployeeWorkingHoursUpdateOne {
	ewhuo.mutation.SetRoleID(id)
	return ewhuo
}

// SetNillableRoleID sets the role edge to Role by id if the given value is not nil.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SetNillableRoleID(id *int) *EmployeeWorkingHoursUpdateOne {
	if id != nil {
		ewhuo = ewhuo.SetRoleID(*id)
	}
	return ewhuo
}

// SetRole sets the role edge to Role.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SetRole(r *Role) *EmployeeWorkingHoursUpdateOne {
	return ewhuo.SetRoleID(r.ID)
}

// Mutation returns the EmployeeWorkingHoursMutation object of the builder.
func (ewhuo *EmployeeWorkingHoursUpdateOne) Mutation() *EmployeeWorkingHoursMutation {
	return ewhuo.mutation
}

// ClearEmployeeWorkingHours clears the EmployeeWorkingHours edge to Employee.
func (ewhuo *EmployeeWorkingHoursUpdateOne) ClearEmployeeWorkingHours() *EmployeeWorkingHoursUpdateOne {
	ewhuo.mutation.ClearEmployeeWorkingHours()
	return ewhuo
}

// ClearDay clears the day edge to Day.
func (ewhuo *EmployeeWorkingHoursUpdateOne) ClearDay() *EmployeeWorkingHoursUpdateOne {
	ewhuo.mutation.ClearDay()
	return ewhuo
}

// ClearShift clears the shift edge to Shift.
func (ewhuo *EmployeeWorkingHoursUpdateOne) ClearShift() *EmployeeWorkingHoursUpdateOne {
	ewhuo.mutation.ClearShift()
	return ewhuo
}

// ClearRole clears the role edge to Role.
func (ewhuo *EmployeeWorkingHoursUpdateOne) ClearRole() *EmployeeWorkingHoursUpdateOne {
	ewhuo.mutation.ClearRole()
	return ewhuo
}

// Save executes the query and returns the updated entity.
func (ewhuo *EmployeeWorkingHoursUpdateOne) Save(ctx context.Context) (*EmployeeWorkingHours, error) {

	var (
		err  error
		node *EmployeeWorkingHours
	)
	if len(ewhuo.hooks) == 0 {
		node, err = ewhuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeWorkingHoursMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ewhuo.mutation = mutation
			node, err = ewhuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ewhuo.hooks) - 1; i >= 0; i-- {
			mut = ewhuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ewhuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ewhuo *EmployeeWorkingHoursUpdateOne) SaveX(ctx context.Context) *EmployeeWorkingHours {
	ewh, err := ewhuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ewh
}

// Exec executes the query on the entity.
func (ewhuo *EmployeeWorkingHoursUpdateOne) Exec(ctx context.Context) error {
	_, err := ewhuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ewhuo *EmployeeWorkingHoursUpdateOne) ExecX(ctx context.Context) {
	if err := ewhuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ewhuo *EmployeeWorkingHoursUpdateOne) sqlSave(ctx context.Context) (ewh *EmployeeWorkingHours, err error) {
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
	id, ok := ewhuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing EmployeeWorkingHours.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ewhuo.mutation.EmployeeWorkingHoursCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employeeworkinghours.EmployeeWorkingHoursTable,
			Columns: []string{employeeworkinghours.EmployeeWorkingHoursColumn},
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
	if nodes := ewhuo.mutation.EmployeeWorkingHoursIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employeeworkinghours.EmployeeWorkingHoursTable,
			Columns: []string{employeeworkinghours.EmployeeWorkingHoursColumn},
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
	if ewhuo.mutation.DayCleared() {
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
	if nodes := ewhuo.mutation.DayIDs(); len(nodes) > 0 {
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
	if ewhuo.mutation.ShiftCleared() {
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
	if nodes := ewhuo.mutation.ShiftIDs(); len(nodes) > 0 {
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
	if ewhuo.mutation.RoleCleared() {
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
	if nodes := ewhuo.mutation.RoleIDs(); len(nodes) > 0 {
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
	ewh = &EmployeeWorkingHours{config: ewhuo.config}
	_spec.Assign = ewh.assignValues
	_spec.ScanValues = ewh.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ewhuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employeeworkinghours.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return ewh, nil
}
