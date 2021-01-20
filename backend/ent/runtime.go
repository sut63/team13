// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/team13/app/ent/assessment"
	"github.com/team13/app/ent/customer"
	"github.com/team13/app/ent/employee"
	"github.com/team13/app/ent/manager"
	"github.com/team13/app/ent/orderproduct"
	"github.com/team13/app/ent/paymentchannel"
	"github.com/team13/app/ent/position"
	"github.com/team13/app/ent/product"
	"github.com/team13/app/ent/promotion"
	"github.com/team13/app/ent/role"
	"github.com/team13/app/ent/salary"
	"github.com/team13/app/ent/schema"
	"github.com/team13/app/ent/typeproduct"
	"github.com/team13/app/ent/zoneproduct"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	assessmentFields := schema.Assessment{}.Fields()
	_ = assessmentFields
	// assessmentDescAssessment is the schema descriptor for assessment field.
	assessmentDescAssessment := assessmentFields[0].Descriptor()
	// assessment.AssessmentValidator is a validator for the "assessment" field. It is called by the builders before save.
	assessment.AssessmentValidator = assessmentDescAssessment.Validators[0].(func(string) error)
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescName is the schema descriptor for name field.
	customerDescName := customerFields[0].Descriptor()
	// customer.NameValidator is a validator for the "name" field. It is called by the builders before save.
	customer.NameValidator = customerDescName.Validators[0].(func(string) error)
	// customerDescEmail is the schema descriptor for email field.
	customerDescEmail := customerFields[1].Descriptor()
	// customer.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	customer.EmailValidator = customerDescEmail.Validators[0].(func(string) error)
	// customerDescPassword is the schema descriptor for password field.
	customerDescPassword := customerFields[2].Descriptor()
	// customer.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	customer.PasswordValidator = customerDescPassword.Validators[0].(func(string) error)
	// customerDescAge is the schema descriptor for age field.
	customerDescAge := customerFields[3].Descriptor()
	// customer.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	customer.AgeValidator = customerDescAge.Validators[0].(func(int) error)
	employeeFields := schema.Employee{}.Fields()
	_ = employeeFields
	// employeeDescName is the schema descriptor for name field.
	employeeDescName := employeeFields[0].Descriptor()
	// employee.NameValidator is a validator for the "name" field. It is called by the builders before save.
	employee.NameValidator = employeeDescName.Validators[0].(func(string) error)
	// employeeDescEmail is the schema descriptor for email field.
	employeeDescEmail := employeeFields[1].Descriptor()
	// employee.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	employee.EmailValidator = employeeDescEmail.Validators[0].(func(string) error)
	// employeeDescPassword is the schema descriptor for password field.
	employeeDescPassword := employeeFields[2].Descriptor()
	// employee.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	employee.PasswordValidator = employeeDescPassword.Validators[0].(func(string) error)
	// employeeDescAge is the schema descriptor for age field.
	employeeDescAge := employeeFields[3].Descriptor()
	// employee.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	employee.AgeValidator = employeeDescAge.Validators[0].(func(int) error)
	managerFields := schema.Manager{}.Fields()
	_ = managerFields
	// managerDescName is the schema descriptor for name field.
	managerDescName := managerFields[0].Descriptor()
	// manager.NameValidator is a validator for the "name" field. It is called by the builders before save.
	manager.NameValidator = managerDescName.Validators[0].(func(string) error)
	// managerDescEmail is the schema descriptor for email field.
	managerDescEmail := managerFields[1].Descriptor()
	// manager.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	manager.EmailValidator = managerDescEmail.Validators[0].(func(string) error)
	// managerDescPassword is the schema descriptor for password field.
	managerDescPassword := managerFields[2].Descriptor()
	// manager.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	manager.PasswordValidator = managerDescPassword.Validators[0].(func(string) error)
	orderproductFields := schema.Orderproduct{}.Fields()
	_ = orderproductFields
	// orderproductDescStock is the schema descriptor for stock field.
	orderproductDescStock := orderproductFields[1].Descriptor()
	// orderproduct.StockValidator is a validator for the "stock" field. It is called by the builders before save.
	orderproduct.StockValidator = orderproductDescStock.Validators[0].(func(int) error)
	// orderproductDescShipment is the schema descriptor for shipment field.
	orderproductDescShipment := orderproductFields[2].Descriptor()
	// orderproduct.ShipmentValidator is a validator for the "shipment" field. It is called by the builders before save.
	orderproduct.ShipmentValidator = orderproductDescShipment.Validators[0].(func(string) error)
	// orderproductDescDetail is the schema descriptor for detail field.
	orderproductDescDetail := orderproductFields[3].Descriptor()
	// orderproduct.DetailValidator is a validator for the "detail" field. It is called by the builders before save.
	orderproduct.DetailValidator = orderproductDescDetail.Validators[0].(func(string) error)
	paymentchannelFields := schema.Paymentchannel{}.Fields()
	_ = paymentchannelFields
	// paymentchannelDescBank is the schema descriptor for Bank field.
	paymentchannelDescBank := paymentchannelFields[0].Descriptor()
	// paymentchannel.BankValidator is a validator for the "Bank" field. It is called by the builders before save.
	paymentchannel.BankValidator = paymentchannelDescBank.Validators[0].(func(string) error)
	positionFields := schema.Position{}.Fields()
	_ = positionFields
	// positionDescPosition is the schema descriptor for position field.
	positionDescPosition := positionFields[0].Descriptor()
	// position.PositionValidator is a validator for the "position" field. It is called by the builders before save.
	position.PositionValidator = positionDescPosition.Validators[0].(func(string) error)
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescNameProduct is the schema descriptor for NameProduct field.
	productDescNameProduct := productFields[0].Descriptor()
	// product.NameProductValidator is a validator for the "NameProduct" field. It is called by the builders before save.
	product.NameProductValidator = productDescNameProduct.Validators[0].(func(string) error)
	promotionFields := schema.Promotion{}.Fields()
	_ = promotionFields
	// promotionDescPromotionName is the schema descriptor for PromotionName field.
	promotionDescPromotionName := promotionFields[0].Descriptor()
	// promotion.PromotionNameValidator is a validator for the "PromotionName" field. It is called by the builders before save.
	promotion.PromotionNameValidator = promotionDescPromotionName.Validators[0].(func(string) error)
	// promotionDescDurationPromotion is the schema descriptor for DurationPromotion field.
	promotionDescDurationPromotion := promotionFields[1].Descriptor()
	// promotion.DurationPromotionValidator is a validator for the "DurationPromotion" field. It is called by the builders before save.
	promotion.DurationPromotionValidator = promotionDescDurationPromotion.Validators[0].(func(string) error)
	// promotionDescPrice is the schema descriptor for Price field.
	promotionDescPrice := promotionFields[2].Descriptor()
	// promotion.PriceValidator is a validator for the "Price" field. It is called by the builders before save.
	promotion.PriceValidator = func() func(float64) error {
		validators := promotionDescPrice.Validators
		fns := [...]func(float64) error{
			validators[0].(func(float64) error),
			validators[1].(func(float64) error),
		}
		return func(_Price float64) error {
			for _, fn := range fns {
				if err := fn(_Price); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescRole is the schema descriptor for Role field.
	roleDescRole := roleFields[0].Descriptor()
	// role.RoleValidator is a validator for the "Role" field. It is called by the builders before save.
	role.RoleValidator = roleDescRole.Validators[0].(func(string) error)
	salaryFields := schema.Salary{}.Fields()
	_ = salaryFields
	// salaryDescSalary is the schema descriptor for Salary field.
	salaryDescSalary := salaryFields[0].Descriptor()
	// salary.SalaryValidator is a validator for the "Salary" field. It is called by the builders before save.
	salary.SalaryValidator = salaryDescSalary.Validators[0].(func(float64) error)
	// salaryDescBonus is the schema descriptor for Bonus field.
	salaryDescBonus := salaryFields[1].Descriptor()
	// salary.BonusValidator is a validator for the "Bonus" field. It is called by the builders before save.
	salary.BonusValidator = salaryDescBonus.Validators[0].(func(float64) error)
	typeproductFields := schema.Typeproduct{}.Fields()
	_ = typeproductFields
	// typeproductDescTypeproduct is the schema descriptor for Typeproduct field.
	typeproductDescTypeproduct := typeproductFields[0].Descriptor()
	// typeproduct.TypeproductValidator is a validator for the "Typeproduct" field. It is called by the builders before save.
	typeproduct.TypeproductValidator = typeproductDescTypeproduct.Validators[0].(func(string) error)
	zoneproductFields := schema.Zoneproduct{}.Fields()
	_ = zoneproductFields
	// zoneproductDescZone is the schema descriptor for Zone field.
	zoneproductDescZone := zoneproductFields[0].Descriptor()
	// zoneproduct.ZoneValidator is a validator for the "Zone" field. It is called by the builders before save.
	zoneproduct.ZoneValidator = zoneproductDescZone.Validators[0].(func(string) error)
}
