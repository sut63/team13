// Code generated by entc, DO NOT EDIT.

package beginwork

const (
	// Label holds the string label denoting the beginwork type in the database.
	Label = "begin_work"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBeginWork holds the string denoting the beginwork field in the database.
	FieldBeginWork = "begin_work"

	// EdgeWhenwork holds the string denoting the whenwork edge name in mutations.
	EdgeWhenwork = "whenwork"

	// Table holds the table name of the beginwork in the database.
	Table = "begin_works"
	// WhenworkTable is the table the holds the whenwork relation/edge.
	WhenworkTable = "employee_working_hours"
	// WhenworkInverseTable is the table name for the EmployeeWorkingHours entity.
	// It exists in this package in order to avoid circular dependency with the "employeeworkinghours" package.
	WhenworkInverseTable = "employee_working_hours"
	// WhenworkColumn is the table column denoting the whenwork relation/edge.
	WhenworkColumn = "begin_work_whenwork"
)

// Columns holds all SQL columns for beginwork fields.
var Columns = []string{
	FieldID,
	FieldBeginWork,
}
