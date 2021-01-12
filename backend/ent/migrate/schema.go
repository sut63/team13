// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// AssessmentsColumns holds the columns for the "assessments" table.
	AssessmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "assessment", Type: field.TypeString},
	}
	// AssessmentsTable holds the schema information for the "assessments" table.
	AssessmentsTable = &schema.Table{
		Name:        "assessments",
		Columns:     AssessmentsColumns,
		PrimaryKey:  []*schema.Column{AssessmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CompaniesColumns holds the columns for the "companies" table.
	CompaniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// CompaniesTable holds the schema information for the "companies" table.
	CompaniesTable = &schema.Table{
		Name:        "companies",
		Columns:     CompaniesColumns,
		PrimaryKey:  []*schema.Column{CompaniesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:        "customers",
		Columns:     CustomersColumns,
		PrimaryKey:  []*schema.Column{CustomersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DaysColumns holds the columns for the "days" table.
	DaysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "day", Type: field.TypeString},
	}
	// DaysTable holds the schema information for the "days" table.
	DaysTable = &schema.Table{
		Name:        "days",
		Columns:     DaysColumns,
		PrimaryKey:  []*schema.Column{DaysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DiscountsColumns holds the columns for the "discounts" table.
	DiscountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "sale", Type: field.TypeInt},
	}
	// DiscountsTable holds the schema information for the "discounts" table.
	DiscountsTable = &schema.Table{
		Name:        "discounts",
		Columns:     DiscountsColumns,
		PrimaryKey:  []*schema.Column{DiscountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// EmployeesColumns holds the columns for the "employees" table.
	EmployeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
	}
	// EmployeesTable holds the schema information for the "employees" table.
	EmployeesTable = &schema.Table{
		Name:        "employees",
		Columns:     EmployeesColumns,
		PrimaryKey:  []*schema.Column{EmployeesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// EmployeeWorkingHoursColumns holds the columns for the "employee_working_hours" table.
	EmployeeWorkingHoursColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "day_whatday", Type: field.TypeInt, Nullable: true},
		{Name: "employee_whose", Type: field.TypeInt, Nullable: true},
		{Name: "role_todo", Type: field.TypeInt, Nullable: true},
		{Name: "shift_when", Type: field.TypeInt, Nullable: true},
	}
	// EmployeeWorkingHoursTable holds the schema information for the "employee_working_hours" table.
	EmployeeWorkingHoursTable = &schema.Table{
		Name:       "employee_working_hours",
		Columns:    EmployeeWorkingHoursColumns,
		PrimaryKey: []*schema.Column{EmployeeWorkingHoursColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "employee_working_hours_days_whatday",
				Columns: []*schema.Column{EmployeeWorkingHoursColumns[1]},

				RefColumns: []*schema.Column{DaysColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "employee_working_hours_employees_whose",
				Columns: []*schema.Column{EmployeeWorkingHoursColumns[2]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "employee_working_hours_roles_todo",
				Columns: []*schema.Column{EmployeeWorkingHoursColumns[3]},

				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "employee_working_hours_shifts_when",
				Columns: []*schema.Column{EmployeeWorkingHoursColumns[4]},

				RefColumns: []*schema.Column{ShiftsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GiveawaysColumns holds the columns for the "giveaways" table.
	GiveawaysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "giveaway_name", Type: field.TypeString},
	}
	// GiveawaysTable holds the schema information for the "giveaways" table.
	GiveawaysTable = &schema.Table{
		Name:        "giveaways",
		Columns:     GiveawaysColumns,
		PrimaryKey:  []*schema.Column{GiveawaysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ManagersColumns holds the columns for the "managers" table.
	ManagersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// ManagersTable holds the schema information for the "managers" table.
	ManagersTable = &schema.Table{
		Name:        "managers",
		Columns:     ManagersColumns,
		PrimaryKey:  []*schema.Column{ManagersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// OrderonlinesColumns holds the columns for the "orderonlines" table.
	OrderonlinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "addedtime", Type: field.TypeTime},
		{Name: "stock", Type: field.TypeInt},
		{Name: "customer_id", Type: field.TypeInt, Nullable: true},
		{Name: "paymentchannel_formpaymentchannel", Type: field.TypeInt, Nullable: true},
		{Name: "product_formproductonline", Type: field.TypeInt, Nullable: true},
		{Name: "typeproduct_from_typeproductonline", Type: field.TypeInt, Nullable: true},
	}
	// OrderonlinesTable holds the schema information for the "orderonlines" table.
	OrderonlinesTable = &schema.Table{
		Name:       "orderonlines",
		Columns:    OrderonlinesColumns,
		PrimaryKey: []*schema.Column{OrderonlinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "orderonlines_customers_formcustomer",
				Columns: []*schema.Column{OrderonlinesColumns[3]},

				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "orderonlines_paymentchannels_formpaymentchannel",
				Columns: []*schema.Column{OrderonlinesColumns[4]},

				RefColumns: []*schema.Column{PaymentchannelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "orderonlines_products_formproductonline",
				Columns: []*schema.Column{OrderonlinesColumns[5]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "orderonlines_typeproducts_fromTypeproductonline",
				Columns: []*schema.Column{OrderonlinesColumns[6]},

				RefColumns: []*schema.Column{TypeproductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrderproductsColumns holds the columns for the "orderproducts" table.
	OrderproductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "addedtime", Type: field.TypeTime},
		{Name: "stock", Type: field.TypeInt},
		{Name: "company_companys", Type: field.TypeInt, Nullable: true},
		{Name: "manager_id", Type: field.TypeInt, Nullable: true},
		{Name: "product_products", Type: field.TypeInt, Nullable: true},
		{Name: "typeproduct_typeproducts", Type: field.TypeInt, Nullable: true},
	}
	// OrderproductsTable holds the schema information for the "orderproducts" table.
	OrderproductsTable = &schema.Table{
		Name:       "orderproducts",
		Columns:    OrderproductsColumns,
		PrimaryKey: []*schema.Column{OrderproductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "orderproducts_companies_companys",
				Columns: []*schema.Column{OrderproductsColumns[3]},

				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "orderproducts_managers_managers",
				Columns: []*schema.Column{OrderproductsColumns[4]},

				RefColumns: []*schema.Column{ManagersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "orderproducts_products_products",
				Columns: []*schema.Column{OrderproductsColumns[5]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "orderproducts_typeproducts_Typeproducts",
				Columns: []*schema.Column{OrderproductsColumns[6]},

				RefColumns: []*schema.Column{TypeproductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PaymentchannelsColumns holds the columns for the "paymentchannels" table.
	PaymentchannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "bank", Type: field.TypeString},
	}
	// PaymentchannelsTable holds the schema information for the "paymentchannels" table.
	PaymentchannelsTable = &schema.Table{
		Name:        "paymentchannels",
		Columns:     PaymentchannelsColumns,
		PrimaryKey:  []*schema.Column{PaymentchannelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PositionsColumns holds the columns for the "positions" table.
	PositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "position", Type: field.TypeString},
	}
	// PositionsTable holds the schema information for the "positions" table.
	PositionsTable = &schema.Table{
		Name:        "positions",
		Columns:     PositionsColumns,
		PrimaryKey:  []*schema.Column{PositionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name_product", Type: field.TypeString},
		{Name: "barcode_product", Type: field.TypeString},
		{Name: "mfg", Type: field.TypeString},
		{Name: "exp", Type: field.TypeString},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:        "products",
		Columns:     ProductsColumns,
		PrimaryKey:  []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PromotionsColumns holds the columns for the "promotions" table.
	PromotionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "promotion_name", Type: field.TypeString, Unique: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "discount_fordiscount", Type: field.TypeInt, Nullable: true},
		{Name: "giveaway_forgiveaway", Type: field.TypeInt, Nullable: true},
		{Name: "product_forproduct", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PromotionsTable holds the schema information for the "promotions" table.
	PromotionsTable = &schema.Table{
		Name:       "promotions",
		Columns:    PromotionsColumns,
		PrimaryKey: []*schema.Column{PromotionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "promotions_discounts_fordiscount",
				Columns: []*schema.Column{PromotionsColumns[3]},

				RefColumns: []*schema.Column{DiscountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "promotions_giveaways_forgiveaway",
				Columns: []*schema.Column{PromotionsColumns[4]},

				RefColumns: []*schema.Column{GiveawaysColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "promotions_products_forproduct",
				Columns: []*schema.Column{PromotionsColumns[5]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:        "roles",
		Columns:     RolesColumns,
		PrimaryKey:  []*schema.Column{RolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SalariesColumns holds the columns for the "salaries" table.
	SalariesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "salary", Type: field.TypeFloat64},
		{Name: "bonus", Type: field.TypeFloat64},
		{Name: "salary_datetime", Type: field.TypeTime},
		{Name: "assessment_formassessment", Type: field.TypeInt, Nullable: true},
		{Name: "employee_formemployee", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "position_formposition", Type: field.TypeInt, Nullable: true},
	}
	// SalariesTable holds the schema information for the "salaries" table.
	SalariesTable = &schema.Table{
		Name:       "salaries",
		Columns:    SalariesColumns,
		PrimaryKey: []*schema.Column{SalariesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "salaries_assessments_formassessment",
				Columns: []*schema.Column{SalariesColumns[4]},

				RefColumns: []*schema.Column{AssessmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "salaries_employees_formemployee",
				Columns: []*schema.Column{SalariesColumns[5]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "salaries_positions_formposition",
				Columns: []*schema.Column{SalariesColumns[6]},

				RefColumns: []*schema.Column{PositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ShiftsColumns holds the columns for the "shifts" table.
	ShiftsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "time_start", Type: field.TypeTime},
		{Name: "time_end", Type: field.TypeTime},
	}
	// ShiftsTable holds the schema information for the "shifts" table.
	ShiftsTable = &schema.Table{
		Name:        "shifts",
		Columns:     ShiftsColumns,
		PrimaryKey:  []*schema.Column{ShiftsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// StocksColumns holds the columns for the "stocks" table.
	StocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "priceproduct", Type: field.TypeFloat64},
		{Name: "amount", Type: field.TypeInt},
		{Name: "time", Type: field.TypeTime},
		{Name: "employee_employeestock", Type: field.TypeInt, Nullable: true},
		{Name: "product_stockproduct", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "typeproduct_typestock", Type: field.TypeInt, Nullable: true},
		{Name: "zoneproduct_zonestock", Type: field.TypeInt, Nullable: true},
	}
	// StocksTable holds the schema information for the "stocks" table.
	StocksTable = &schema.Table{
		Name:       "stocks",
		Columns:    StocksColumns,
		PrimaryKey: []*schema.Column{StocksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "stocks_employees_employeestock",
				Columns: []*schema.Column{StocksColumns[4]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "stocks_products_stockproduct",
				Columns: []*schema.Column{StocksColumns[5]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "stocks_typeproducts_typestock",
				Columns: []*schema.Column{StocksColumns[6]},

				RefColumns: []*schema.Column{TypeproductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "stocks_zoneproducts_zonestock",
				Columns: []*schema.Column{StocksColumns[7]},

				RefColumns: []*schema.Column{ZoneproductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TypeproductsColumns holds the columns for the "typeproducts" table.
	TypeproductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "typeproduct", Type: field.TypeString},
	}
	// TypeproductsTable holds the schema information for the "typeproducts" table.
	TypeproductsTable = &schema.Table{
		Name:        "typeproducts",
		Columns:     TypeproductsColumns,
		PrimaryKey:  []*schema.Column{TypeproductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ZoneproductsColumns holds the columns for the "zoneproducts" table.
	ZoneproductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "zone", Type: field.TypeString},
	}
	// ZoneproductsTable holds the schema information for the "zoneproducts" table.
	ZoneproductsTable = &schema.Table{
		Name:        "zoneproducts",
		Columns:     ZoneproductsColumns,
		PrimaryKey:  []*schema.Column{ZoneproductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AssessmentsTable,
		CompaniesTable,
		CustomersTable,
		DaysTable,
		DiscountsTable,
		EmployeesTable,
		EmployeeWorkingHoursTable,
		GiveawaysTable,
		ManagersTable,
		OrderonlinesTable,
		OrderproductsTable,
		PaymentchannelsTable,
		PositionsTable,
		ProductsTable,
		PromotionsTable,
		RolesTable,
		SalariesTable,
		ShiftsTable,
		StocksTable,
		TypeproductsTable,
		ZoneproductsTable,
	}
)

func init() {
	EmployeeWorkingHoursTable.ForeignKeys[0].RefTable = DaysTable
	EmployeeWorkingHoursTable.ForeignKeys[1].RefTable = EmployeesTable
	EmployeeWorkingHoursTable.ForeignKeys[2].RefTable = RolesTable
	EmployeeWorkingHoursTable.ForeignKeys[3].RefTable = ShiftsTable
	OrderonlinesTable.ForeignKeys[0].RefTable = CustomersTable
	OrderonlinesTable.ForeignKeys[1].RefTable = PaymentchannelsTable
	OrderonlinesTable.ForeignKeys[2].RefTable = ProductsTable
	OrderonlinesTable.ForeignKeys[3].RefTable = TypeproductsTable
	OrderproductsTable.ForeignKeys[0].RefTable = CompaniesTable
	OrderproductsTable.ForeignKeys[1].RefTable = ManagersTable
	OrderproductsTable.ForeignKeys[2].RefTable = ProductsTable
	OrderproductsTable.ForeignKeys[3].RefTable = TypeproductsTable
	PromotionsTable.ForeignKeys[0].RefTable = DiscountsTable
	PromotionsTable.ForeignKeys[1].RefTable = GiveawaysTable
	PromotionsTable.ForeignKeys[2].RefTable = ProductsTable
	SalariesTable.ForeignKeys[0].RefTable = AssessmentsTable
	SalariesTable.ForeignKeys[1].RefTable = EmployeesTable
	SalariesTable.ForeignKeys[2].RefTable = PositionsTable
	StocksTable.ForeignKeys[0].RefTable = EmployeesTable
	StocksTable.ForeignKeys[1].RefTable = ProductsTable
	StocksTable.ForeignKeys[2].RefTable = TypeproductsTable
	StocksTable.ForeignKeys[3].RefTable = ZoneproductsTable
}
