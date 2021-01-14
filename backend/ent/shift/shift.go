// Code generated by entc, DO NOT EDIT.

package shift

const (
	// Label holds the string label denoting the shift type in the database.
	Label = "shift"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTimeStart holds the string denoting the timestart field in the database.
	FieldTimeStart = "time_start"
	// FieldTimeEnd holds the string denoting the timeend field in the database.
	FieldTimeEnd = "time_end"

	// EdgeWhen holds the string denoting the when edge name in mutations.
	EdgeWhen = "when"

	// Table holds the table name of the shift in the database.
	Table = "shifts"
	// WhenTable is the table the holds the when relation/edge.
	WhenTable = "employee_working_hours"
	// WhenInverseTable is the table name for the EmployeeWorkingHours entity.
	// It exists in this package in order to avoid circular dependency with the "employeeworkinghours" package.
	WhenInverseTable = "employee_working_hours"
	// WhenColumn is the table column denoting the when relation/edge.
	WhenColumn = "shift_when"
)

// Columns holds all SQL columns for shift fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldTimeStart,
	FieldTimeEnd,
}
