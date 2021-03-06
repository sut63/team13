// Code generated by entc, DO NOT EDIT.

package role

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"

	// EdgeTodo holds the string denoting the todo edge name in mutations.
	EdgeTodo = "todo"

	// Table holds the table name of the role in the database.
	Table = "roles"
	// TodoTable is the table the holds the todo relation/edge.
	TodoTable = "employee_working_hours"
	// TodoInverseTable is the table name for the EmployeeWorkingHours entity.
	// It exists in this package in order to avoid circular dependency with the "employeeworkinghours" package.
	TodoInverseTable = "employee_working_hours"
	// TodoColumn is the table column denoting the todo relation/edge.
	TodoColumn = "role_todo"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldRole,
}

var (
	// RoleValidator is a validator for the "Role" field. It is called by the builders before save.
	RoleValidator func(string) error
)
