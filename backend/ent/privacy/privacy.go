// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/team13/app/ent"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type decisionCtxKey struct{}

// DecisionContext creates a decision context.
func DecisionContext(parent context.Context, decision error) context.Context {
	if decision == nil || errors.Is(decision, Skip) {
		return parent
	}
	return context.WithValue(parent, decisionCtxKey{}, decision)
}

func decisionFromContext(ctx context.Context) (error, bool) {
	decision, ok := ctx.Value(decisionCtxKey{}).(error)
	if ok && errors.Is(decision, Allow) {
		decision = nil
	}
	return decision, ok
}

type (
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}
)

// EvalQuery evaluates a query against a query policy.
func (policy QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalQuery(ctx, q); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}
)

// EvalMutation evaluates a mutation against a mutation policy.
func (policy MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalMutation(ctx, m); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// MutationRuleFunc type is an adapter to allow the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutation policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is the interface that groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AssessmentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AssessmentQueryRuleFunc func(context.Context, *ent.AssessmentQuery) error

// EvalQuery return f(ctx, q).
func (f AssessmentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AssessmentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AssessmentQuery", q)
}

// The AssessmentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AssessmentMutationRuleFunc func(context.Context, *ent.AssessmentMutation) error

// EvalMutation calls f(ctx, m).
func (f AssessmentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AssessmentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AssessmentMutation", m)
}

// The CompanyQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CompanyQueryRuleFunc func(context.Context, *ent.CompanyQuery) error

// EvalQuery return f(ctx, q).
func (f CompanyQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CompanyQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CompanyQuery", q)
}

// The CompanyMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CompanyMutationRuleFunc func(context.Context, *ent.CompanyMutation) error

// EvalMutation calls f(ctx, m).
func (f CompanyMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CompanyMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CompanyMutation", m)
}

// The CustomerQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CustomerQueryRuleFunc func(context.Context, *ent.CustomerQuery) error

// EvalQuery return f(ctx, q).
func (f CustomerQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CustomerQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CustomerQuery", q)
}

// The CustomerMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CustomerMutationRuleFunc func(context.Context, *ent.CustomerMutation) error

// EvalMutation calls f(ctx, m).
func (f CustomerMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CustomerMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CustomerMutation", m)
}

// The DayQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DayQueryRuleFunc func(context.Context, *ent.DayQuery) error

// EvalQuery return f(ctx, q).
func (f DayQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DayQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DayQuery", q)
}

// The DayMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DayMutationRuleFunc func(context.Context, *ent.DayMutation) error

// EvalMutation calls f(ctx, m).
func (f DayMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DayMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DayMutation", m)
}

// The DiscountQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DiscountQueryRuleFunc func(context.Context, *ent.DiscountQuery) error

// EvalQuery return f(ctx, q).
func (f DiscountQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DiscountQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DiscountQuery", q)
}

// The DiscountMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DiscountMutationRuleFunc func(context.Context, *ent.DiscountMutation) error

// EvalMutation calls f(ctx, m).
func (f DiscountMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DiscountMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DiscountMutation", m)
}

// The EmployeeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EmployeeQueryRuleFunc func(context.Context, *ent.EmployeeQuery) error

// EvalQuery return f(ctx, q).
func (f EmployeeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EmployeeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EmployeeQuery", q)
}

// The EmployeeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EmployeeMutationRuleFunc func(context.Context, *ent.EmployeeMutation) error

// EvalMutation calls f(ctx, m).
func (f EmployeeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EmployeeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EmployeeMutation", m)
}

// The EmployeeWorkingHoursQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EmployeeWorkingHoursQueryRuleFunc func(context.Context, *ent.EmployeeWorkingHoursQuery) error

// EvalQuery return f(ctx, q).
func (f EmployeeWorkingHoursQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EmployeeWorkingHoursQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EmployeeWorkingHoursQuery", q)
}

// The EmployeeWorkingHoursMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EmployeeWorkingHoursMutationRuleFunc func(context.Context, *ent.EmployeeWorkingHoursMutation) error

// EvalMutation calls f(ctx, m).
func (f EmployeeWorkingHoursMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EmployeeWorkingHoursMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EmployeeWorkingHoursMutation", m)
}

// The EndWorkQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EndWorkQueryRuleFunc func(context.Context, *ent.EndWorkQuery) error

// EvalQuery return f(ctx, q).
func (f EndWorkQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EndWorkQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EndWorkQuery", q)
}

// The EndWorkMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EndWorkMutationRuleFunc func(context.Context, *ent.EndWorkMutation) error

// EvalMutation calls f(ctx, m).
func (f EndWorkMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EndWorkMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EndWorkMutation", m)
}

// The GiveawayQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GiveawayQueryRuleFunc func(context.Context, *ent.GiveawayQuery) error

// EvalQuery return f(ctx, q).
func (f GiveawayQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GiveawayQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GiveawayQuery", q)
}

// The GiveawayMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GiveawayMutationRuleFunc func(context.Context, *ent.GiveawayMutation) error

// EvalMutation calls f(ctx, m).
func (f GiveawayMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GiveawayMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GiveawayMutation", m)
}

// The ManagerQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ManagerQueryRuleFunc func(context.Context, *ent.ManagerQuery) error

// EvalQuery return f(ctx, q).
func (f ManagerQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ManagerQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ManagerQuery", q)
}

// The ManagerMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ManagerMutationRuleFunc func(context.Context, *ent.ManagerMutation) error

// EvalMutation calls f(ctx, m).
func (f ManagerMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ManagerMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ManagerMutation", m)
}

// The OrderonlineQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OrderonlineQueryRuleFunc func(context.Context, *ent.OrderonlineQuery) error

// EvalQuery return f(ctx, q).
func (f OrderonlineQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OrderonlineQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.OrderonlineQuery", q)
}

// The OrderonlineMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OrderonlineMutationRuleFunc func(context.Context, *ent.OrderonlineMutation) error

// EvalMutation calls f(ctx, m).
func (f OrderonlineMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.OrderonlineMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.OrderonlineMutation", m)
}

// The OrderproductQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OrderproductQueryRuleFunc func(context.Context, *ent.OrderproductQuery) error

// EvalQuery return f(ctx, q).
func (f OrderproductQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OrderproductQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.OrderproductQuery", q)
}

// The OrderproductMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OrderproductMutationRuleFunc func(context.Context, *ent.OrderproductMutation) error

// EvalMutation calls f(ctx, m).
func (f OrderproductMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.OrderproductMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.OrderproductMutation", m)
}

// The PaymentchannelQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PaymentchannelQueryRuleFunc func(context.Context, *ent.PaymentchannelQuery) error

// EvalQuery return f(ctx, q).
func (f PaymentchannelQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PaymentchannelQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PaymentchannelQuery", q)
}

// The PaymentchannelMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PaymentchannelMutationRuleFunc func(context.Context, *ent.PaymentchannelMutation) error

// EvalMutation calls f(ctx, m).
func (f PaymentchannelMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PaymentchannelMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PaymentchannelMutation", m)
}

// The PositionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PositionQueryRuleFunc func(context.Context, *ent.PositionQuery) error

// EvalQuery return f(ctx, q).
func (f PositionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PositionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PositionQuery", q)
}

// The PositionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PositionMutationRuleFunc func(context.Context, *ent.PositionMutation) error

// EvalMutation calls f(ctx, m).
func (f PositionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PositionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PositionMutation", m)
}

// The ProductQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ProductQueryRuleFunc func(context.Context, *ent.ProductQuery) error

// EvalQuery return f(ctx, q).
func (f ProductQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProductQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ProductQuery", q)
}

// The ProductMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ProductMutationRuleFunc func(context.Context, *ent.ProductMutation) error

// EvalMutation calls f(ctx, m).
func (f ProductMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ProductMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ProductMutation", m)
}

// The PromotionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PromotionQueryRuleFunc func(context.Context, *ent.PromotionQuery) error

// EvalQuery return f(ctx, q).
func (f PromotionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PromotionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PromotionQuery", q)
}

// The PromotionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PromotionMutationRuleFunc func(context.Context, *ent.PromotionMutation) error

// EvalMutation calls f(ctx, m).
func (f PromotionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PromotionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PromotionMutation", m)
}

// The RoleQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RoleQueryRuleFunc func(context.Context, *ent.RoleQuery) error

// EvalQuery return f(ctx, q).
func (f RoleQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RoleQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RoleQuery", q)
}

// The RoleMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RoleMutationRuleFunc func(context.Context, *ent.RoleMutation) error

// EvalMutation calls f(ctx, m).
func (f RoleMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RoleMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RoleMutation", m)
}

// The SalaryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SalaryQueryRuleFunc func(context.Context, *ent.SalaryQuery) error

// EvalQuery return f(ctx, q).
func (f SalaryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SalaryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SalaryQuery", q)
}

// The SalaryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SalaryMutationRuleFunc func(context.Context, *ent.SalaryMutation) error

// EvalMutation calls f(ctx, m).
func (f SalaryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SalaryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SalaryMutation", m)
}

// The StartWorkQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type StartWorkQueryRuleFunc func(context.Context, *ent.StartWorkQuery) error

// EvalQuery return f(ctx, q).
func (f StartWorkQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.StartWorkQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.StartWorkQuery", q)
}

// The StartWorkMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type StartWorkMutationRuleFunc func(context.Context, *ent.StartWorkMutation) error

// EvalMutation calls f(ctx, m).
func (f StartWorkMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.StartWorkMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.StartWorkMutation", m)
}

// The StockQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type StockQueryRuleFunc func(context.Context, *ent.StockQuery) error

// EvalQuery return f(ctx, q).
func (f StockQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.StockQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.StockQuery", q)
}

// The StockMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type StockMutationRuleFunc func(context.Context, *ent.StockMutation) error

// EvalMutation calls f(ctx, m).
func (f StockMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.StockMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.StockMutation", m)
}

// The TypeproductQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TypeproductQueryRuleFunc func(context.Context, *ent.TypeproductQuery) error

// EvalQuery return f(ctx, q).
func (f TypeproductQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TypeproductQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TypeproductQuery", q)
}

// The TypeproductMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TypeproductMutationRuleFunc func(context.Context, *ent.TypeproductMutation) error

// EvalMutation calls f(ctx, m).
func (f TypeproductMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TypeproductMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TypeproductMutation", m)
}

// The ZoneproductQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ZoneproductQueryRuleFunc func(context.Context, *ent.ZoneproductQuery) error

// EvalQuery return f(ctx, q).
func (f ZoneproductQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ZoneproductQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ZoneproductQuery", q)
}

// The ZoneproductMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ZoneproductMutationRuleFunc func(context.Context, *ent.ZoneproductMutation) error

// EvalMutation calls f(ctx, m).
func (f ZoneproductMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ZoneproductMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ZoneproductMutation", m)
}
