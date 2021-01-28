// Code generated by entc, DO NOT EDIT.

package salary

const (
	// Label holds the string label denoting the salary type in the database.
	Label = "salary"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSalary holds the string denoting the salary field in the database.
	FieldSalary = "salary"
	// FieldBonus holds the string denoting the bonus field in the database.
	FieldBonus = "bonus"
	// FieldSalaryDatetime holds the string denoting the salarydatetime field in the database.
	FieldSalaryDatetime = "salary_datetime"
	// FieldIDEmployee holds the string denoting the idemployee field in the database.
	FieldIDEmployee = "id_employee"
	// FieldAccountNumber holds the string denoting the accountnumber field in the database.
	FieldAccountNumber = "account_number"

	// EdgeAssessment holds the string denoting the assessment edge name in mutations.
	EdgeAssessment = "assessment"
	// EdgePosition holds the string denoting the position edge name in mutations.
	EdgePosition = "position"
	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"

	// Table holds the table name of the salary in the database.
	Table = "salaries"
	// AssessmentTable is the table the holds the assessment relation/edge.
	AssessmentTable = "salaries"
	// AssessmentInverseTable is the table name for the Assessment entity.
	// It exists in this package in order to avoid circular dependency with the "assessment" package.
	AssessmentInverseTable = "assessments"
	// AssessmentColumn is the table column denoting the assessment relation/edge.
	AssessmentColumn = "assessment_formassessment"
	// PositionTable is the table the holds the position relation/edge.
	PositionTable = "salaries"
	// PositionInverseTable is the table name for the Position entity.
	// It exists in this package in order to avoid circular dependency with the "position" package.
	PositionInverseTable = "positions"
	// PositionColumn is the table column denoting the position relation/edge.
	PositionColumn = "position_formposition"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "salaries"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "employee_formemployee"
)

// Columns holds all SQL columns for salary fields.
var Columns = []string{
	FieldID,
	FieldSalary,
	FieldBonus,
	FieldSalaryDatetime,
	FieldIDEmployee,
	FieldAccountNumber,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Salary type.
var ForeignKeys = []string{
	"assessment_formassessment",
	"employee_formemployee",
	"position_formposition",
}

var (
	// SalaryValidator is a validator for the "Salary" field. It is called by the builders before save.
	SalaryValidator func(float64) error
	// BonusValidator is a validator for the "Bonus" field. It is called by the builders before save.
	BonusValidator func(float64) error
	// IDEmployeeValidator is a validator for the "IDEmployee" field. It is called by the builders before save.
	IDEmployeeValidator func(string) error
	// AccountNumberValidator is a validator for the "AccountNumber" field. It is called by the builders before save.
	AccountNumberValidator func(string) error
)
