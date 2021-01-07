// Code generated by entc, DO NOT EDIT.

package employeeworkinghours

const (
	// Label holds the string label denoting the employeeworkinghours type in the database.
	Label = "employeeworkinghours"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeWorkinghour holds the string denoting the workinghour edge name in mutations.
	EdgeWorkinghour = "workinghour"
	// EdgeDay holds the string denoting the day edge name in mutations.
	EdgeDay = "day"
	// EdgeShift holds the string denoting the shift edge name in mutations.
	EdgeShift = "shift"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"

	// Table holds the table name of the employeeworkinghours in the database.
	Table = "employeeworkinghours"
	// WorkinghourTable is the table the holds the workinghour relation/edge.
	WorkinghourTable = "employeeworkinghours"
	// WorkinghourInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	WorkinghourInverseTable = "employees"
	// WorkinghourColumn is the table column denoting the workinghour relation/edge.
	WorkinghourColumn = "employee_whose"
	// DayTable is the table the holds the day relation/edge.
	DayTable = "employeeworkinghours"
	// DayInverseTable is the table name for the Day entity.
	// It exists in this package in order to avoid circular dependency with the "day" package.
	DayInverseTable = "days"
	// DayColumn is the table column denoting the day relation/edge.
	DayColumn = "day_whatday"
	// ShiftTable is the table the holds the shift relation/edge.
	ShiftTable = "employeeworkinghours"
	// ShiftInverseTable is the table name for the Shift entity.
	// It exists in this package in order to avoid circular dependency with the "shift" package.
	ShiftInverseTable = "shifts"
	// ShiftColumn is the table column denoting the shift relation/edge.
	ShiftColumn = "shift_when"
	// RoleTable is the table the holds the role relation/edge.
	RoleTable = "employeeworkinghours"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "roles"
	// RoleColumn is the table column denoting the role relation/edge.
	RoleColumn = "role_todo"
)

// Columns holds all SQL columns for employeeworkinghours fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Employeeworkinghours type.
var ForeignKeys = []string{
	"day_whatday",
	"employee_whose",
	"role_todo",
	"shift_when",
}