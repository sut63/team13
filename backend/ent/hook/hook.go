// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/team13/app/ent"
)

// The AssessmentFunc type is an adapter to allow the use of ordinary
// function as Assessment mutator.
type AssessmentFunc func(context.Context, *ent.AssessmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AssessmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AssessmentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AssessmentMutation", m)
	}
	return f(ctx, mv)
}

// The BeginWorkFunc type is an adapter to allow the use of ordinary
// function as BeginWork mutator.
type BeginWorkFunc func(context.Context, *ent.BeginWorkMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BeginWorkFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BeginWorkMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BeginWorkMutation", m)
	}
	return f(ctx, mv)
}

// The CompanyFunc type is an adapter to allow the use of ordinary
// function as Company mutator.
type CompanyFunc func(context.Context, *ent.CompanyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompanyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CompanyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompanyMutation", m)
	}
	return f(ctx, mv)
}

// The CustomerFunc type is an adapter to allow the use of ordinary
// function as Customer mutator.
type CustomerFunc func(context.Context, *ent.CustomerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CustomerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CustomerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CustomerMutation", m)
	}
	return f(ctx, mv)
}

// The DayFunc type is an adapter to allow the use of ordinary
// function as Day mutator.
type DayFunc func(context.Context, *ent.DayMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DayFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DayMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DayMutation", m)
	}
	return f(ctx, mv)
}

// The DiscountFunc type is an adapter to allow the use of ordinary
// function as Discount mutator.
type DiscountFunc func(context.Context, *ent.DiscountMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DiscountFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DiscountMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DiscountMutation", m)
	}
	return f(ctx, mv)
}

// The EmployeeFunc type is an adapter to allow the use of ordinary
// function as Employee mutator.
type EmployeeFunc func(context.Context, *ent.EmployeeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmployeeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EmployeeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmployeeMutation", m)
	}
	return f(ctx, mv)
}

// The EmployeeWorkingHoursFunc type is an adapter to allow the use of ordinary
// function as EmployeeWorkingHours mutator.
type EmployeeWorkingHoursFunc func(context.Context, *ent.EmployeeWorkingHoursMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmployeeWorkingHoursFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EmployeeWorkingHoursMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmployeeWorkingHoursMutation", m)
	}
	return f(ctx, mv)
}

// The GetOffWorkFunc type is an adapter to allow the use of ordinary
// function as GetOffWork mutator.
type GetOffWorkFunc func(context.Context, *ent.GetOffWorkMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GetOffWorkFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GetOffWorkMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GetOffWorkMutation", m)
	}
	return f(ctx, mv)
}

// The GiveawayFunc type is an adapter to allow the use of ordinary
// function as Giveaway mutator.
type GiveawayFunc func(context.Context, *ent.GiveawayMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GiveawayFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GiveawayMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GiveawayMutation", m)
	}
	return f(ctx, mv)
}

// The ManagerFunc type is an adapter to allow the use of ordinary
// function as Manager mutator.
type ManagerFunc func(context.Context, *ent.ManagerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ManagerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ManagerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ManagerMutation", m)
	}
	return f(ctx, mv)
}

// The OrderonlineFunc type is an adapter to allow the use of ordinary
// function as Orderonline mutator.
type OrderonlineFunc func(context.Context, *ent.OrderonlineMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderonlineFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderonlineMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderonlineMutation", m)
	}
	return f(ctx, mv)
}

// The OrderproductFunc type is an adapter to allow the use of ordinary
// function as Orderproduct mutator.
type OrderproductFunc func(context.Context, *ent.OrderproductMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderproductFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderproductMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderproductMutation", m)
	}
	return f(ctx, mv)
}

// The PaymentchannelFunc type is an adapter to allow the use of ordinary
// function as Paymentchannel mutator.
type PaymentchannelFunc func(context.Context, *ent.PaymentchannelMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaymentchannelFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PaymentchannelMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaymentchannelMutation", m)
	}
	return f(ctx, mv)
}

// The PositionFunc type is an adapter to allow the use of ordinary
// function as Position mutator.
type PositionFunc func(context.Context, *ent.PositionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PositionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PositionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PositionMutation", m)
	}
	return f(ctx, mv)
}

// The ProductFunc type is an adapter to allow the use of ordinary
// function as Product mutator.
type ProductFunc func(context.Context, *ent.ProductMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProductFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProductMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProductMutation", m)
	}
	return f(ctx, mv)
}

// The PromotionFunc type is an adapter to allow the use of ordinary
// function as Promotion mutator.
type PromotionFunc func(context.Context, *ent.PromotionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PromotionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PromotionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PromotionMutation", m)
	}
	return f(ctx, mv)
}

// The RoleFunc type is an adapter to allow the use of ordinary
// function as Role mutator.
type RoleFunc func(context.Context, *ent.RoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoleMutation", m)
	}
	return f(ctx, mv)
}

// The SalaryFunc type is an adapter to allow the use of ordinary
// function as Salary mutator.
type SalaryFunc func(context.Context, *ent.SalaryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SalaryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SalaryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SalaryMutation", m)
	}
	return f(ctx, mv)
}

// The StockFunc type is an adapter to allow the use of ordinary
// function as Stock mutator.
type StockFunc func(context.Context, *ent.StockMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StockFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StockMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StockMutation", m)
	}
	return f(ctx, mv)
}

// The TypeproductFunc type is an adapter to allow the use of ordinary
// function as Typeproduct mutator.
type TypeproductFunc func(context.Context, *ent.TypeproductMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TypeproductFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TypeproductMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TypeproductMutation", m)
	}
	return f(ctx, mv)
}

// The ZoneproductFunc type is an adapter to allow the use of ordinary
// function as Zoneproduct mutator.
type ZoneproductFunc func(context.Context, *ent.ZoneproductMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ZoneproductFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ZoneproductMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ZoneproductMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	Hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(_ context.Context, m ent.Mutation) (ent.Value, error) {
			return nil, fmt.Errorf("%s operation is not allowed", m.Op())
		})
	}
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
